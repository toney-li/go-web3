package admin

import (
	"github.com/toney-li/go-web3/providers"
	"github.com/toney-li/go-web3/dto"
)

// Admin - The Admin Module
type Admin struct {
	provider providers.ProviderInterface
}

// NewAdmin - Admin Module constructor to set the default provider
func NewAdmin(provider providers.ProviderInterface) *Admin {
	personal := new(Admin)
	personal.provider = provider
	return personal
}

func (admin *Admin) DataDir() (string, error) {
	pointer := &dto.RequestResult{}
	err := admin.provider.SendRequest(pointer, "admin_datadir", nil)

	if err != nil {
		return "", err
	}
	return pointer.ToString()
}
