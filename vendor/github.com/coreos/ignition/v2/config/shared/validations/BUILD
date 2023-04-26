load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "validations",
    srcs = ["unit.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/v2/config/shared/validations",
    importpath = "github.com/coreos/ignition/v2/config/shared/validations",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-systemd/v22/unit",
        "//vendor/github.com/coreos/ignition/v2/config/shared/errors",
    ],
)
