load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "highperf",
    srcs = ["highperf_controller.go"],
    importpath = "k8s.io/cloud-provider-gcp/pkg/controller/highperf",
    visibility = ["//visibility:public"],
    deps = [
        "//providers/gce",
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1:api",
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned",
        "//vendor/k8s.io/api/core/v1:core",
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime",
        "//vendor/k8s.io/apimachinery/pkg/util/wait",
        "//vendor/k8s.io/client-go/informers/core/v1:core",
        "//vendor/k8s.io/client-go/tools/cache",
        "//vendor/k8s.io/client-go/util/workqueue",
        "//vendor/k8s.io/cloud-provider-gcp/crd/apis/network/v1:network",
        "//vendor/k8s.io/cloud-provider-gcp/crd/apis/network/v1alpha1",
        "//vendor/k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned",
        "//vendor/k8s.io/component-base/metrics/prometheus/controllers",
        "//vendor/k8s.io/klog/v2:klog",
    ],
)

go_test(
    name = "highperf_test",
    srcs = ["highperf_controller_test.go"],
    embed = [":highperf"],
    deps = [
        "//pkg/controller/testutil",
        "//providers/gce",
        "//vendor/github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/fake",
        "//vendor/k8s.io/api/core/v1:core",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/client-go/informers",
        "//vendor/k8s.io/client-go/informers/core/v1:core",
        "//vendor/k8s.io/client-go/kubernetes/fake",
        "//vendor/k8s.io/client-go/tools/cache",
        "//vendor/k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned/fake",
        "//vendor/k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/network/v1:network",
        "//vendor/k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/network/v1alpha1",
        "//vendor/k8s.io/component-base/metrics/prometheus/controllers",
    ],
)
