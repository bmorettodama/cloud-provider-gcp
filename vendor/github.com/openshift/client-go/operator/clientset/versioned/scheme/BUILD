load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "scheme",
    srcs = [
        "doc.go",
        "register.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/operator/clientset/versioned/scheme",
    importpath = "github.com/openshift/client-go/operator/clientset/versioned/scheme",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/operator/v1:operator",
        "//vendor/github.com/openshift/api/operator/v1alpha1",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime",
    ],
)
