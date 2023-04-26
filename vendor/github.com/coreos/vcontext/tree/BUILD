load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tree",
    srcs = ["tree.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/vcontext/tree",
    importpath = "github.com/coreos/vcontext/tree",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/coreos/vcontext/path"],
)
