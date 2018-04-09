///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package helpers

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
	"github.com/vmware/dispatch/pkg/events"
)

// NO TESTS

// CloudEventFromSwagger creates CloudEvent struct from Swagger model
func CloudEventFromSwagger(e *models.CloudEvent) *events.CloudEvent {
	if e == nil {
		return nil
	}
	return &events.CloudEvent{
		EventType:          *e.EventType,
		EventTypeVersion:   e.EventTypeVersion,
		CloudEventsVersion: *e.CloudEventsVersion,
		Source:             *e.Source,
		EventID:            *e.EventID,
		EventTime:          time.Time(e.EventTime),
		SchemaURL:          e.SchemaURL,
		ContentType:        e.ContentType,
		Extensions:         events.CloudEventExtensions(e.Extensions),
		Data:               e.Data,
	}
}

// CloudEventToSwagger creates Swagger model from CloudEvent struct
func CloudEventToSwagger(e *events.CloudEvent) *models.CloudEvent {
	if e == nil {
		return nil
	}
	return &models.CloudEvent{
		CloudEventsVersion: swag.String(e.CloudEventsVersion),
		ContentType:        e.ContentType,
		Data:               e.Data,
		EventID:            swag.String(e.EventID),
		EventTime:          strfmt.DateTime(e.EventTime),
		EventType:          swag.String(e.EventType),
		EventTypeVersion:   e.EventTypeVersion,
		Extensions:         e.Extensions,
		SchemaURL:          e.SchemaURL,
		Source:             swag.String(e.Source),
	}
}
