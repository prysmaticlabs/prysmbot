load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# Download the rules_docker repository at release v0.10.1
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "9ff889216e28c918811b77999257d4ac001c26c1f7c7fb17a79bc28abf74182e",
    strip_prefix = "rules_docker-0.10.1",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.10.1/rules_docker-v0.10.1.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
    ],
    sha256 = "513c12397db1bc9aa46dd62f02dd94b49a9b5d17444d49b5a04c5a89f3053c1c",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
    ],
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

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
