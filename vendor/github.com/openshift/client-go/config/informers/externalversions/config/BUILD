load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = ["interface.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/config/informers/externalversions/config",
    importpath = "github.com/openshift/client-go/config/informers/externalversions/config",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/client-go/config/informers/externalversions/config/v1:config",
        "//vendor/github.com/openshift/client-go/config/informers/externalversions/internalinterfaces",
    ],
)
