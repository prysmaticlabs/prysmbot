load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_docker",
    commit = "9bfcd7dbf0294ed9d11a99da6363fc28df904502",
    remote = "https://github.com/bazelbuild/rules_docker",
    shallow_since = "1596824487 -0400",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "513c12397db1bc9aa46dd62f02dd94b49a9b5d17444d49b5a04c5a89f3053c1c",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

git_repository(
    name = "com_google_protobuf",
    commit = "4cf5bfee9546101d98754d23ff378ff718ba8438",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

go_repository(
    name = "com_github_bwmarrin_discordgo",
    importpath = "github.com/bwmarrin/discordgo",
    sum = "h1:AxjcHGbyBFSC0a3Zx5nDQwbOjU7xai5dXjRnZ0YB7nU=",
    version = "v0.20.3",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_prysmaticlabs_ethereumapis",
    importpath = "github.com/prysmaticlabs/ethereumapis",
    sum = "h1:4Yk1lnr9Q6V8Oxo5YDLiEvOIQGMKA9fDgxFRhiqWRdM=",
    version = "v0.0.0-20200422200834-867e307fa50f",
)

go_repository(
    name = "com_github_prysmaticlabs_prysm",
    importpath = "github.com/prysmaticlabs/prysm",
    sum = "h1:BxCWTHJ3C81N7MTt888k3Of+9WyBrNhBom/BikxEmBo=",
    version = "v0.3.10",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:1N5EYkVAPEywqZRJd7cwnRtCb6xJx7NH3T3WUTF980Q=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_wealdtech_go_bytesutil",
    importpath = "github.com/wealdtech/go-bytesutil",
    sum = "h1:ocEg3Ke2GkZ4vQw5lp46rmO+pfqCCTgq35gqOy8JKVc=",
    version = "v1.1.1",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/grpc",
    sum = "h1:zvIju4sqAGvwKspUQOhwnpcqSbzi7/H6QomNNjTL4sk=",
    version = "v1.27.1",
)

go_repository(
    name = "grpc_ecosystem_grpc_gateway",
    commit = "da7a886035e25b2f274f89b6f3c64bf70a9f6780",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:Ovk5o4KCHLscsOf7hkmbjHsEDk/ycaM9+urOBLus0gI=",
    version = "v0.0.0-20200423195118-18b771bd64f1",
)

go_repository(
    name = "com_github_ferranbt_fastssz",
    importpath = "github.com/ferranbt/fastssz",
    sum = "h1:CxaMtGnKgr9Ar2xLMVddPhnMwYLDsY56w/LxQ/wnsKA=",
    version = "v0.0.0-20200415074633-b062b680417b",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:QPwSajcTUrFriMF1nJ3XzgoqakqQEsnZf9LdXdi2nkI=",
    version = "v0.0.0-20200421231249-e086a090c8fd",
)

go_repository(
    name = "in_gopkg_urfave_cli_v2",
    importpath = "gopkg.in/urfave/cli.v2",
    sum = "h1:OvXt/p4cdwNl+mwcWMq/AxaKFkhdxcjx+tx+qf4EOvY=",
    version = "v2.0.0-20190806201727-b62605953717",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_prysmaticlabs_go_bitfield",
    importpath = "github.com/prysmaticlabs/go-bitfield",
    sum = "h1:cX6YRZnZ9sgMqM5U14llxUiXVNJ3u07Res1IIjTOgtI=",
    version = "v0.0.0-20200322041314-62c2aee71669",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_deckarep_golang_set",
    commit = "cbaa98ba5575e67703b32b4b19f73c91f3c4159e",  # v1.7.1
    importpath = "github.com/deckarep/golang-set",
)

go_repository(
    name = "com_github_google_uuid",
    commit = "0cd6bf5da1e1c83f8b45653022c74f71af0538a4",  # v1.1.1
    importpath = "github.com/google/uuid",
)

go_repository(
    name = "com_github_aristanetworks_goarista",
    commit = "728bce664cf5dfb921941b240828f989a2c8f8e3",
    importpath = "github.com/aristanetworks/goarista",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    commit = "306aecffea325e97f513b3ff0cf7895a5310651d",
    importpath = "github.com/btcsuite/btcd",
)

go_repository(
    name = "com_github_rs_cors",
    commit = "db0fe48135e83b5812a5a31be0eea66984b1b521",  # v1.7.0
    importpath = "github.com/rs/cors",
)

go_repository(
    name = "com_github_pborman_uuid",
    commit = "8b1b92947f46224e3b97bb1a3a5b0382be00d31e",  # v1.2.0
    importpath = "github.com/pborman/uuid",
)

go_repository(
    name = "com_github_go_stack_stack",
    commit = "2fee6af1a9795aafbe0253a0cfbdf668e1fb8a9a",  # v1.8.0
    importpath = "github.com/go-stack/stack",
)

go_repository(
    name = "com_github_ethereum_go_ethereum",
    commit = "0beb54b2147b3473a4c55e5ce6f02643ce403b14",
    importpath = "github.com/ethereum/go-ethereum",
    # Note: go-ethereum is not bazel-friendly with regards to cgo. We have a
    # a fork that has resolved these issues by disabling HID/USB support and
    # some manual fixes for c imports in the crypto package. This is forked
    # branch should be updated from time to time with the latest go-ethereum
    # code.
    remote = "https://github.com/prysmaticlabs/bazel-go-ethereum",
    vcs = "git",
)