package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSONB struct {
	Bytes []byte
	Valid bool
}

func (dst *JSONB) Set(src interface{}) error {
	if src == nil {
		*dst = JSONB{Valid: false}
		return nil
	}

	switch value := src.(type) {
	case string:
		*dst = JSONB{Valid: true, Bytes: []byte(value)}
	case *string:
		if value == nil {
			*dst = JSONB{Valid: false}
		} else {
			*dst = JSONB{Valid: true, Bytes: []byte(*value)}
		}
	case []byte:
		if value == nil {
			*dst = JSONB{Valid: false}
		} else {
			*dst = JSONB{Valid: true, Bytes: value}
		}
	default:
		buf, err := json.Marshal(src)
		if err != nil {
			return err
		}
		*dst = JSONB{Valid: true, Bytes: buf}
	}

	return nil
}

func (dst *JSONB) Get() interface{} {
	if dst.Valid {
		var i interface{}
		err := json.Unmarshal(dst.Bytes, &i)
		if err != nil {
			return dst
		}

		return i
	}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *JSONB) Scan(src interface{}) error {
	if src == nil {
		*dst = JSONB{Valid: false}
		return nil
	}

	switch src := src.(type) {
	case string:
		*dst = JSONB{Valid: true, Bytes: []byte(src)}
		return nil
	case []byte:
		*dst = JSONB{Valid: true, Bytes: src}
		return nil
	}

	return errors.New(fmt.Sprintf("cannot scan %T", src))
}

// // Value implements the database/sql/driver Valuer interface.
func (src JSONB) Value() (driver.Value, error) {
	if src.Valid {
		return src.Bytes, nil
	}

	return nil, nil
}
