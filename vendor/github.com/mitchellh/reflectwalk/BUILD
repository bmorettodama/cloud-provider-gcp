load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "reflectwalk",
    srcs = [
        "location.go",
        "location_string.go",
        "reflectwalk.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/mitchellh/reflectwalk",
    importpath = "github.com/mitchellh/reflectwalk",
    visibility = ["//visibility:public"],
)
