module github.com/ilovelili/dongfeng-core

go 1.12

require (
	github.com/aliyun/aliyun-oss-go-sdk v1.9.8
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-log/log v0.1.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.2.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/consul v1.4.4
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-immutable-radix v1.0.0
	github.com/hashicorp/go-rootcerts v1.0.0
	github.com/hashicorp/golang-lru v0.5.1
	github.com/hashicorp/serf v0.8.2
	github.com/ilovelili/aliyun-client v0.0.0-20190605074008-4fbaa377d984
	github.com/ilovelili/dongfeng-error-code v0.0.0-20190404061658-b59a7f3fe1a3
	github.com/ilovelili/dongfeng-logger v0.0.0-20190403091018-f20598e7c461
	github.com/ilovelili/dongfeng-notification v0.0.0-20190403091005-c2220ec717b4
	github.com/ilovelili/dongfeng-protobuf v0.0.0-20190404052200-ec920597149a
	github.com/ilovelili/dongfeng-shared-lib v0.0.0-20190108085915-4093ff764c36
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/lestrrat-go/jwx v0.0.0-20190331105938-e346d0eba260
	github.com/lestrrat-go/pdebug v0.0.0-20180220043849-39f9a71bcabe
	github.com/mafredri/cdp v0.23.2
	github.com/micro/cli v0.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.0.0
	github.com/micro/go-plugins v0.16.2
	github.com/micro/go-rcache v0.3.0
	github.com/micro/h2c v1.0.0
	github.com/micro/mdns v0.1.0
	github.com/micro/util v0.2.0
	github.com/miekg/dns v1.1.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olivere/dapper v0.0.0-20160315092308-11d01232a968
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.1
	golang.org/x/crypto v0.0.0-20190403202508-8e1b8d32e692
	golang.org/x/net v0.0.0-20190403144856-b630fd6fe46b
	golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e
	golang.org/x/text v0.3.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/appengine v1.5.0
)

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.1

replace github.com/ilovelili/dongfeng-protobuf => github.com/ilovelili/dongfeng-protobuf v0.0.0-20190618065646-870ed0f2e9aa

replace github.com/ilovelili/dongfeng-error-code => github.com/ilovelili/dongfeng-error-code v0.0.0-20190618065903-9bcc1dd6022c

replace github.com/ilovelili/aliyun-client => github.com/ilovelili/aliyun-client v0.0.0-20190605074008-4fbaa377d984

replace github.com/ilovelili/dongfeng-shared-lib => github.com/ilovelili/dongfeng-shared-lib v0.0.0-20190508095909-3b2fd2edf57b

replace github.com/micro/go-micro v1.0.0 => github.com/micro/go-micro v0.14.1
