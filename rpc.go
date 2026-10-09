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
	if number == "latest" {
		return "0x2", nil
	}
	return "0x1", nil
}
