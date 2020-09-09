package jsonoperation

import (
	"encoding/json"
	"fmt"
	"testing"
)

var str = []byte(`{
    "registry-mirrors": [
        "https://wlzfs4t4.mirror.aliyuncs.com",
        "https://wlzfs4t4.mirror.aliyuncs.com"
    ],

    "insecure-registries": [
        "dockerhub.private.rockcontrol.com:5000"
    ],

    "bip":"169.254.31.1/24",
    "max-concurrent-downloads": 10,
    "log-driver": "json-file",
    "log-level": "warn",
    "log-opts": {
        "max-size": "10m",
        "max-file": "3"
        },
    "data-root": "/data/var/lib/docker"
    }`)

func TestMain(t *testing.T) {

	v := make(map[string]interface{})

	err := json.Unmarshal(str, &v)
	if err != nil {
		panic(err)
	}
	// spew.Dump(v)

	v["default-runtime"] = "nvidia"
	v["runtimes"] = map[string]interface{}{
		"nvidia": map[string]interface{}{
			"path":        "/usr/bin/nvidia-container-runtime",
			"runtimeArgs": []string{},
		},
	}

	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
