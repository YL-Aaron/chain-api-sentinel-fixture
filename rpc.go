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

func (api *PublicAPI) BlockNumber(number string) (string, error) {
	return "0x1", nil
}
func (api *PublicAPI) ChainID(number string) (string, error) {
	return "0x1", nil
}
