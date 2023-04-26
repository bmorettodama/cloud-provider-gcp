load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "render",
    srcs = [
        "funcs.go",
        "render.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/render",
    importpath = "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/render",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Masterminds/sprig/v3:sprig",
        "//vendor/github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1:machineconfiguration_openshift_io",
        "//vendor/github.com/openshift/machine-config-operator/pkg/controller/common",
        "//vendor/github.com/pkg/errors",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured",
        "//vendor/k8s.io/apimachinery/pkg/util/yaml",
    ],
)
