load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "validate",
    srcs = ["validate.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/v2/config/validate",
    importpath = "github.com/coreos/ignition/v2/config/validate",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/ignition/v2/config/shared/errors",
        "//vendor/github.com/coreos/ignition/v2/config/util",
        "//vendor/github.com/coreos/vcontext/json",
        "//vendor/github.com/coreos/vcontext/path",
        "//vendor/github.com/coreos/vcontext/report",
        "//vendor/github.com/coreos/vcontext/tree",
        "//vendor/github.com/coreos/vcontext/validate",
    ],
)
