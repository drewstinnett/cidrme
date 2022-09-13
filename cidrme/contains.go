package cidrme

import (
	"fmt"
	"net/netip"
)

type ContainsResult struct {
	Result  bool   `yaml:"result"`
	Network string `yaml:"network"`
	Address string `yaml:"address"`
	Message string `yaml:"message"`
}

func Contains(n netip.Prefix, a netip.Addr) *ContainsResult {
	r := &ContainsResult{
		Result:  n.Contains(a),
		Network: n.String(),
		Address: a.String(),
	}
	if r.Result {
		r.Message = fmt.Sprintf("%v is part of %v", r.Address, r.Network)
	} else {
		r.Message = fmt.Sprintf("%v is NOT part of %v", r.Address, r.Network)
	}
	return r
}

func ContainsWithStrings(n, a string) (*ContainsResult, error) {
	nn, err := NewCIDR(n)
	if err != nil {
		return nil, err
	}
	aa, err := NewCIDR(a)
	if err != nil {
		return nil, err
	}
	return Contains(nn.Prefix, aa.Prefix.Addr()), nil
}
