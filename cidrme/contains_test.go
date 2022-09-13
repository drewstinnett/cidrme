package cidrme

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	tests := map[string]struct {
		network string
		ip      string
		want    bool
	}{
		"plain-in": {
			network: "192.168.1.0/24",
			ip:      "192.168.1.5",
			want:    true,
		},
		"plain-out": {
			network: "192.168.1.0/24",
			ip:      "192.168.2.5",
			want:    false,
		},
	}
	for _, tt := range tests {
		n := netip.MustParsePrefix(tt.network)
		a := netip.MustParseAddr(tt.ip)
		got := Contains(n, a)
		require.Equal(t, tt.want, got.Result)
	}

	for _, tt := range tests {
		got, err := ContainsWithStrings(tt.network, tt.ip)
		require.NoError(t, err)
		require.Equal(t, tt.want, got.Result)
	}
}
