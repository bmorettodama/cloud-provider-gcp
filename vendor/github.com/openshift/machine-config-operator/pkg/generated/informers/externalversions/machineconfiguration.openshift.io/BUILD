load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "machineconfiguration_openshift_io",
    srcs = ["interface.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io",
    importpath = "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/internalinterfaces",
        "//vendor/github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1:machineconfiguration_openshift_io",
    ],
)
