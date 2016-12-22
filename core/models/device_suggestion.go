package models




import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// DeviceSuggestion device suggestion
// swagger:model DeviceSuggestion
type DeviceSuggestion struct {

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// place suggestion
	PlaceSuggestion []*PlaceSuggestion `json:"placeSuggestion"`
}

// Validate validates this device suggestion
func (m *DeviceSuggestion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlaceSuggestion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceSuggestion) validatePlaceSuggestion(formats strfmt.Registry) error {

	if swag.IsZero(m.PlaceSuggestion) { // not required
		return nil
	}

	for i := 0; i < len(m.PlaceSuggestion); i++ {

		if swag.IsZero(m.PlaceSuggestion[i]) { // not required
			continue
		}

		if m.PlaceSuggestion[i] != nil {

			if err := m.PlaceSuggestion[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}