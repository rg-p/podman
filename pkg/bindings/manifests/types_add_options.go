// Code generated by go generate; DO NOT EDIT.
package manifests

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *AddOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *AddOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAll set field All to given value
func (o *AddOptions) WithAll(value bool) *AddOptions {
	o.All = &value
	return o
}

// GetAll returns value of field All
func (o *AddOptions) GetAll() bool {
	if o.All == nil {
		var z bool
		return z
	}
	return *o.All
}

// WithAnnotation set field Annotation to given value
func (o *AddOptions) WithAnnotation(value map[string]string) *AddOptions {
	o.Annotation = value
	return o
}

// GetAnnotation returns value of field Annotation
func (o *AddOptions) GetAnnotation() map[string]string {
	if o.Annotation == nil {
		var z map[string]string
		return z
	}
	return o.Annotation
}

// WithArch set field Arch to given value
func (o *AddOptions) WithArch(value string) *AddOptions {
	o.Arch = &value
	return o
}

// GetArch returns value of field Arch
func (o *AddOptions) GetArch() string {
	if o.Arch == nil {
		var z string
		return z
	}
	return *o.Arch
}

// WithFeatures set field Features to given value
func (o *AddOptions) WithFeatures(value []string) *AddOptions {
	o.Features = value
	return o
}

// GetFeatures returns value of field Features
func (o *AddOptions) GetFeatures() []string {
	if o.Features == nil {
		var z []string
		return z
	}
	return o.Features
}

// WithImages set field Images to given value
func (o *AddOptions) WithImages(value []string) *AddOptions {
	o.Images = value
	return o
}

// GetImages returns value of field Images
func (o *AddOptions) GetImages() []string {
	if o.Images == nil {
		var z []string
		return z
	}
	return o.Images
}

// WithOS set field OS to given value
func (o *AddOptions) WithOS(value string) *AddOptions {
	o.OS = &value
	return o
}

// GetOS returns value of field OS
func (o *AddOptions) GetOS() string {
	if o.OS == nil {
		var z string
		return z
	}
	return *o.OS
}

// WithOSVersion set field OSVersion to given value
func (o *AddOptions) WithOSVersion(value string) *AddOptions {
	o.OSVersion = &value
	return o
}

// GetOSVersion returns value of field OSVersion
func (o *AddOptions) GetOSVersion() string {
	if o.OSVersion == nil {
		var z string
		return z
	}
	return *o.OSVersion
}

// WithVariant set field Variant to given value
func (o *AddOptions) WithVariant(value string) *AddOptions {
	o.Variant = &value
	return o
}

// GetVariant returns value of field Variant
func (o *AddOptions) GetVariant() string {
	if o.Variant == nil {
		var z string
		return z
	}
	return *o.Variant
}
