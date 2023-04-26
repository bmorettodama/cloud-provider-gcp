load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "report",
    srcs = ["report.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/vcontext/report",
    importpath = "github.com/coreos/vcontext/report",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/vcontext/path",
        "//vendor/github.com/coreos/vcontext/tree",
    ],
)
