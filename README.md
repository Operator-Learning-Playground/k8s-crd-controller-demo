## 半手写k8s自定义资源控制器

#### 项目思路与背景
首先使用operator，快速又简单的方法是采用Kubebuilder脚手架完成，但这种封装好的方式，不利于学习创建一个扩展性高的operator。

如果使用client-go包纯手工的创建一个自定义控制器，会需要写非常多k8s内部机制的代码。

因此本项目采用半手写的方式创建控制器，藉由controller-runtime包提供的功能搭建自定义资源的控制器。

```
├── README.md
├── crd # crd资源的yaml
│   ├── crd-test.yaml
│   └── crd.yaml
├── go.mod
├── go.sum
├── main.go
├── pkg
│   ├── common
│   │   └── common.go
│   └── k8sconfig
│       ├── SchemeBuilder.go
│       ├── controller.go
│       ├── initConfig.go
│       ├── tester_deepcopy.go
│       └── tester_types.go
└── resource
    └── config

```