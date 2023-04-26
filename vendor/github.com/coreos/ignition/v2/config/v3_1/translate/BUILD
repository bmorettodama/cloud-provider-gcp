load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "translate",
    srcs = ["translate.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/v2/config/v3_1/translate",
    importpath = "github.com/coreos/ignition/v2/config/v3_1/translate",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/ignition/v2/config/translate",
        "//vendor/github.com/coreos/ignition/v2/config/util",
        "//vendor/github.com/coreos/ignition/v2/config/v3_0/types",
        "//vendor/github.com/coreos/ignition/v2/config/v3_1/types",
    ],
)
