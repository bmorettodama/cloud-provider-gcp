load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "operator",
    srcs = ["interface.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/operator/informers/externalversions/operator",
    importpath = "github.com/openshift/client-go/operator/informers/externalversions/operator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces",
        "//vendor/github.com/openshift/client-go/operator/informers/externalversions/operator/v1:operator",
        "//vendor/github.com/openshift/client-go/operator/informers/externalversions/operator/v1alpha1",
    ],
)
