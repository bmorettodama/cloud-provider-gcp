load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v2_2",
    srcs = [
        "append.go",
        "cloudinit.go",
        "config.go",
        "translate.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/config/v2_2",
    importpath = "github.com/coreos/ignition/config/v2_2",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ajeddeloh/go-json",
        "//vendor/github.com/coreos/go-semver/semver",
        "//vendor/github.com/coreos/ignition/config/shared/errors",
        "//vendor/github.com/coreos/ignition/config/util",
        "//vendor/github.com/coreos/ignition/config/v2_1",
        "//vendor/github.com/coreos/ignition/config/v2_1/types",
        "//vendor/github.com/coreos/ignition/config/v2_2/types",
        "//vendor/github.com/coreos/ignition/config/validate",
        "//vendor/github.com/coreos/ignition/config/validate/report",
    ],
)
