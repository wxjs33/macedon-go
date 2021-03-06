package macedon

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const test_macedon_conf string = `
[default]
addr: 172.30.23.39:8888

log: macedon.log
level: debug

location: /dns
purge_ips: 172.30.19.35
purge_cmd: touch /tmp/ggg01
ssh_key: /root/.ssh/id_rsa
ssh_port: 22
ssh_user: root
ssh_timeout: 20


consul_addrs: 172.30.23.39:8500
register_location: /v1/agent/service/register
deregister_location: /v1/agent/service/deregister/
read_location: /v1/catalog/service/
domain: lianjia.com
`


func testReadConf(t *testing.T, data string) {
	conf := &Config{}
	tempDir, err := ioutil.TempDir("", "test_log")
	if err != nil {
		t.Fatalf("tempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	path := filepath.Join(tempDir, "test_conf")
	err = ioutil.WriteFile(path, []byte(test_macedon_conf), 0644)
	if err != nil {
		t.Fatalf("writeFile: %v", err)
	}

	_, err = conf.ReadConf(path)
	if err != nil {
		t.Fatal("Test read conf failed")
	}
	t.Log("Test read conf ok")
}

func TestReadConfOk(t *testing.T) {
	testReadConf(t, test_macedon_conf)
}
