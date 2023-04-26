load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "astjson",
    srcs = ["node.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/config/validate/astjson",
    importpath = "github.com/coreos/ignition/config/validate/astjson",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ajeddeloh/go-json",
        "//vendor/github.com/coreos/ignition/config/validate/astnode",
        "//vendor/go4.org/errorutil",
    ],
)
