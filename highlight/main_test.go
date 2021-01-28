package main

import (
	"os"
	"testing"

	"github.com/alecthomas/chroma/quick"
)

const (
	Dockerfile = `FROM alpine:3.12
ADD go-bin /usr/local/bin/go-bin
RUN chmod +x /usr/local/bin/go-bin
ENTRYPOINT ["/usr/local/bin/go-bin"]
`
)

const (
	YAML = `apiVersion: v1
	kind: Service
	metadata:
	  annotations:
		helmx/project: '{"name":"web-ops","version":"0.0.0-9c688c4","group":"devops","description":"${PROJECT_DESCRIPTION}"}'
		helmx/upstreams: ""
	  creationTimestamp: "2019-08-28T08:13:22Z"
	  name: web-ops
	  namespace: devops
	spec:
	  clusterIP: 10.64.204.164
	  ports:
	  - name: http-80
		port: 80
		protocol: TCP
		targetPort: 80
	  selector:
		srv: web-ops
	  sessionAffinity: None
	  type: ClusterIP
	status:
	  loadBalancer: {}`
)

func Test_YAML(t *testing.T) {

	if err := quick.Highlight(
		os.Stdout, YAML,
		"YAML", "terminal", "monokai"); err != nil {

		panic(err)
	}

}
func Test_Docker(t *testing.T) {

	if err := quick.Highlight(
		os.Stdout, Dockerfile,
		"Docker", "terminal", "vim"); err != nil {

		panic(err)
	}

}
