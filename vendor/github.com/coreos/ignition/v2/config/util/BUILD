load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "config.go",
        "helpers.go",
        "interfaces.go",
        "parsingErrors.go",
        "reflection.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/v2/config/util",
    importpath = "github.com/coreos/ignition/v2/config/util",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-semver/semver",
        "//vendor/github.com/coreos/ignition/v2/config/shared/errors",
        "//vendor/github.com/coreos/vcontext/path",
        "//vendor/github.com/coreos/vcontext/report",
        "//vendor/github.com/coreos/vcontext/tree",
    ],
)
