package digitalocean

import (
	"fmt"
	"github.com/kris-nova/kubicorn/vpn"
)

type OpenVPN struct {
	Name string `json:"name,omitempty"`
}

func (o OpenVPN) Test(){
	fmt.Println("AAAAAAAAAA")
}

func NewOpenVPN() vpn.VPN {
	return OpenVPN{
		Name: "Abc",
	}
}