// +build !linux

package device

import (
	"github.com/Skyxim/wireguard/conn"
	"github.com/Skyxim/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
