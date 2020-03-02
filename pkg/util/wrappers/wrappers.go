package wrappers

import (
	"time"

	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

// NullString wraps sql.NullString and implements Marshalling for serialization / deserialization
type NullString struct {
	sql.NullString
}

// NewNullString constructor for NullTime
func NewNullString(str *string) *NullString {

	var nullString sql.NullString

	if str == nil {
		nullString = sql.NullString{
			String: "",
			Valid:  false}
	} else {
		nullString = sql.NullString{
			String: *str,
			Valid:  true}
	}

	return &NullString{nullString}
}

// NullInt64 wraps sql.NullInt64 and implements Marshalling for serialization / deserialization
type NullInt64 struct {
	sql.NullInt64
}

// NullTime wraps sql.NullInt64 and implements Marshalling for serialization / deserialization
type NullTime struct {
	sql.NullTime
}

// NewNullTime constructor for NullTime
func NewNullTime(tim *time.Time) *NullTime {
	var nullTime sql.NullTime

	if tim == nil {
		nullTime = sql.NullTime{
			Time:  time.Now(),
			Valid: false}
	} else {
		nullTime = sql.NullTime{
			Time:  *tim,
			Valid: true}
	}

	return &NullTime{nullTime}
}

// NullUUID wraps uuid.UUID and implements Marshalling
type NullUUID struct {
	id uuid.UUID
	valid bool
}

// MarshalJSON ...
func (nullString NullString) MarshalJSON() ([]byte, error) {
	if nullString.Valid {
		return json.Marshal(nullString.String)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (nullString *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var String *string
	if err := json.Unmarshal(data, &String); err != nil {
		return err
	}
	if String != nil {
		nullString.Valid = true
		nullString.String = *String
	} else {
		nullString.Valid = false
	}
	return nil
}

// MarshalJSON ...
func (nullInt64 NullInt64) MarshalJSON() ([]byte, error) {
	if nullInt64.Valid {
		return json.Marshal(nullInt64.NullInt64)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (nullInt64 *NullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var Int64 *int64
	if err := json.Unmarshal(data, &Int64); err != nil {
		return err
	}
	if Int64 != nil {
		nullInt64.Int64 = *Int64
		nullInt64.Valid = true
	} else {
		nullInt64.Valid = false
	}
	return nil
}

// MarshalJSON ...
func (nullTime NullTime) MarshalJSON() ([]byte, error) {
	if nullTime.Valid {
		return json.Marshal(nullTime.NullTime)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (nullTime *NullTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var Time *time.Time
	if err := json.Unmarshal(data, &Time); err != nil {
		return err
	}
	if Time != nil {
		nullTime.Time = *Time
		nullTime.Valid = true
	} else {
		nullTime.Valid = false
	}
	return nil
}

// MarshalJSON ...
func (nullUUID NullUUID) MarshalJSON() ([]byte, error) {
	if nullUUID.valid {
		return json.Marshal(nullUUID.id)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (nullUUID *NullUUID) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var id *uuid.UUID
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	if id != nil {
		nullUUID.id = *id
		nullUUID.valid = true
	} else {
		nullUUID.valid = false
	}
	return nil
}
