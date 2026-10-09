package fixture

type API struct {
	Namespace     string
	Service       any
	Authenticated bool
}

type PublicAPI struct{}

var APIs = []API{
	{
		Namespace: "eth",
		Service:   &PublicAPI{},
	},
}

func (api *PublicAPI) ChainID() (string, error) {
	return "0x1", nil
}
