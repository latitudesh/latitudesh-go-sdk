// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type ServerExitRescueModeRequest struct {
	ServerID string `pathParam:"style=simple,explode=false,name=server_id"`
}

func (o *ServerExitRescueModeRequest) GetServerID() string {
	if o == nil {
		return ""
	}
	return o.ServerID
}

type ServerExitRescueModeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	ServerRescue *components.ServerRescue
}

func (o *ServerExitRescueModeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ServerExitRescueModeResponse) GetServerRescue() *components.ServerRescue {
	if o == nil {
		return nil
	}
	return o.ServerRescue
}
