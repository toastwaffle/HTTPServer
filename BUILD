load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        ":response",
        ":router",
        ":server",
    ],
)

go_library(
    name = "server",
    importpath = "server",
    srcs = ["server.go"],
    deps = [
        ":request",
        ":response",
        ":router",
        ":status",
    ],
)

go_library(
    name = "request",
    importpath = "request",
    srcs = ["request.go"],
)

go_library(
    name = "response",
    importpath = "response",
    srcs = ["response.go"],
    deps = [
        ":request",
        ":status",
    ],
)

go_library(
    name = "status",
    importpath = "status",
    srcs = ["status.go"],
)

go_library(
    name = "router",
    importpath = "router",
    srcs = ["router.go"],
    deps = [
        ":request",
        ":response",
        ":status",
    ],
)
