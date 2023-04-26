load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go-json",
    srcs = [
        "decode.go",
        "encode.go",
        "fold.go",
        "indent.go",
        "scanner.go",
        "stream.go",
        "tags.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/ajeddeloh/go-json",
    importpath = "github.com/ajeddeloh/go-json",
    visibility = ["//visibility:public"],
)
