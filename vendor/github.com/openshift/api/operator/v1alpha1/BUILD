load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1alpha1",
    srcs = [
        "doc.go",
        "register.go",
        "types.go",
        "types_image_content_source_policy.go",
        "zz_generated.deepcopy.go",
        "zz_generated.swagger_doc_generated.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/openshift/api/operator/v1alpha1",
    importpath = "github.com/openshift/api/operator/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/openshift/api/config/v1:config",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
    ],
)
