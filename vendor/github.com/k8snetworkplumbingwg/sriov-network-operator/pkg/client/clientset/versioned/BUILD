load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "versioned",
    srcs = [
        "clientset.go",
        "doc.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned",
    importpath = "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1:sriovnetwork",
        "//vendor/k8s.io/client-go/discovery",
        "//vendor/k8s.io/client-go/rest",
        "//vendor/k8s.io/client-go/util/flowcontrol",
    ],
)
