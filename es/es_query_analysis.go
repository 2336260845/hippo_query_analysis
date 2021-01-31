package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AnalysisResponse struct {
	Tokens []AnalysisToken `json:"tokens"`
}

type AnalysisToken struct {
	Token       string `json:"token"`
	StartOffset int    `json:"start_offset"`
	EndOffset   int    `json:"end_offset"`
	Type        string `json:"type"`
	Position    int    `json:"position"`
}

func QueryAnalysis(analyzer, query string) ([]string, error) {
	querys := map[string]interface{}{
		"analyzer": analyzer, //智能分词用：ik_smart，最大化分词用：ik_max_word
		"text":     query,
	}

	jsonBody, err := json.Marshal(querys)
	if err != nil {
		return []string{}, fmt.Errorf("QueryAnalysis marshal error, err=%+v", err.Error())
	}

	req, err := http.NewRequest("GET", "/_analyze?pretty=true", bytes.NewReader(jsonBody))
	if err != nil {
		return []string{}, fmt.Errorf("QueryAnalysis NewRequest error, err=%+v", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := esClient.Perform(req)
	if err != nil {
		return []string{}, fmt.Errorf("QueryAnalysis Perform error, err=%+v", err.Error())
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return []string{}, fmt.Errorf("QueryAnalysis Perform error, err=%+v", err.Error())
	}

	//fmt.Println("sss", string(buf.Bytes()))

	var analysis AnalysisResponse
	err = json.Unmarshal(buf.Bytes(), &analysis)
	if err != nil {
		return []string{}, fmt.Errorf("QueryAnalysis Unmarshal error, err=%+v", err.Error())
	}

	var reList []string
	for _, v := range analysis.Tokens {
		reList = append(reList, v.Token)
	}

	return reList, nil
}
