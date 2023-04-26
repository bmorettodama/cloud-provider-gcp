load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "versioned",
    srcs = [
        "clientset.go",
        "doc.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/operator/clientset/versioned",
    importpath = "github.com/openshift/client-go/operator/clientset/versioned",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1:operator",
        "//vendor/github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1alpha1",
        "//vendor/k8s.io/client-go/discovery",
        "//vendor/k8s.io/client-go/rest",
        "//vendor/k8s.io/client-go/util/flowcontrol",
    ],
)
