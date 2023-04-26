load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = [
        "apiserver.go",
        "authentication.go",
        "build.go",
        "clusteroperator.go",
        "clusterversion.go",
        "console.go",
        "dns.go",
        "expansion_generated.go",
        "featuregate.go",
        "image.go",
        "infrastructure.go",
        "ingress.go",
        "network.go",
        "oauth.go",
        "operatorhub.go",
        "project.go",
        "proxy.go",
        "scheduler.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/client-go/config/listers/config/v1",
    importpath = "github.com/openshift/client-go/config/listers/config/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/config/v1:config",
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
