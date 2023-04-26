load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "apiextensions",
    srcs = ["interface.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions",
    importpath = "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions/v1:apiextensions",
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions/v1beta1",
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/internalinterfaces",
    ],
)
