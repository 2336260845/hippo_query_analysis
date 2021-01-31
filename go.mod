module hippo_query_analysis

go 1.15

require (
	github.com/apache/thrift/lib/go/thrift v0.0.0-20210120171102-e27e82c46ba4
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
)

replace github.com/apache/thrift/lib/go/thrift => ./go_thrift
