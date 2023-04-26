load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "externalversions",
    srcs = [
        "factory.go",
        "generic.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/config/informers/externalversions",
    importpath = "github.com/openshift/client-go/config/informers/externalversions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/config/v1:config",
        "//vendor/github.com/openshift/client-go/config/clientset/versioned",
        "//vendor/github.com/openshift/client-go/config/informers/externalversions/config",
        "//vendor/github.com/openshift/client-go/config/informers/externalversions/internalinterfaces",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
