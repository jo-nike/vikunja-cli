#!/usr/bin/env bash
# Assemble a self-contained vikunja-cli skill bundle: the skill source plus a
# matching binary, zipped into build/ for attachment to the release. Invoked by
# GoReleaser's before-hook (see .goreleaser.yaml) with the release version as $1.
#
# Usage: scripts/build-skill-bundle.sh <version> [goos] [goarch]
set -euo pipefail

VERSION="${1:-dev}"
GOOS_BUNDLE="${2:-darwin}"
GOARCH_BUNDLE="${3:-arm64}"

PKG="github.com/jo-nike/vikunja-cli"
SKILL_DIR=".claude/skills/vikunja-cli"
COMMIT="$(git rev-parse --short HEAD 2>/dev/null || echo none)"
DATE="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
# Stage outside dist/: GoReleaser's --clean wipes dist/ and then refuses to run if
# the before-hook has repopulated it. release.extra_files globs this path instead.
OUT="build/vikunja-cli-skill-${GOOS_BUNDLE}-${GOARCH_BUNDLE}.zip"

mkdir -p build "${SKILL_DIR}/bin"

echo "building skill binary ${GOOS_BUNDLE}/${GOARCH_BUNDLE} (version ${VERSION})"
GOOS="${GOOS_BUNDLE}" GOARCH="${GOARCH_BUNDLE}" CGO_ENABLED=0 \
	go build -ldflags "-s -w -X ${PKG}/cmd.Version=${VERSION} -X ${PKG}/cmd.Commit=${COMMIT} -X ${PKG}/cmd.Date=${DATE}" \
	-o "${SKILL_DIR}/bin/vikunja-cli" .

echo "zipping skill bundle -> ${OUT}"
rm -f "${OUT}"
# Zip with the vikunja-cli/ prefix preserved so it extracts as a ready-to-use skill dir.
( cd "$(dirname "${SKILL_DIR}")" && zip -q -r -X "${OLDPWD}/${OUT}" "$(basename "${SKILL_DIR}")" )

echo "done: ${OUT}"
