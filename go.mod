module github.com/bk201/cosplay

go 1.16

replace (
	k8s.io/api => github.com/rancher/kubernetes/staging/src/k8s.io/api v1.19.3-k3s1
	k8s.io/apimachinery => github.com/rancher/kubernetes/staging/src/k8s.io/apimachinery v1.19.3-k3s1
	k8s.io/client-go => github.com/rancher/kubernetes/staging/src/k8s.io/client-go v1.19.3-k3s1
)

require (
	github.com/harvester/harvester-installer v0.1.9-0.20210906063158-207a68ab74c7
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	gopkg.in/yaml.v2 v2.4.0
)
