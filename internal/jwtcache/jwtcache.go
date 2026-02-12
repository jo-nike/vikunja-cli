package jwtcache

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LoadCachedJWT reads a cached JWT for the given server URL.
// Returns the token if valid, or empty string if missing/expired.
func LoadCachedJWT(serverURL string) (string, error) {
	path := cachePath(serverURL)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", nil
	}

	exp, err := parseJWTExp(token)
	if err != nil {
		return "", err
	}

	// Require at least 30 seconds of remaining validity
	if time.Now().Unix()+30 >= exp {
		return "", nil
	}

	return token, nil
}

// CacheJWT writes a JWT to the cache file for the given server URL.
func CacheJWT(serverURL, token string) error {
	dir := cacheDir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("creating cache dir: %w", err)
	}
	path := cachePath(serverURL)
	return os.WriteFile(path, []byte(token), 0600)
}

// parseJWTExp extracts the exp claim from a JWT by base64-decoding the payload.
func parseJWTExp(token string) (int64, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid JWT: expected 3 parts, got %d", len(parts))
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, fmt.Errorf("decoding JWT payload: %w", err)
	}

	var claims struct {
		Exp int64 `json:"exp"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return 0, fmt.Errorf("parsing JWT claims: %w", err)
	}

	if claims.Exp == 0 {
		return 0, fmt.Errorf("JWT has no exp claim")
	}

	return claims.Exp, nil
}

func cacheDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		dir = filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return filepath.Join(dir, "vikunja-cli")
}

func cachePath(serverURL string) string {
	h := sha256.Sum256([]byte(serverURL))
	return filepath.Join(cacheDir(), fmt.Sprintf("jwt-%x", h[:8]))
}
