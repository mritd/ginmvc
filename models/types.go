package models

import (
	"database/sql"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type NullInt32 struct {
	sql.NullInt32
}

func (v NullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return jsoniter.Marshal(v.Int32)
	} else {
		return jsoniter.Marshal(0)
	}
}

func (v *NullInt32) UnmarshalJSON(data []byte) error {
	var x *int32
	if err := jsoniter.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int32 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return jsoniter.Marshal(v.Int64)
	} else {
		return jsoniter.Marshal(nil)
	}
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := jsoniter.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return jsoniter.Marshal(v.String)
	} else {
		return jsoniter.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := jsoniter.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullBool struct {
	sql.NullBool
}

func (v NullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return jsoniter.Marshal(v.Bool)
	} else {
		return jsoniter.Marshal(nil)
	}
}

func (v *NullBool) UnmarshalJSON(data []byte) error {
	var x *bool
	if err := jsoniter.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Bool = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return jsoniter.Marshal(v.Float64)
	} else {
		return jsoniter.Marshal(nil)
	}
}

func (v *NullFloat64) UnmarshalJSON(data []byte) error {
	var x *float64
	if err := jsoniter.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullTime struct {
	sql.NullTime
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return jsoniter.Marshal(v.NullTime)
	} else {
		return jsoniter.Marshal(nil)
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	var x *time.Time
	if err := jsoniter.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}
	return nil
}
