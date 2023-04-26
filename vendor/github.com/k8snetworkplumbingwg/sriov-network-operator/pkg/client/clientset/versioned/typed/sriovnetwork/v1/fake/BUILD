load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fake",
    srcs = [
        "doc.go",
        "fake_sriovnetwork.go",
        "fake_sriovnetwork_client.go",
        "fake_sriovnetworknodepolicy.go",
        "fake_sriovnetworknodestate.go",
        "fake_sriovoperatorconfig.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1/fake",
    importpath = "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1/fake",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1:api",
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1:sriovnetwork",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/rest",
        "//vendor/k8s.io/client-go/testing",
    ],
)
