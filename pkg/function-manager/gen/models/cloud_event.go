///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudEvent cloud event
// swagger:model CloudEvent
type CloudEvent struct {

	// cloud events version
	// Required: true
	CloudEventsVersion *string `json:"cloudEventsVersion"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// data
	// Max Length: 0
	Data string `json:"data,omitempty"`

	// event ID
	// Required: true
	EventID *string `json:"eventID"`

	// event time
	EventTime strfmt.DateTime `json:"eventTime,omitempty"`

	// event type
	// Required: true
	// Max Length: 128
	// Pattern: ^[\w\d\-\.]+$
	EventType *string `json:"eventType"`

	// event type version
	EventTypeVersion string `json:"eventTypeVersion,omitempty"`

	// extensions
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	// schema URL
	SchemaURL string `json:"schemaURL,omitempty"`

	// source
	// Required: true
	Source *string `json:"source"`
}

// Validate validates this cloud event
func (m *CloudEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudEventsVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudEvent) validateCloudEventsVersion(formats strfmt.Registry) error {

	if err := validate.Required("cloudEventsVersion", "body", m.CloudEventsVersion); err != nil {
		return err
	}

	return nil
}

func (m *CloudEvent) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if err := validate.MaxLength("data", "body", string(m.Data), 0); err != nil {
		return err
	}

	return nil
}

func (m *CloudEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventID", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *CloudEvent) validateEventTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudEvent) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("eventType", "body", m.EventType); err != nil {
		return err
	}

	if err := validate.MaxLength("eventType", "body", string(*m.EventType), 128); err != nil {
		return err
	}

	if err := validate.Pattern("eventType", "body", string(*m.EventType), `^[\w\d\-\.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *CloudEvent) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudEvent) UnmarshalBinary(b []byte) error {
	var res CloudEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
