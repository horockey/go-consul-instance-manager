// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.7
// Revision: bf63e108589bbd2327b13ec2c5da532aad234029
// Build Date: 2023-07-25T23:27:55Z
// Built By: goreleaser

package consul_instance_manager

import (
	"errors"
	"fmt"
)

const (
	// InstanceStatusAlive is a InstanceStatus of type Alive.
	InstanceStatusAlive InstanceStatus = iota
	// InstanceStatusPending is a InstanceStatus of type Pending.
	InstanceStatusPending
)

var ErrInvalidInstanceStatus = errors.New("not a valid InstanceStatus")

const _InstanceStatusName = "alivepending"

var _InstanceStatusMap = map[InstanceStatus]string{
	InstanceStatusAlive:   _InstanceStatusName[0:5],
	InstanceStatusPending: _InstanceStatusName[5:12],
}

// String implements the Stringer interface.
func (x InstanceStatus) String() string {
	if str, ok := _InstanceStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("InstanceStatus(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x InstanceStatus) IsValid() bool {
	_, ok := _InstanceStatusMap[x]
	return ok
}

var _InstanceStatusValue = map[string]InstanceStatus{
	_InstanceStatusName[0:5]:  InstanceStatusAlive,
	_InstanceStatusName[5:12]: InstanceStatusPending,
}

// ParseInstanceStatus attempts to convert a string to a InstanceStatus.
func ParseInstanceStatus(name string) (InstanceStatus, error) {
	if x, ok := _InstanceStatusValue[name]; ok {
		return x, nil
	}
	return InstanceStatus(0), fmt.Errorf("%s is %w", name, ErrInvalidInstanceStatus)
}
