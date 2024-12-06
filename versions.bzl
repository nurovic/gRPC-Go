"""Defines the dependencies' versions.

This allow us to keep the book generic and update cheaply the versions.
"""

# Go
# The last version can be found here: https://go.dev/dl/
# Note: You do not need to download the last version locally,
#       Just update this variable. Bazel will download it for you.
GO_VERSION = "1.22.2"

# Rules_go
# The last version and SHA256 can be found here: https://github.com/bazelbuild/rules_go/releases
RULES_GO_VERSION = "v0.46.0"
RULES_GO_SHA256 = "80a98277ad1311dacd837f9b16db62887702e9f1d1c4c9f796d0121a46c8e184"

# Gazelle
# The last version and SHA256 can be found here: https://github.com/bazelbuild/bazel-gazelle/blob/master/README.rst
GAZELLE_VERSION = "v0.36.0"
GAZELLE_SHA256 = "75df288c4b31c81eb50f51e2e14f4763cb7548daae126817247064637fd9ea62"

# Protobuf
# The last version can be found here: https://github.com/protocolbuffers/protobuf/releases
# Note: You do not need to download the last version locally,
#       Just update this variable. Bazel will download it for you.
PROTO_VERSION = "v25.0"

# protoc-gen-validate
# The last version can be found here: https://github.com/bufbuild/protoc-gen-validate/releases
PROTOC_GEN_VALIDATE_VERSION = "v1.0.1"
