module dpCall

go 1.17

replace (
	github.com/containerd/containerd => github.com/containerd/containerd v1.3.10
	github.com/containerd/cri => github.com/containerd/cri v1.19.0
	github.com/cri-o/cri-o => github.com/cri-o/cri-o v1.15.4
	github.com/docker/distribution => github.com/docker/distribution v2.8.0-beta.1+incompatible
	github.com/docker/docker => github.com/docker/docker v1.13.1
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.3
	github.com/kubernetes/cri-api => k8s.io/cri-api v0.22.3
	github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.0-rc5
	golang.org/x/net => golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc => google.golang.org/grpc v1.30.1
	k8s.io/api => k8s.io/api v0.18.19
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.19
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.17
	k8s.io/apiserver => k8s.io/apiserver v0.18.19
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.19
	k8s.io/client-go => k8s.io/client-go v0.18.19
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.19
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.19
	k8s.io/code-generator => k8s.io/code-generator v0.18.19
	k8s.io/component-base => k8s.io/component-base v0.18.19
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.14
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.14
	k8s.io/cri-api => k8s.io/cri-api v0.18.19
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.19
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.19
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.19
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.19
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.19
	k8s.io/kubectl => k8s.io/kubectl v0.18.19
	k8s.io/kubelet => k8s.io/kubelet v0.18.19
	k8s.io/kubernetes => k8s.io/kubernetes v1.23.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.19
	k8s.io/metrics => k8s.io/metrics v0.18.19
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.14
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.22.5
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.19
)

require (
	github.com/aws/aws-sdk-go v1.42.22
	github.com/beevik/etree v1.1.0
	github.com/cenkalti/rpc2 v0.0.0-20210604223624-c1acbc6ec984
	github.com/codeskyblue/go-sh v0.0.0-20200712050446-30169cf553fe
	github.com/containerd/containerd v1.4.11
	github.com/containerd/typeurl v1.0.2
	github.com/cri-o/cri-o v0.0.0-00010101000000-000000000000
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.7+incompatible
	github.com/docker/go-units v0.4.0
	github.com/ericchiang/k8s v1.2.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/go-version v1.4.0
	github.com/hashicorp/serf v0.9.7
	github.com/mitchellh/mapstructure v1.4.3
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/pquerna/cachecontrol v0.0.0-20171018203845-0dec1b30a021
	github.com/russellhaering/goxmldsig v1.1.1
	github.com/samalba/dockerclient v0.0.0-20160531175551-a30362618471
	github.com/sirupsen/logrus v1.8.1
	github.com/streadway/simpleuuid v0.0.0-20130420165545-6617b501e485
	github.com/stretchr/testify v1.7.0
	github.com/vishvananda/netlink v1.1.0
	github.com/vishvananda/netns v0.0.0-20211101163701-50045581ed74
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158
	google.golang.org/grpc v1.40.0
	gopkg.in/ldap.v2 v2.5.1
	gopkg.in/square/go-jose.v2 v2.5.1
	k8s.io/api v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/cri-api v0.0.0
)

require (
	github.com/Microsoft/go-winio v0.4.17 // indirect
	github.com/Microsoft/hcsshim v0.8.22 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/containerd/cgroups v1.0.1 // indirect
	github.com/containerd/continuity v0.1.0 // indirect
	github.com/containerd/fifo v1.0.0 // indirect
	github.com/containerd/ttrpc v1.0.2 // indirect
	github.com/containers/storage v1.13.7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opencontainers/image-spec v1.0.2-0.20190823105129-775207bd45b6 // indirect
	github.com/opencontainers/runc v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.6-0.20210820212750-d4cc65f0b2ff // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210831024726-fe130286e0e2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/klog v1.0.0 // indirect
	lukechampine.com/uint128 v1.1.1 // indirect
	modernc.org/cc/v3 v3.35.22 // indirect
	modernc.org/ccgo/v3 v3.15.1 // indirect
	modernc.org/libc v1.14.1 // indirect
	modernc.org/mathutil v1.4.1 // indirect
	modernc.org/memory v1.0.5 // indirect
	modernc.org/opt v0.1.1 // indirect
	modernc.org/sqlite v1.14.5 // indirect
	modernc.org/strutil v1.1.1 // indirect
	modernc.org/token v1.0.0 // indirect
)

require (
	github.com/cenk/hub v1.0.1 // indirect
	github.com/cenkalti/hub v1.0.1 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/knqyf263/go-rpmdb v0.0.0-20220209103220-0f7a6d951a6d
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
)
