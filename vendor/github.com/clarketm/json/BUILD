load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = [
        "decode.go",
        "encode.go",
        "fold.go",
        "indent.go",
        "scanner.go",
        "stream.go",
        "tables.go",
        "tags.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/clarketm/json",
    importpath = "github.com/clarketm/json",
    visibility = ["//visibility:public"],
)
