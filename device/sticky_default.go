//go:build !linux

package device

import (
	"github.com/akhokhlow/obfs-wireguard-go/conn"
	"github.com/akhokhlow/obfs-wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
