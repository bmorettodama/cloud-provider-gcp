load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "translate",
    srcs = [
        "set.go",
        "translate.go",
        "util.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/fcct/translate",
    importpath = "github.com/coreos/fcct/translate",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/ignition/v2/config/util",
        "//vendor/github.com/coreos/vcontext/path",
    ],
)
