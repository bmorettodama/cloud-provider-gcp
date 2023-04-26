load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v0_1",
    srcs = [
        "schema.go",
        "translate.go",
        "validate.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/fcct/base/v0_1",
    importpath = "github.com/coreos/fcct/base/v0_1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/fcct/translate",
        "//vendor/github.com/coreos/ignition/v2/config/v3_0/types",
        "//vendor/github.com/coreos/vcontext/path",
        "//vendor/github.com/coreos/vcontext/report",
        "//vendor/github.com/vincent-petithory/dataurl",
    ],
)
