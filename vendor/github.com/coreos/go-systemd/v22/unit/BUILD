load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "unit",
    srcs = [
        "deserialize.go",
        "escape.go",
        "option.go",
        "section.go",
        "serialize.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/go-systemd/v22/unit",
    importpath = "github.com/coreos/go-systemd/v22/unit",
    visibility = ["//visibility:public"],
)
