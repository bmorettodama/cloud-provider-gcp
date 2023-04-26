load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v31tov22",
    srcs = ["v31tov22.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ign-converter/translate/v31tov22",
    importpath = "github.com/coreos/ign-converter/translate/v31tov22",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/ign-converter/util",
        "//vendor/github.com/coreos/ignition/config/v2_2/types",
        "//vendor/github.com/coreos/ignition/config/validate",
        "//vendor/github.com/coreos/ignition/v2/config/v3_1/types",
        "//vendor/github.com/coreos/ignition/v2/config/validate",
    ],
)
