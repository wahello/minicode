module github.com/rancher/registry-controller

go 1.13

replace k8s.io/client-go => k8s.io/client-go v0.18.0

require (
	github.com/rancher/wrangler v0.6.0
	github.com/rancher/wrangler-api v0.6.0
	github.com/xshrim/gol v0.0.0-20200808053019-d18c98ab37ec
	k8s.io/api v0.18.0
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
