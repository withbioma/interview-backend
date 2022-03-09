package contexthelper

import (
	"context"
	"reflect"
)

type ContextKey uint

const (
	// keep-sorted

	Appointment        = ContextKey(iota)
	Control            = ContextKey(iota)
	CurrentUser        = ContextKey(iota)
	CurrentUserID      = ContextKey(iota)
	Institution        = ContextKey(iota)
	Invitation         = ContextKey(iota)
	Patient            = ContextKey(iota)
	CurrentPatientID   = ContextKey(iota)
	CurrentPatient     = ContextKey(iota)
	CurrentPatients    = ContextKey(iota)
	RecurringTimeSlot  = ContextKey(iota)
	TimeSlot           = ContextKey(iota)
	TimeSlotQuota      = ContextKey(iota)
	Symptom            = ContextKey(iota)
	Vaccine            = ContextKey(iota)
	VaccinationAttempt = ContextKey(iota)
	VaccineEligibility = ContextKey(iota)
)

// WithValue wraps context with an arbitrary, non-model value.
func WithValue(
	ctx context.Context,
	key ContextKey,
	value interface{},
) context.Context {
	return context.WithValue(ctx, key, value)
}

// WithModelValue wraps context with model value.
// @param ctx original context.
// @param key the context key found above.
// @param model the model interface in the form of a model pointer.
func WithModelValue(
	ctx context.Context,
	key ContextKey,
	model interface{},
) context.Context {
	modelValue := reflect.ValueOf(model)

	if modelValue.Kind() == reflect.Ptr {
		return WithModelValue(ctx, key, modelValue.Elem().Interface())
	} else if modelValue.Kind() == reflect.Struct {
		return context.WithValue(ctx, key, convertToPointer(model))
	}

	panic("model passed is not a struct or pointer to a struct")
}

func convertToPointer(v interface{}) interface{} {
	p := reflect.New(reflect.TypeOf(v))
	p.Elem().Set(reflect.ValueOf(v))

	return p.Interface()
}
