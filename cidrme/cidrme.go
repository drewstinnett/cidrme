package cidrme

import (
	"bytes"
	"fmt"
	"net"
	"net/netip"
	"strings"
)

type CIDR struct {
	CIDR   string
	Prefix netip.Prefix
}

type CIDRSummary struct {
	CIDR   string
	Prefix int
	Mask   string
	Count  int64
}

func NewCIDR(c string) (*CIDR, error) {
	if !strings.Contains(c, "/") {
		c = c + "/32"
	} else {
		parts := strings.Split(c, "/")
		ip := parts[0]
		maybeMask := parts[1]
		if strings.Contains(maybeMask, ".") {
			mask := net.IPMask(net.ParseIP(maybeMask).To4())
			length, _ := mask.Size()
			c = fmt.Sprintf("%v/%v", ip, length)
		}
	}

	cidr := CIDR{
		CIDR: c,
	}
	var err error
	cidr.Prefix, err = netip.ParsePrefix(c)
	if err != nil {
		return nil, err
	}
	return &cidr, nil
}

func (c *CIDR) PrefixLength() int {
	return c.Prefix.Masked().Bits()
}

func (c *CIDR) Mask() string {
	return localMask(c.Prefix.Bits())
}

func (c *CIDR) Count() int64 {
	var count int64
	for addr := c.Prefix.Addr(); c.Prefix.Contains(addr); addr = addr.Next() {
		count++
	}
	if count < 2 {
		return 1
	}
	return count - 2
}

func (c *CIDR) Summarize() *CIDRSummary {
	return &CIDRSummary{
		CIDR:   c.CIDR,
		Prefix: c.PrefixLength(),
		Mask:   c.Mask(),
		Count:  c.Count(),
	}
}

// GetIPs returns a list of all ips in a given CIDR as strings
func (c *CIDR) GetIPs() []string {
	r := []string{}
	for addr := c.Prefix.Addr(); c.Prefix.Contains(addr); addr = addr.Next() {
		r = append(r, addr.String())
	}
	return r
}

// GetUsableIPs returns a list of all ips in a given CIDR as strings, excluding the first and last addresses
func (c *CIDR) GetUsableIPs() []string {
	r := []string{}
	for addr := c.Prefix.Addr(); c.Prefix.Contains(addr); addr = addr.Next() {
		r = append(r, addr.String())
	}
	return r[1 : len(r)-1]
}

func localMask(p int) string {
	mask := (0xFFFFFFFF << (32 - p)) & 0xFFFFFFFF
	var dmask uint64
	dmask = 32
	localmask := make([]uint64, 0, 4)
	var tmp uint64
	for i := 1; i <= 4; i++ {
		tmp = uint64(mask) >> (dmask - 8) & 0xFF
		localmask = append(localmask, tmp)
		dmask -= 8
	}

	return joinMask(localmask)
}

func joinMask(a []uint64) string {
	var buffer bytes.Buffer
	for i := 0; i < len(a); i++ {
		buffer.WriteString(fmt.Sprintf("%v", a[i]))
		if i != len(a)-1 {
			buffer.WriteString(".")
		}
	}

	return buffer.String()
}
