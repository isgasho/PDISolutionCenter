package pdiutil

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

var contentTypeMapping = map[string]string{
	".absl":        "ABSL",
	".bo":          "BUSINESS_OBJECT",
	".qry":         "QUERYDEF",
	".xbo":         "EXTENSION_ENTITY",
	".bco":         "BCO",
	".bcc":         "BCSET",
	".uicomponent": "UICOMPONENT",
}

var checkMessageCategoryReg = map[string]*regexp.Regexp{
	"Type Assignment Error":   regexp.MustCompile("Assignment of the type '.*?' to the type '.*?' is not possible."),
	"String Exceed Warning":   regexp.MustCompile("Please ensure that the value doesn’t exceed [\\d]+ characters, otherwise it will be cut off at runtime."),
	"Query Type Warning":      regexp.MustCompile("Use of data type '.*?' is not supported in queries"),
	"Type Recommendation":     regexp.MustCompile("Do not store external document data in unrestricted data type '.*?'. Recommendation .*?$"),
	"Performation Warning":    regexp.MustCompile("Performance Alert: .*?$"),
	"Database Length Warning": regexp.MustCompile("Length of data type '.*?' is restricted to [\\d]+ characters in the data base"),
	"Query Warning":           regexp.MustCompile("Query is not using any selection parameters, which may lead to a long runtime."),
	"Definition Warning":      regexp.MustCompile("Definition of value '.*?' not found; Default value can not be verified."),
	"Namespace Warning":       regexp.MustCompile("Namespace '.*?' already imported"),
}

// CheckMessage is backend check result
type CheckMessage struct {
	Column   string
	Row      string
	Severity string
	FileName string
	Message  string
}

// GetMessageCategory string
func (m CheckMessage) GetMessageCategory() string {
	rt := "Unknown"
	for k, r := range checkMessageCategoryReg {
		if r.MatchString(m.Message) {
			return k
		}
	}
	return rt
}

// IsError message
func (m CheckMessage) IsError() bool {
	return m.Severity == "E"
}

// GetMessageLevel formatted level
// Warning or Error
func (m CheckMessage) GetMessageLevel() string {
	rt := "Unknown"
	switch m.Severity {
	case "W":
		rt = "Warning"
	case "E":
		rt = "Error"
	}
	return rt
}

func (c *PDIClient) backendCheck(xrepPath string) (bool, *[]CheckMessage) {
	canCheck, msgLst := false, []CheckMessage{}

	contentType := contentTypeMapping[filepath.Ext(xrepPath)]

	_, filename := filepath.Split(xrepPath)

	if contentType != "" {
		canCheck = true
		payload := map[string]interface{}{
			"IMPORTING": map[string]interface{}{
				"IV_CONTENT_TYPE": contentType,
				"IT_XREP_PATH":    []string{xrepPath},
			},
		}

		respBody := c.xrepRequest("00163E0115B01DDFB194EC88B8EE8C9B", payload)

		switch contentType {
		case "UICOMPONENT":
			msgList := gjson.Get(respBody, "EXPORTING.ET_MESSAGES").Array()
			for _, msg := range msgList {
				checkMessage := CheckMessage{
					Severity: msg.Get("SEVERITY").String(),
					Message:  msg.Get("TEXT").String(),
					FileName: filename,
				}
				msgLst = append(msgLst, checkMessage)
			}
		default:
			msgList := gjson.Get(respBody, "EXPORTING.ET_MSG_LIST").Array()
			for _, msg := range msgList {
				checkMessage := CheckMessage{
					Column:   strings.TrimSpace(msg.Get("COLUMN_NO").String()),
					Row:      strings.TrimSpace(msg.Get("LINE_NO").String()),
					Severity: msg.Get("SEVERITY").String(),
					Message:  msg.Get("TEXT").String(),
					FileName: filename,
				}
				msgLst = append(msgLst, checkMessage)
			}
		}

	}

	return canCheck, &msgLst
}

// CheckBackendMessageAPI information
func (c *PDIClient) CheckBackendMessageAPI(solution string, concurrent int) []CheckMessage {
	files := c.GetSolutionXrepFileList(solution)
	fileCount := len(files)

	responses := []CheckMessage{}

	asyncResponses := make([]chan *[]CheckMessage, fileCount)
	parallelController := make(chan bool, concurrent)

	for idx, file := range files {
		asyncResponses[idx] = make(chan *[]CheckMessage, 1)
		parallelController <- true
		go func(file string, done chan *[]CheckMessage) {
			_, checkMessage := c.backendCheck(file)
			done <- checkMessage
			<-parallelController
		}(file, asyncResponses[idx])
	}

	for _, response := range asyncResponses {
		r := *(<-response)
		responses = append(responses, r...)
	}

	return responses

}

// CheckBackendMessage information
func (c *PDIClient) CheckBackendMessage(solution string, concurrent int) {

	responses := c.CheckBackendMessageAPI(solution, concurrent)

	for _, r := range responses {
		_, filename := filepath.Split(r.FileName)

		if r.Severity == "E" {
			// any error will cause exit not as zero
			c.exitCode = c.exitCode + 1
		}

		log.Printf("[%s]\t%s(%s,%s): %s\n", r.GetMessageLevel(), filename, r.Row, r.Column, r.Message)
	}

}
