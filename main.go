package main

import (
	"golanglearning/new_project/crd_practice/pkg/k8sconfig"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"os"
)

func main() {
	logf.SetLogger(zap.New())
	var logs = logf.Log.WithName("crdPractice")

	// 初始化manager
	mgr, err := manager.New(k8sconfig.K8sRestConfig(), manager.Options{})
	if err != nil {
		logs.Error(err, "unable to set up manager")
		os.Exit(1)
	}

	// 管理器加入scheme
	err = k8sconfig.SchemeBuilder.AddToScheme(mgr.GetScheme())
	if err != nil {
		logs.Error(err, "unable add schema")
		os.Exit(1)
	}

	// 控制器
	err = builder.ControllerManagedBy(mgr).
		For(&k8sconfig.Tester{}).Complete(k8sconfig.NewCrdPracticeController())

	if err != nil {
		logs.Error(err, "unable to create manager")
		os.Exit(1)
	}

	// 管理器启动
	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		logs.Error(err, "unable to start manager")
		os.Exit(1)
	}

}

