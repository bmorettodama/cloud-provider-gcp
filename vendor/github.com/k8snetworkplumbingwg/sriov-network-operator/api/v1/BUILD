load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "groupversion_info.go",
        "helper.go",
        "sriovibnetwork_types.go",
        "sriovnetwork_types.go",
        "sriovnetworknodepolicy_types.go",
        "sriovnetworknodestate_types.go",
        "sriovnetworkpoolconfig_types.go",
        "sriovoperatorconfig_types.go",
        "zz_generated.deepcopy.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1",
    importpath = "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1:k8s_cni_cncf_io",
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/render",
        "//vendor/k8s.io/api/core/v1:core",
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/client-go/kubernetes",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/client",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/log",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/scheme",
    ],
)
