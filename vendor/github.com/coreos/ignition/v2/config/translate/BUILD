load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "translate",
    srcs = ["translate.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/v2/config/translate",
    importpath = "github.com/coreos/ignition/v2/config/translate",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/coreos/ignition/v2/config/util"],
)
