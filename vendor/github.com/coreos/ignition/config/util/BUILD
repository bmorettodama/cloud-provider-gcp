load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "helpers.go",
        "parsingErrors.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/config/util",
    importpath = "github.com/coreos/ignition/config/util",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ajeddeloh/go-json",
        "//vendor/github.com/coreos/ignition/config/shared/errors",
        "//vendor/github.com/coreos/ignition/config/v2_5_experimental/types",
        "//vendor/github.com/coreos/ignition/config/validate/report",
        "//vendor/go4.org/errorutil",
    ],
)
