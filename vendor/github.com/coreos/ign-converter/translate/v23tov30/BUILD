load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v23tov30",
    srcs = ["v23tov30.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/ign-converter/translate/v23tov30",
    importpath = "github.com/coreos/ign-converter/translate/v23tov30",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/ign-converter/util",
        "//vendor/github.com/coreos/ignition/config/v2_3/types",
        "//vendor/github.com/coreos/ignition/config/validate",
        "//vendor/github.com/coreos/ignition/v2/config/v3_0/types",
        "//vendor/github.com/coreos/ignition/v2/config/validate",
    ],
)
