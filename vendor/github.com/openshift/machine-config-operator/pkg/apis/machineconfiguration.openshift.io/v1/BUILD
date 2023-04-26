load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "machineconfiguration_openshift_io",
    srcs = [
        "doc.go",
        "helpers.go",
        "register.go",
        "types.go",
        "zz_generated.deepcopy.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1",
    importpath = "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/config/v1:config",
        "//vendor/k8s.io/api/core/v1:core",
        "//vendor/k8s.io/apimachinery/pkg/api/resource",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/apimachinery/pkg/util/intstr",
    ],
)
