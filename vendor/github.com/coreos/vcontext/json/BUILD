load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = ["json.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/vcontext/json",
    importpath = "github.com/coreos/vcontext/json",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-json",
        "//vendor/github.com/coreos/vcontext/tree",
    ],
)
