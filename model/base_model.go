package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type BaseModel struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
}

func (this BaseModel) IsPersisted() bool {
	return this.ID > 0
}

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	} else {
		return json.Marshal(nil)
	}
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}

	return nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var r *float64
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	if r != nil {
		nf.Valid = true
		nf.Float64 = *r
	} else {
		nf.Valid = false
	}

	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (ni *NullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var r *int64
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	if r != nil {
		ni.Valid = true
		ni.Int64 = *r
	} else {
		ni.Valid = false
	}

	return nil
}
