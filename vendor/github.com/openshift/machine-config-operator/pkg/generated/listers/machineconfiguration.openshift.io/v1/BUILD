load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "machineconfiguration_openshift_io",
    srcs = [
        "containerruntimeconfig.go",
        "controllerconfig.go",
        "expansion_generated.go",
        "kubeletconfig.go",
        "machineconfig.go",
        "machineconfigpool.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1",
    importpath = "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1:machineconfiguration_openshift_io",
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
