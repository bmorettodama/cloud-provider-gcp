load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "k8s_cni_cncf_io",
    srcs = ["register.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io",
    importpath = "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io",
    visibility = ["//visibility:public"],
)
