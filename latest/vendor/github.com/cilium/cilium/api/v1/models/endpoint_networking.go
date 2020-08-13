// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointNetworking Unique identifiers for this endpoint from outside cilium
//
// swagger:model EndpointNetworking
type EndpointNetworking struct {

	// IP4/6 addresses assigned to this Endpoint
	Addressing []*AddressPair `json:"addressing"`

	// host addressing
	HostAddressing *NodeAddressing `json:"host-addressing,omitempty"`

	// MAC address
	HostMac string `json:"host-mac,omitempty"`

	// Index of network device
	InterfaceIndex int64 `json:"interface-index,omitempty"`

	// Name of network device
	InterfaceName string `json:"interface-name,omitempty"`

	// MAC address
	Mac string `json:"mac,omitempty"`
}

// Validate validates this endpoint networking
func (m *EndpointNetworking) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAddressing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointNetworking) validateAddressing(formats strfmt.Registry) error {

	if swag.IsZero(m.Addressing) { // not required
		return nil
	}

	for i := 0; i < len(m.Addressing); i++ {
		if swag.IsZero(m.Addressing[i]) { // not required
			continue
		}

		if m.Addressing[i] != nil {
			if err := m.Addressing[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addressing" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EndpointNetworking) validateHostAddressing(formats strfmt.Registry) error {

	if swag.IsZero(m.HostAddressing) { // not required
		return nil
	}

	if m.HostAddressing != nil {
		if err := m.HostAddressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-addressing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointNetworking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointNetworking) UnmarshalBinary(b []byte) error {
	var res EndpointNetworking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
