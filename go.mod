module github.com/alauda/helm-crds

go 1.12

replace k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab

replace k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190606205144-71ebb8303503

replace sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.2.0-beta.3

replace helm.sh/helm => github.com/alauda/helm v3.0.0-alpha.1.0.20190829021852-0235ba407f6d+incompatible

replace github.com/russross/blackfriday => github.com/russross/blackfriday v1.5.2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190606210616-f848dc7be4a4

replace github.com/deislabs/oras => github.com/deislabs/oras v0.6.0

require (
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/alauda/component-base v0.0.0-20190628064654-a4dafcfd3446
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/fatih/structs v1.1.0
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/thoas/go-funk v0.4.0
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8 // indirect
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980 // indirect
	golang.org/x/sys v0.0.0-20190613124609-5ed2794edfdc // indirect
	google.golang.org/appengine v1.6.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	helm.sh/helm v3.0.0-alpha.1.0.20190613170622-c35dbb7aabf8+incompatible
	k8s.io/api v0.0.0-20190612125737-db0771252981
	k8s.io/apimachinery v0.0.0-20190624085041-961b39a1baa0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v0.3.3
	k8s.io/kube-openapi v0.0.0-20190603182131-db7b694dc208 // indirect
	k8s.io/utils v0.0.0-20190607212802-c55fbcfc754a // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)
