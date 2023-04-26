load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "types",
    srcs = [
        "config.go",
        "disk.go",
        "file.go",
        "filesystem.go",
        "group.go",
        "networkd.go",
        "partition.go",
        "passwd.go",
        "path.go",
        "raid.go",
        "storage.go",
        "systemd.go",
        "unit.go",
        "user.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/config/v1/types",
    importpath = "github.com/coreos/ignition/config/v1/types",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-semver/semver",
        "//vendor/github.com/coreos/ignition/config/shared/errors",
        "//vendor/github.com/coreos/ignition/config/validate/report",
    ],
)
