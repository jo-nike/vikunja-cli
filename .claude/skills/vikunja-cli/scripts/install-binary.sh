#!/usr/bin/env bash
# Fetch the `vikunja-cli` binary for this machine into the skill's bin/ directory.
# The skill ships instructions but not the binary (platform-specific + large),
# so this runs once on first use. Downloads the matching build from the GitHub
# release, verifies its SHA256 against checksums.txt, and installs it.
#
# Usage: bash install-binary.sh [version]   (version defaults to "latest")
#   Override the repo with VIKUNJA_CLI_REPO=owner/repo if you forked it.
set -euo pipefail

REPO="${VIKUNJA_CLI_REPO:-jo-nike/vikunja-cli}"
VERSION="${1:-latest}"

# Resolve the skill dir (this script lives in <skill>/scripts/) and target bin/.
SKILL_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BIN_DIR="${SKILL_DIR}/bin"

# --- detect platform -------------------------------------------------------
os="$(uname -s)"; arch="$(uname -m)"
case "$os" in
	Darwin) goos="darwin" ;;
	Linux)  goos="linux" ;;
	MINGW*|MSYS*|CYGWIN*) goos="windows" ;;
	*) echo "unsupported OS: $os" >&2; exit 1 ;;
esac
case "$arch" in
	arm64|aarch64) goarch="arm64" ;;
	x86_64|amd64)  goarch="amd64" ;;
	*) echo "unsupported arch: $arch" >&2; exit 1 ;;
esac
if [ "$goos" = "windows" ] && [ "$goarch" = "arm64" ]; then
	echo "no windows/arm64 build is published" >&2; exit 1
fi

if [ "$goos" = "windows" ]; then ext="zip"; binname="vikunja-cli.exe"; else ext="tar.gz"; binname="vikunja-cli"; fi
asset="vikunja-cli-${goos}-${goarch}.${ext}"

if [ "$VERSION" = "latest" ]; then
	base="https://github.com/${REPO}/releases/latest/download"
else
	base="https://github.com/${REPO}/releases/download/${VERSION}"
fi

# --- pick downloader / hasher ----------------------------------------------
if command -v curl >/dev/null 2>&1; then dl() { curl -fsSL "$1" -o "$2"; }
elif command -v wget >/dev/null 2>&1; then dl() { wget -qO "$2" "$1"; }
else echo "need curl or wget to download the binary" >&2; exit 1; fi

if command -v sha256sum >/dev/null 2>&1; then sha() { sha256sum "$1" | awk '{print $1}'; }
elif command -v shasum >/dev/null 2>&1; then sha() { shasum -a 256 "$1" | awk '{print $1}'; }
else echo "need sha256sum or shasum to verify the download" >&2; exit 1; fi

# --- download + verify + install -------------------------------------------
tmp="$(mktemp -d)"; trap 'rm -rf "$tmp"' EXIT

echo "downloading ${asset} from ${REPO} (${VERSION})..."
dl "${base}/${asset}" "${tmp}/${asset}"
dl "${base}/checksums.txt" "${tmp}/checksums.txt"

want="$(awk -v a="$asset" '$2==a {print $1}' "${tmp}/checksums.txt")"
if [ -z "$want" ]; then echo "no checksum for ${asset} in checksums.txt" >&2; exit 1; fi
got="$(sha "${tmp}/${asset}")"
if [ "$want" != "$got" ]; then
	echo "checksum mismatch for ${asset}" >&2
	echo "  expected $want" >&2; echo "  got      $got" >&2
	exit 1
fi
echo "checksum verified."

case "$ext" in
	tar.gz) tar -xzf "${tmp}/${asset}" -C "$tmp" ;;
	zip)    unzip -qo "${tmp}/${asset}" -d "$tmp" ;;
esac

mkdir -p "$BIN_DIR"
install -m 0755 "${tmp}/${binname}" "${BIN_DIR}/${binname}" 2>/dev/null \
	|| { cp "${tmp}/${binname}" "${BIN_DIR}/${binname}"; chmod 0755 "${BIN_DIR}/${binname}"; }

echo "installed ${BIN_DIR}/${binname}"
"${BIN_DIR}/${binname}" version || true
