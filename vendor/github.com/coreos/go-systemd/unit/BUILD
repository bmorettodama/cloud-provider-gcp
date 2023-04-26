load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "unit",
    srcs = [
        "deserialize.go",
        "escape.go",
        "option.go",
        "serialize.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/go-systemd/unit",
    importpath = "github.com/coreos/go-systemd/unit",
    visibility = ["//visibility:public"],
)
