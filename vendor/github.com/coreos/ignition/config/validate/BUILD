load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "validate",
    srcs = ["validate.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/config/validate",
    importpath = "github.com/coreos/ignition/config/validate",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ajeddeloh/go-json",
        "//vendor/github.com/coreos/ignition/config/validate/astjson",
        "//vendor/github.com/coreos/ignition/config/validate/astnode",
        "//vendor/github.com/coreos/ignition/config/validate/report",
    ],
)
