package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PluginInterfaceType plugin interface type
// swagger:model PluginInterfaceType
type PluginInterfaceType struct {

	// capability
	// Required: true
	Capability string `json:"Capability"`

	// prefix
	// Required: true
	Prefix string `json:"Prefix"`

	// version
	// Required: true
	Version string `json:"Version"`
}
