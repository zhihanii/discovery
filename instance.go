package discovery

type Instance interface {
	Network() string
	Address() string
	Weight() int
	Tag(key string) (value string, exist bool)
}

func NewInstance(network, address string, weight int, tags map[string]string) Instance {
	return &instance{
		network: network,
		addr: address,
		weight: weight,
		tags: tags,
	}
}

type instance struct {
	network string
	addr string
	weight int
	tags map[string]string
}

func (i *instance) Network() string {
	return i.network
}

func (i *instance) Address() string {
	return i.addr
}

func (i *instance) Weight() int {
	return i.weight
}

func (i *instance) Tag(key string) (value string, exist bool) {
	value, exist = i.tags[key]
	return
}
