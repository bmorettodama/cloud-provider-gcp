load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1beta1",
    srcs = [
        "customresourcedefinition.go",
        "interface.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions/v1beta1",
    importpath = "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions/v1beta1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset",
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/internalinterfaces",
        "//vendor/k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1beta1",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
