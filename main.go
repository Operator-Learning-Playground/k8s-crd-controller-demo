package main

import (
	"golanglearning/new_project/crd_practice/pkg/controller"
	"golanglearning/new_project/crd_practice/pkg/k8sconfig"
	"golanglearning/new_project/crd_practice/pkg/type_object"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	logf.SetLogger(zap.New())
	var logs = logf.Log.WithName("crdpractice")

	// 1. 初始化manager
	mgr, err := manager.New(k8sconfig.K8sRestConfig(), manager.Options{})
	if err != nil {
		logs.Error(err, "unable to set up manager")
		os.Exit(1)
	}

	// 2. 管理器加入scheme
	err = k8sconfig.SchemeBuilder.AddToScheme(mgr.GetScheme())
	if err != nil {
		logs.Error(err, "unable add schema")
		os.Exit(1)
	}

	// 3. 控制器
	err = builder.ControllerManagedBy(mgr).
		For(&type_object.Tester{}).
		Complete(controller.NewCrdPracticeController())

	if err != nil {
		logs.Error(err, "unable to create manager")
		os.Exit(1)
	}

	// 4. 管理器启动
	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		logs.Error(err, "unable to start manager")
		os.Exit(1)
	}

}
