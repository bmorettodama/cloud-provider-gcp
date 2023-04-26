load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "sriovnetwork",
    srcs = [
        "doc.go",
        "generated_expansion.go",
        "sriovnetwork.go",
        "sriovnetwork_client.go",
        "sriovnetworknodepolicy.go",
        "sriovnetworknodestate.go",
        "sriovoperatorconfig.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1",
    importpath = "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1:api",
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/scheme",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/rest",
    ],
)
