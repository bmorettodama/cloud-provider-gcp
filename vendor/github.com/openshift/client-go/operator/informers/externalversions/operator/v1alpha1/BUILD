load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1alpha1",
    srcs = [
        "imagecontentsourcepolicy.go",
        "interface.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/operator/informers/externalversions/operator/v1alpha1",
    importpath = "github.com/openshift/client-go/operator/informers/externalversions/operator/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/operator/v1alpha1",
        "//vendor/github.com/openshift/client-go/operator/clientset/versioned",
        "//vendor/github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces",
        "//vendor/github.com/openshift/client-go/operator/listers/operator/v1alpha1",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
