load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [":server"],
)

go_library(
    name = "server",
    importpath = "server",
    srcs = ["server.go"],
    deps = [
        ":request",
        ":response",
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
