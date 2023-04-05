// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type LockersLockStatus string

const (
	LockersLockStatusLocked   LockersLockStatus = "locked"
	LockersLockStatusUnlocked LockersLockStatus = "unlocked"
)

func (e *LockersLockStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LockersLockStatus(s)
	case string:
		*e = LockersLockStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for LockersLockStatus: %T", src)
	}
	return nil
}

type NullLockersLockStatus struct {
	LockersLockStatus LockersLockStatus
	Valid             bool // Valid is true if LockersLockStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLockersLockStatus) Scan(value interface{}) error {
	if value == nil {
		ns.LockersLockStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LockersLockStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLockersLockStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.LockersLockStatus, nil
}

type LockersStatus string

const (
	LockersStatusUsed        LockersStatus = "used"
	LockersStatusAvailable   LockersStatus = "available"
	LockersStatusMalfunction LockersStatus = "malfunction"
)

func (e *LockersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LockersStatus(s)
	case string:
		*e = LockersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for LockersStatus: %T", src)
	}
	return nil
}

type NullLockersStatus struct {
	LockersStatus LockersStatus
	Valid         bool // Valid is true if LockersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLockersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.LockersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LockersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLockersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.LockersStatus, nil
}

type SensorsKind string

const (
	SensorsKindTemperature SensorsKind = "temperature"
	SensorsKindMoisture    SensorsKind = "moisture"
	SensorsKindServo       SensorsKind = "servo"
	SensorsKindSpeaker     SensorsKind = "speaker"
	SensorsKindLock        SensorsKind = "lock"
)

func (e *SensorsKind) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SensorsKind(s)
	case string:
		*e = SensorsKind(s)
	default:
		return fmt.Errorf("unsupported scan type for SensorsKind: %T", src)
	}
	return nil
}

type NullSensorsKind struct {
	SensorsKind SensorsKind
	Valid       bool // Valid is true if SensorsKind is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSensorsKind) Scan(value interface{}) error {
	if value == nil {
		ns.SensorsKind, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SensorsKind.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSensorsKind) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.SensorsKind, nil
}

type UsersRole string

const (
	UsersRoleAdmin    UsersRole = "admin"
	UsersRoleCustomer UsersRole = "customer"
)

func (e *UsersRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersRole(s)
	case string:
		*e = UsersRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersRole: %T", src)
	}
	return nil
}

type NullUsersRole struct {
	UsersRole UsersRole
	Valid     bool // Valid is true if UsersRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersRole) Scan(value interface{}) error {
	if value == nil {
		ns.UsersRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UsersRole, nil
}

type Locker struct {
	ID           int32             `json:"id"`
	LockerNumber int32             `json:"lockerNumber"`
	Location     string            `json:"location"`
	Status       LockersStatus     `json:"status"`
	NfcSig       string            `json:"nfcSig"`
	LastAccessed sql.NullTime      `json:"lastAccessed"`
	LockStatus   LockersLockStatus `json:"lockStatus"`
	CreatedAt    sql.NullTime      `json:"createdAt"`
	LastModified sql.NullTime      `json:"lastModified"`
}

type LockerSensor struct {
	ID           int32        `json:"id"`
	SensorID     int32        `json:"sensorID"`
	LockerID     int32        `json:"lockerID"`
	CreatedAt    sql.NullTime `json:"createdAt"`
	LastModified sql.NullTime `json:"lastModified"`
}

type LockerUser struct {
	ID           int32        `json:"id"`
	UserID       int32        `json:"userID"`
	LockerID     int32        `json:"lockerID"`
	CreatedAt    sql.NullTime `json:"createdAt"`
	LastModified sql.NullTime `json:"lastModified"`
}

type Sensor struct {
	ID           int32        `json:"id"`
	FeedKey      string       `json:"feedKey"`
	Kind         SensorsKind  `json:"kind"`
	CreatedAt    sql.NullTime `json:"createdAt"`
	LastModified sql.NullTime `json:"lastModified"`
}

type User struct {
	ID             int32        `json:"id"`
	Name           string       `json:"name"`
	PasswordHashed string       `json:"passwordHashed"`
	Email          string       `json:"email"`
	Role           UsersRole    `json:"role"`
	CreatedAt      sql.NullTime `json:"createdAt"`
	LastModified   sql.NullTime `json:"lastModified"`
}
