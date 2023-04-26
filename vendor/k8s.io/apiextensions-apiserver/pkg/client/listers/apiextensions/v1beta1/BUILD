load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1beta1",
    srcs = [
        "customresourcedefinition.go",
        "expansion_generated.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1beta1",
    importpath = "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1beta1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
