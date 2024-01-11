package discovery

type Instance interface {
	Network() string
	Address() string
	Port() int
	Weight() int
	Tag(key string) (value string, exist bool)
}

func NewInstance(network, address string, port int, weight int, tags map[string]string) Instance {
	return &instance{
		network: network,
		addr:    address,
		port:    port,
		weight:  weight,
		tags:    tags,
	}
}

type instance struct {
	network string
	addr    string
	port    int
	weight  int
	tags    map[string]string
}

func (i *instance) Network() string {
	return i.network
}

func (i *instance) Address() string {
	return i.addr
}

func (i *instance) Port() int {
	return i.port
}

func (i *instance) Weight() int {
	return i.weight
}

func (i *instance) Tag(key string) (value string, exist bool) {
	value, exist = i.tags[key]
	return
}
