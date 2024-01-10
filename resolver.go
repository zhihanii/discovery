package discovery

import "context"

type Resolver interface {
	Target(ctx context.Context) string
	Resolve(ctx context.Context, target string) (Result, error)
	Close()
}
