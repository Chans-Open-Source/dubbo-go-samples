module github.com/apache/dubbo-go-samples

require (
	github.com/alibaba/sentinel-golang v0.6.2
	github.com/apache/dubbo-go v1.5.2-0.20201213161311-ab774b578174
	github.com/apache/dubbo-go-hessian2 v1.7.0
	github.com/dubbogo/gost v1.9.5
	github.com/emicklei/go-restful/v3 v3.4.0
	github.com/golang/protobuf v1.4.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0
	github.com/stretchr/testify v1.6.1
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace (
	github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.8.0
	github.com/shirou/gopsutil => github.com/shirou/gopsutil v0.0.0-20181107111621-48177ef5f880
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20140225173054-eb6ee6f84d0a
)

go 1.13
