package qniblib // import "github.com/qnib/k8s-device-plugin-rcuda/lib"

import (
	"fmt"
	"github.com/zpatrick/go-config"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
	"log"
	"strings"
)

const (
	rcudaCfg = "/etc/qnib-device-plugin/rcuda.ini"
	cfgPrefix = "devices"
)

func check(err error) {
	if err != nil {
		log.Panicln("Fatal:", err)
	}
}

func getDevs(cfg *config.Config, host string) (devs []string) {
	key := fmt.Sprintf("%s.%s", cfgPrefix, host)
	val, err := cfg.String(key)
	if err != nil {
		log.Fatalf("No key '%s' holdinfg a list of device IDs", host)
	}
	devs = strings.Split(val, ",")
	return

}
func getHosts(cfg *config.Config) (hosts []string) {
	key := fmt.Sprintf("%s.%s", cfgPrefix, "hosts")
	val, err := cfg.String(key)
	if err != nil {
		log.Fatal("No key 'hosts' holdinfg a list of hosts with remote GPUs")
	}
	hosts = strings.Split(val, ",")
	return
}

func GetDevices() (devs []*pluginapi.Device) {
	// Read configuration and setup remote GPUs accesible by this node
	cfg, err := NewConfig(rcudaCfg)
	if err != nil {
		return
	}
	for _, host := range getHosts(cfg) {
		for _, devId := range getDevs(cfg, host) {
			dId := fmt.Sprintf("%s:%s", host, devId)
			log.Printf("Add Device: %s", dId)
			devs = append(devs, &pluginapi.Device{ID: dId, Health: pluginapi.Healthy})
		}
	}
	return devs
}

func DeviceExists(devs []*pluginapi.Device, id string) bool {
	for _, d := range devs {
		if d.ID == id {
			return true
		}
	}
	return false
}

