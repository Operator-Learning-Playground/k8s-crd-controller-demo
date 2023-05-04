package k8sconfig

import (
	"golanglearning/new_project/crd_practice/pkg/common"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"log"
	"os"
)

func K8sRestConfig() *rest.Config {

	// 读取配置
	if os.Getenv("Release") == "1" {
		klog.Info("run in the cluster")
		return k8sRestConfigInPod()
	}

	// 本地测试
	path := common.GetWd()
	config, err := clientcmd.BuildConfigFromFlags("", path+"/resource/config")

	if err != nil {
		log.Fatal(err)
	}
	config.Insecure = true
	return config
}

// k8sRestConfigInPod 集群内部Pod使用
func k8sRestConfigInPod() *rest.Config {

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}
