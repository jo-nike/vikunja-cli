package cmdutil

import (
	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/config"
	"github.com/jo-nike/vikunja-cli/internal/jwtcache"
	"github.com/jo-nike/vikunja-cli/internal/output"
)

// MustClient creates a client from config, exiting on error.
// When username/password are configured, it authenticates with a JWT
// (cached to disk) instead of the API token, since JWTs are required
// for write operations to persist.
func MustClient() *client.Client {
	cfg, err := config.Load()
	if err != nil {
		output.Error(err)
	}
	c, err := client.New(cfg)
	if err != nil {
		output.Error(err)
	}

	if cfg.Username != "" && cfg.Password != "" {
		if cached, _ := jwtcache.LoadCachedJWT(cfg.URL); cached != "" {
			c.SetToken(cached)
		} else {
			jwt, err := c.Login(cfg.Username, cfg.Password)
			if err != nil {
				output.Error(err)
			}
			_ = jwtcache.CacheJWT(cfg.URL, jwt)
			c.SetToken(jwt)
		}
	}

	return c
}

// AddPaginationFlags adds --page, --per-page, and --all flags to a command.
func PaginationOpts(page, perPage int) []client.RequestOption {
	var opts []client.RequestOption
	if page > 0 {
		opts = append(opts, client.WithPage(page))
	}
	if perPage > 0 {
		opts = append(opts, client.WithPerPage(perPage))
	}
	return opts
}
