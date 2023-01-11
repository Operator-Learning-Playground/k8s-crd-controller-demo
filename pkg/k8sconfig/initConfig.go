package k8sconfig

import (
	"golanglearning/new_project/crd_practice/pkg/common"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)


func K8sRestConfig() *rest.Config{
	path := common.GetWd()
	config, err := clientcmd.BuildConfigFromFlags("",path + "/resource/config" )

	if err != nil {
	   log.Fatal(err)
	}
	config.Insecure = true
	return config
}
