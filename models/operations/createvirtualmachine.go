// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type CreateVirtualMachineResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Created
	VirtualMachine *components.VirtualMachine
}

func (o *CreateVirtualMachineResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateVirtualMachineResponse) GetVirtualMachine() *components.VirtualMachine {
	if o == nil {
		return nil
	}
	return o.VirtualMachine
}
