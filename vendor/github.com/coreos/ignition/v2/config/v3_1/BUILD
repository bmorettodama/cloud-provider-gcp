load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v3_1",
    srcs = ["config.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/v2/config/v3_1",
    importpath = "github.com/coreos/ignition/v2/config/v3_1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-semver/semver",
        "//vendor/github.com/coreos/ignition/v2/config/merge",
        "//vendor/github.com/coreos/ignition/v2/config/shared/errors",
        "//vendor/github.com/coreos/ignition/v2/config/util",
        "//vendor/github.com/coreos/ignition/v2/config/v3_0",
        "//vendor/github.com/coreos/ignition/v2/config/v3_1/translate",
        "//vendor/github.com/coreos/ignition/v2/config/v3_1/types",
        "//vendor/github.com/coreos/ignition/v2/config/validate",
        "//vendor/github.com/coreos/vcontext/report",
    ],
)
