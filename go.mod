module github.com/vdaas/vald

go 1.17

replace (
	cloud.google.com/go => cloud.google.com/go v0.97.0
	cloud.google.com/go/monitoring => cloud.google.com/go/monitoring v1.1.1-0.20211106003501-0138b86d100d
	cloud.google.com/go/profiler => cloud.google.com/go/profiler v0.1.2-0.20211106003501-0138b86d100d
	cloud.google.com/go/storage => cloud.google.com/go/storage v1.18.3-0.20211106003501-0138b86d100d
	cloud.google.com/go/trace => cloud.google.com/go/trace v1.0.1-0.20211106003501-0138b86d100d
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.1-0.20211129235838-50e09bb39af1+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.11.23-0.20211129235838-50e09bb39af1
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.9.18-0.20211129235838-50e09bb39af1
	github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.42.17
	github.com/chzyer/logex => github.com/chzyer/logex v1.1.11-0.20170329064859-445be9e134b2
	github.com/coreos/etcd => go.etcd.io/etcd v3.3.27+incompatible
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/docker/docker => github.com/moby/moby v20.10.11+incompatible
	github.com/envoyproxy/protoc-gen-validate => github.com/envoyproxy/protoc-gen-validate v0.6.2
	github.com/go-sql-driver/mysql => github.com/go-sql-driver/mysql v1.6.0
	github.com/gocql/gocql => github.com/gocql/gocql v0.0.0-20211015133455-b225f9b53fa1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf => github.com/golang/protobuf v1.5.2
	github.com/google/gnostic => github.com/google/gnostic v0.5.7
	github.com/google/go-cmp => github.com/google/go-cmp v0.5.6
	github.com/google/pprof => github.com/google/pprof v0.0.0-20211122183932-1daafda22083
	github.com/gophercloud/gophercloud => github.com/gophercloud/gophercloud v0.23.0
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	github.com/hailocab/go-hostpool => github.com/kpango/go-hostpool v0.0.0-20210303030322-aab80263dcd0
	github.com/klauspost/compress => github.com/klauspost/compress v1.13.7-0.20211201171849-25adde538381
	github.com/kpango/glg => github.com/kpango/glg v1.6.4
	github.com/tensorflow/tensorflow => github.com/tensorflow/tensorflow v2.1.2+incompatible
	github.com/zeebo/xxh3 => github.com/zeebo/xxh3 v1.0.1
	go.uber.org/goleak => go.uber.org/goleak v1.1.12
	go.uber.org/multierr => go.uber.org/multierr v1.7.0
	go.uber.org/zap => go.uber.org/zap v1.19.1
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871
	golang.org/x/exp => golang.org/x/exp v0.0.0-20211129234152-8a230f1f7d7a
	golang.org/x/image => golang.org/x/image v0.0.0-20211028202545-6944b10bf410
	golang.org/x/lint => golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/mobile => golang.org/x/mobile v0.0.0-20211109191125-d61a72f26a1a
	golang.org/x/mod => golang.org/x/mod v0.5.1
	golang.org/x/net => golang.org/x/net v0.0.0-20211201190559-0a0e4e1bb54c
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/sync => golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys => golang.org/x/sys v0.0.0-20211124211545-fe61309f8881
	golang.org/x/term => golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	golang.org/x/text => golang.org/x/text v0.3.7
	golang.org/x/time => golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11
	golang.org/x/tools => golang.org/x/tools v0.1.7
	golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/api => google.golang.org/api v0.60.0
	google.golang.org/appengine => google.golang.org/appengine v1.6.7
	google.golang.org/grpc => google.golang.org/grpc v1.42.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc => google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	k8s.io/api => k8s.io/api v0.22.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.2
	k8s.io/apiserver => k8s.io/apiserver v0.22.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.2
	k8s.io/client-go => k8s.io/client-go v0.22.2
	k8s.io/code-generator => k8s.io/code-generator v0.22.2
	k8s.io/component-base => k8s.io/component-base v0.22.2
	k8s.io/gengo => k8s.io/gengo v0.0.0-20211129171323-c02415ce4185
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.10.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
	k8s.io/metrics => k8s.io/metrics v0.22.2
	k8s.io/utils => k8s.io/utils v0.0.0-20211116205334-6203023598ed
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.26
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.10.2
	sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.10.1
	sigs.k8s.io/kustomize/kyaml => sigs.k8s.io/kustomize/kyaml v0.13.0
	sigs.k8s.io/structured-merge-diff/v4 => sigs.k8s.io/structured-merge-diff/v4 v4.2.0
	sigs.k8s.io/yaml => sigs.k8s.io/yaml v1.3.0
)

require (
	cloud.google.com/go/profiler v0.0.0-00010101000000-000000000000
	cloud.google.com/go/storage v1.16.1
	code.cloudfoundry.org/bytefmt v0.0.0-20211005130812-5bb3c17173e5
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.4.0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.10
	github.com/aws/aws-sdk-go v1.40.34
	github.com/envoyproxy/protoc-gen-validate v0.6.2
	github.com/fsnotify/fsnotify v1.5.1
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gocql/gocql v0.0.0-20200131111108-92af2e088537
	github.com/gocraft/dbr/v2 v2.7.2
	github.com/google/go-cmp v0.5.6
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-version v1.3.0
	github.com/json-iterator/go v1.1.12
	github.com/klauspost/compress v1.13.6
	github.com/kpango/fastime v1.0.17
	github.com/kpango/fuid v0.0.0-20210407064122-2990e29e1ea5
	github.com/kpango/gache v1.2.6
	github.com/kpango/glg v1.6.2
	github.com/kr/pretty v0.3.0 // indirect
	github.com/leanovate/gopter v0.2.9
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/pierrec/lz4/v3 v3.3.4
	github.com/quasilyte/go-ruleguard v0.3.13
	github.com/quasilyte/go-ruleguard/dsl v0.3.10
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/scylladb/gocqlx v1.5.0
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/tensorflow/tensorflow v0.0.0-00010101000000-000000000000
	github.com/zeebo/xxh3 v0.12.0
	go.opencensus.io v0.23.0
	go.uber.org/automaxprocs v1.4.0
	go.uber.org/goleak v1.1.11-0.20210813005559-691160354723
	go.uber.org/zap v1.19.0
	gocloud.dev v0.24.0
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486
	golang.org/x/tools v0.1.8-0.20211029000441-d6a9af8af023
	gonum.org/v1/hdf5 v0.0.0-20210714002203-8c5d23bc6946
	gonum.org/v1/plot v0.10.0
	google.golang.org/api v0.63.0
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	inet.af/netaddr v0.0.0-20211027220019-c74959edd3b6
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/cli-runtime v0.0.0-00010101000000-000000000000
	k8s.io/client-go v0.22.2
	k8s.io/metrics v0.0.0-00010101000000-000000000000
	sigs.k8s.io/controller-runtime v0.0.0-00010101000000-000000000000
)
