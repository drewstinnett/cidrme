package cidrme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCIDR(t *testing.T) {
	tests := map[string]struct {
		cidr    string
		wantErr bool
	}{
		"plain-network": {cidr: "192.168.1.0/24"},
		"plain-ip":      {cidr: "192.168.1.1"},
		"not-cidr":      {cidr: "foo", wantErr: true},
	}
	for k, tt := range tests {
		got, err := NewCIDR(tt.cidr)
		if tt.wantErr {
			require.Error(t, err, k)
			require.Nil(t, got, k)
		} else {
			require.NoError(t, err, k)
			require.NotNil(t, got, k)
		}
	}
}

func TestCIDRPrefix(t *testing.T) {
	tests := map[string]struct {
		cidr string
		want int
	}{
		"plain-network":     {cidr: "192.168.1.0/24", want: 24},
		"network-with-mask": {cidr: "192.168.1.0/255.255.255.0", want: 24},
		"plain-ip":          {cidr: "192.168.1.1", want: 32},
	}
	for k, tt := range tests {
		c, err := NewCIDR(tt.cidr)
		require.NoError(t, err, k)
		got := c.PrefixLength()
		require.Equal(t, tt.want, got, k)
	}
}

func TestCIDRMask(t *testing.T) {
	tests := map[string]struct {
		cidr string
		want string
	}{
		"plain-network": {cidr: "192.168.1.0/24", want: "255.255.255.0"},
		"plain-ip":      {cidr: "192.168.1.1", want: "255.255.255.255"},
	}
	for k, tt := range tests {
		c, err := NewCIDR(tt.cidr)
		require.NoError(t, err, k)
		got := c.Mask()
		require.Equal(t, tt.want, got, k)
	}
}

func TestSummarize(t *testing.T) {
	c, err := NewCIDR("192.168.1.0/24")
	require.NoError(t, err)
	got := c.Summarize()
	require.NotNil(t, got)
	require.Equal(t, 24, got.Prefix)
	require.Equal(t, "255.255.255.0", got.Mask)
	require.Equal(t, int64(254), got.Count)
}

func TestSummarizeSingle(t *testing.T) {
	c, err := NewCIDR("192.168.1.5/32")
	require.NoError(t, err)
	got := c.Summarize()
	require.NotNil(t, got)
	require.Equal(t, 32, got.Prefix)
	require.Equal(t, "255.255.255.255", got.Mask)
	require.Equal(t, int64(1), got.Count)
}

func TestGetIPs(t *testing.T) {
	c, err := NewCIDR("192.168.1.0/24")
	require.NoError(t, err)
	require.Equal(t, 256, len(c.GetIPs()))
}

func TestGetUsableIPs(t *testing.T) {
	c, err := NewCIDR("192.168.1.0/24")
	require.NoError(t, err)
	got := c.GetUsableIPs()
	require.Equal(t, 254, len(got))
}
