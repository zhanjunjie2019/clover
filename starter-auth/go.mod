module github.com/zhanjunjie2019/clover/starter-auth

go 1.19

replace (
	github.com/zhanjunjie2019/clover/core => ../core
	github.com/zhanjunjie2019/clover/global => ../global
	github.com/zhanjunjie2019/clover/share => ../share
)

require (
	github.com/alibaba/ioc-golang v1.0.2
	github.com/gin-gonic/gin v1.8.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/swaggo/swag v1.8.8
	github.com/zhanjunjie2019/clover/core v0.0.0-20221117082526-cf9f91f74a29
	github.com/zhanjunjie2019/clover/global v0.0.0-20221125090618-3417abc3a8bf
	github.com/zhanjunjie2019/clover/share v0.0.0-20221117082526-cf9f91f74a29
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/gorm v1.24.2
)

require (
	dubbo.apache.org/dubbo-go/v3 v3.0.2 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/MrAlias/flow v0.1.5 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20221026131551-cf6655e29de4 // indirect
	github.com/Shopify/sarama v1.37.2 // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/alibaba/sentinel-golang v1.0.4 // indirect
	github.com/apache/dubbo-getty v1.4.9 // indirect
	github.com/apache/thrift v0.17.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.7.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cloudflare/circl v1.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/badger/v3 v3.2103.4 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dubbogo/gost v1.13.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-micro/plugins/v4/config/encoder/yaml v1.2.0 // indirect
	github.com/go-micro/plugins/v4/registry/consul v1.2.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/runtime v0.24.2 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.22.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/go-redis/redis/extra/rediscmd/v9 v9.0.0-rc.1 // indirect
	github.com/go-redis/redis/extra/redisotel/v9 v9.0.0-rc.1 // indirect
	github.com/go-redis/redis/v9 v9.0.0-rc.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/gocql/gocql v1.2.1 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v22.11.23+incompatible // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/consul/api v1.17.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.3.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.6 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/log15 v0.0.0-20221122034931-555555054819 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jaegertracing/jaeger v1.39.1-0.20221110195127-14c11365a856 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.3 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/knadh/koanf v1.4.4 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/lufia/plan9stats v0.0.0-20220913051719-115f729f3c8c // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.1.17 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/nsqio/go-nsq v1.1.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/olivere/elastic v6.2.37+incompatible // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.66.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.66.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/petermattis/goid v0.0.0-20221018141743-354ef7f2fd21 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20220216144756-c35f1ee13d7c // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shirou/gopsutil/v3 v3.22.10 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.14.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a // indirect
	github.com/swaggo/gin-swagger v1.5.3 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.17 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.17 // indirect
	github.com/urfave/cli/v2 v2.23.5 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go-micro.dev/v4 v4.9.0 // indirect
	go.mongodb.org/mongo-driver v1.11.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.66.0 // indirect
	go.opentelemetry.io/collector/component v0.66.0 // indirect
	go.opentelemetry.io/collector/consumer v0.66.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.66.0 // indirect
	go.opentelemetry.io/collector/pdata v0.66.0 // indirect
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.66.0 // indirect
	go.opentelemetry.io/collector/semconv v0.66.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.36.4 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.36.4 // indirect
	go.opentelemetry.io/otel v1.11.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.11.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.11.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.11.1 // indirect
	go.opentelemetry.io/otel/metric v0.33.0 // indirect
	go.opentelemetry.io/otel/sdk v1.11.1 // indirect
	go.opentelemetry.io/otel/trace v1.11.1 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/term v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.4.4 // indirect
	gorm.io/driver/postgres v1.4.5 // indirect
)
