package discovery

type Result struct {
	Cacheable bool
	CacheKey string
	Instances []Instance
}
