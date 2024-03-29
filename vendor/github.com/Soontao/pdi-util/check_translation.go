package pdiutil

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// TranslationStatus message
type TranslationStatus struct {
	FileName     string
	AllTextCount string
	Info         map[string]TranslationStatusInfo
}

// TranslationStatusInfo detail
type TranslationStatusInfo struct {
	Language        string
	TranslatedCount string
}

var translationCheckList = map[string]bool{
	".uicomponent": true, ".bo": true, ".codelist": true,
}

func (c *PDIClient) translationInformation(xrepPath string) (bool, *TranslationStatus) {
	canCheck := translationCheckList[filepath.Ext(xrepPath)]
	rt := &TranslationStatus{
		FileName: xrepPath,
		Info:     map[string]TranslationStatusInfo{},
	}

	payload := map[string]interface{}{
		"IMPORTING": map[string]interface{}{
			"IT_PATH": []string{xrepPath},
		},
	}

	if canCheck {
		respBody := c.xrepRequest("00163E01138A1EE0AFEA287164321C26", payload)
		textCount := strings.TrimSpace(gjson.Get(respBody, "EXPORTING.EV_NUMBER_OF_TEXTS").String())
		rt.AllTextCount = textCount
		checkInfoList := gjson.Get(respBody, "EXPORTING.ET_CHECK_INFO").Array()
		for _, jsonInfo := range checkInfoList {
			info := TranslationStatusInfo{}
			info.Language = jsonInfo.Get("LANGUAGE").String()
			info.TranslatedCount = strings.TrimSpace(jsonInfo.Get("TEXTCOUNT").String())

			rt.Info[info.Language] = info
		}

	}

	return canCheck, rt
}

// CheckTranslationAPI used for programming
func (c *PDIClient) CheckTranslationAPI(solution string, concurrent int) []TranslationStatus {
	files := c.GetSolutionXrepFileList(solution)
	fileCount := len(files)

	responses := []TranslationStatus{}

	asyncResponses := make([]chan *TranslationStatus, fileCount)
	parallexController := make(chan bool, concurrent)

	for idx, file := range files {
		asyncResponses[idx] = make(chan *TranslationStatus, 1)
		parallexController <- true
		go func(file string, done chan *TranslationStatus) {
			canCheck, checkMessage := c.translationInformation(file)
			if canCheck {
				done <- checkMessage
			} else {
				done <- nil
			}
			<-parallexController
		}(file, asyncResponses[idx])
	}

	for _, response := range asyncResponses {
		r := (<-response)
		if r != nil {
			responses = append(responses, *r)
		}
	}

	return responses

}

// CheckTranslation information
func (c *PDIClient) CheckTranslation(solution string, concurrent int, language string) {

	// need check language here

	responses := c.CheckTranslationAPI(solution, concurrent)

	for _, r := range responses {
		_, filename := filepath.Split(r.FileName)
		targetCount, err := strconv.Atoi(r.AllTextCount)

		if err != nil {
			panic(err)
		}

		translatedInfo := r.Info[language]
		translatedCount, err := strconv.Atoi(translatedInfo.TranslatedCount)
		if err != nil {
			panic(err)
		}
		if translatedCount < targetCount {
			c.exitCode = c.exitCode + 1
			log.Printf("For language %s, %d of %d texts have beed translated, file(%s)\n", language, translatedCount, targetCount, filename)
		} else {
			log.Printf("For language %s, fully translated, file(%s)\n", language, filename)
		}

	}

}
