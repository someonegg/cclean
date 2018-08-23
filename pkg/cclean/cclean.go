package cclean

import (
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

func Clean(addr string) {

	config := api.DefaultConfig()
	if addr != "" {
		config.Address = addr
	}

	client, err := api.NewClient(config)
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	allNodes, _, err := client.Catalog().Nodes(nil)
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	allClients := map[string]*api.Client{}
	for _, node := range allNodes {
		tmpConfig := api.DefaultConfig()
		tmpConfig.Address = node.Address + ":8500"
		tmpClient, err := api.NewClient(tmpConfig)
		if err != nil {
			logrus.Errorf("Client: %s create Failed!", tmpConfig.Address)
		} else {
			allClients[tmpConfig.Address] = tmpClient
		}
	}

	for address, tmpClient := range allClients {

		logrus.Infof("Clean client ===> %s", address)

		allChecks, err := tmpClient.Agent().Checks()
		if err != nil {
			logrus.Errorf("Failed to get client [%s] checks!", address)
			continue
		}

		for _, v := range allChecks {
			if v.Status == "critical" {
				logrus.Infof("Deregister: %s", v.ServiceID)
				err := tmpClient.Agent().ServiceDeregister(v.ServiceID)
				if err != nil {
					logrus.Errorf("Failed to clean [%s] service %s", address, v.ServiceID)
				}
			}
		}

		logrus.Infoln("Done.")
	}

}
