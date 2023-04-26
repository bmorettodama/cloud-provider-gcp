load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "clients",
    srcs = ["builder.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/machine-config-operator/internal/clients",
    importpath = "github.com/openshift/machine-config-operator/internal/clients",
    visibility = ["//vendor/github.com/openshift/machine-config-operator:__subpackages__"],
    deps = [
        "//vendor/github.com/golang/glog",
        "//vendor/github.com/openshift/client-go/config/clientset/versioned",
        "//vendor/github.com/openshift/client-go/operator/clientset/versioned",
        "//vendor/github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned",
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset",
        "//vendor/k8s.io/client-go/kubernetes",
        "//vendor/k8s.io/client-go/rest",
        "//vendor/k8s.io/client-go/tools/clientcmd",
    ],
)
