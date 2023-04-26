load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dataurl",
    srcs = [
        "dataurl.go",
        "doc.go",
        "lex.go",
        "rfc2396.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/vincent-petithory/dataurl",
    importpath = "github.com/vincent-petithory/dataurl",
    visibility = ["//visibility:public"],
)
