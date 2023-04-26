load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "machineconfiguration_openshift_io",
    srcs = [
        "containerruntimeconfig.go",
        "controllerconfig.go",
        "interface.go",
        "kubeletconfig.go",
        "machineconfig.go",
        "machineconfigpool.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1",
    importpath = "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1:machineconfiguration_openshift_io",
        "//vendor/github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned",
        "//vendor/github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/internalinterfaces",
        "//vendor/github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1:machineconfiguration_openshift_io",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
