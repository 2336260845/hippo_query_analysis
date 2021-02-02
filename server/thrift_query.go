package server

import (
	"context"
	"fmt"
	"github.com/2336260845/hippo_search/gen-go/query_analysis"
	"github.com/apache/thrift/lib/go/thrift"
	"hippo_query_analysis/config"
	"hippo_query_analysis/es"
)

type QueryAnalysisServer struct{}

func (qas *QueryAnalysisServer) QueryAnalysis(_ context.Context, req *query_analysis.QueryParam) (r []string, err error) {
	if req.Query == "" {
		return []string{}, fmt.Errorf("query is empty")
	}

	if req.Analysis == "" {
		return []string{}, fmt.Errorf("analysisModel is empty")
	}

	return es.QueryAnalysis(req.Analysis, req.Query)
}

func ThriftInit(conf *config.Config) {
	transport, err := thrift.NewTServerSocket(conf.ThriftAddress)
	if err != nil {
		panic(fmt.Sprintf("ThriftInit NewTServerSocket error, err=%+v", err.Error()))
	}

	handler := &QueryAnalysisServer{}
	processor := query_analysis.NewQueryAnalysisServiceProcessor(handler)
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(fmt.Sprintf("ThriftInit thrift Serve error, err=%+v", err.Error()))
	}
}
