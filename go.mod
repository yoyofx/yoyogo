module github.com/maxzhang1985/yoyogo

go 1.13

require (
	github.com/defval/inject v1.5.2
	github.com/golang/protobuf v1.3.2
	github.com/lucas-clemente/quic-go v0.12.1
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go/codec v1.1.7
	gopkg.in/yaml.v2 v2.2.4
)

replace (
	github.com/defval/inject v1.5.2 => github.com/maxzhang1985/inject v1.5.0
	github.com/golang/protobuf v1.3.0 => github.com/golang/protobuf v1.3.2
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191029031824-8986dd9e96cf
	golang.org/x/net => github.com/golang/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191029155521-f43be2a4598c
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20191030003036-b2a7f28a184a
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191011141410-1b5146add898
)
