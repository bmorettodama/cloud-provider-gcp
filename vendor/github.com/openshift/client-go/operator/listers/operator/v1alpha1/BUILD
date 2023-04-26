load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1alpha1",
    srcs = [
        "expansion_generated.go",
        "imagecontentsourcepolicy.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/operator/listers/operator/v1alpha1",
    importpath = "github.com/openshift/client-go/operator/listers/operator/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/operator/v1alpha1",
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
