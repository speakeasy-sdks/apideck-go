// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/utils"
)

type PassThroughQuery struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// All passthrough query parameters are passed along to the connector as is (?pass_through[search]=leads becomes ?search=leads)
	ExampleDownstreamProperty *string `queryParam:"name=example_downstream_property"`
}

func (p PassThroughQuery) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PassThroughQuery) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PassThroughQuery) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *PassThroughQuery) GetExampleDownstreamProperty() *string {
	if o == nil {
		return nil
	}
	return o.ExampleDownstreamProperty
}