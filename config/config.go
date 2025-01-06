package config

import (
	"log"
	"os"
	"strconv"
)

func GetEtcdAddr() string {
	if etcdAddr := os.Getenv("ETCD_ADDR"); etcdAddr != "" {
		return etcdAddr
	}
	panic("config.GetEtcdAddr: etcd addr is empty")
}

func GetAvailablePort() int {
	portStr := os.Getenv("AVAILABLE_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}
