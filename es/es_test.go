package es

import (
	"hippo_query_analysis/config"
	"testing"
)

func TestQueryAnalysis(t *testing.T) {
	config.InitConfig("../script")

	InitEsClient(config.GetConfig())
	re, err := QueryAnalysis(AnalyzerIkSmart, "你是不是我最疼爱的人")
	if err != nil {
		t.Fatalf("QueryAnalysis error, err=%+v", err.Error())
	}

	t.Logf("TestQueryAnalysis re=%+v", re)

	re, err = QueryAnalysis(AnalyzerIkMaxWord, "你是不是我最疼爱的人")
	if err != nil {
		t.Fatalf("QueryAnalysis error, err=%+v", err.Error())
	}

	t.Logf("TestQueryAnalysis re=%+v", re)
}
