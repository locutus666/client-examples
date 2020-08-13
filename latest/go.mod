module github.com/cilium/client-examples/latest

go 1.15

require github.com/cilium/cilium v1.8.0-rc1.0.20200813184924-df367832502e

replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	k8s.io/client-go => github.com/cilium/client-go v0.0.0-20200725133211-0bdb134c37db
)
