module github.com/cilium/client-examples/v1.8

go 1.14

require github.com/cilium/cilium v1.8.2

replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	k8s.io/client-go => github.com/cilium/client-go v0.0.0-20200725133211-0bdb134c37db
)
