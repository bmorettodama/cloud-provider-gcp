load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "types",
    srcs = [
        "config.go",
        "directory.go",
        "disk.go",
        "file.go",
        "filesystem.go",
        "ignition.go",
        "mode.go",
        "node.go",
        "partition.go",
        "passwd.go",
        "path.go",
        "raid.go",
        "schema.go",
        "unit.go",
        "url.go",
        "verification.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ignition/config/v2_1/types",
    importpath = "github.com/coreos/ignition/config/v2_1/types",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-semver/semver",
        "//vendor/github.com/coreos/go-systemd/unit",
        "//vendor/github.com/coreos/ignition/config/shared/errors",
        "//vendor/github.com/coreos/ignition/config/shared/validations",
        "//vendor/github.com/coreos/ignition/config/validate/report",
        "//vendor/github.com/vincent-petithory/dataurl",
    ],
)
