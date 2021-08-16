// +build !linux

package device

import (
	"github.com/Skyxim/wireguard-go/conn"
	"github.com/Skyxim/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
