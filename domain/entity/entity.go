package entity

import (
	"time"

	"github.com/guregu/null"
)

type Branch struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt null.Time `db:"updated_at,omitempty"`
	DeletedAt null.Time `db:"deleted_at,omitempty"`
}

type Event struct {
	ID          string      `db:"id"`
	AdminId     null.String `db:"admin_id,omitempty"`
	BranchId    string      `db:"branch_id"`
	Name        string      `db:"name"`
	StartTime   null.Time   `db:"start_time,omitempty"`
	EndTime     null.Time   `db:"end_time,omitempty"`
	Location    string      `db:"location"`
	Description string      `db:"description"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   null.Time   `db:"updated_at,omitempty"`
	DeletedAt   null.Time   `db:"deleted_at,omitempty"`
}

type Documentation struct {
	ID          string      `db:"id"`
	AdminId     null.String `db:"admin_id,omitempty"`
	BranchId    string      `db:"branch_id"`
	Date        time.Time   `db:"date"`
	Participant int         `db:"participant"`
	Location    string      `db:"location"`
	Description string      `db:"description"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   null.Time   `db:"updated_at,omitempty"`
	DeletedAt   null.Time   `db:"deleted_at,omitempty"`
}

type Participant struct {
	ID        string            `db:"id"`
	UserId    null.String       `db:"user_id,omitempty"`
	AdminId   null.String       `db:"admin_id,omitempty"`
	EventId   string            `db:"event_id"`
	Status    StatusParticipant `db:"status"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt null.Time         `db:"updated_at,omitempty"`
	DeletedAt null.Time         `db:"deleted_at,omitempty"`
}

// can be stored as object value
type StatusParticipant string

const (
	Hadir      StatusParticipant = "hadir"
	TidakHadir StatusParticipant = "tidak_hadir"
)

type Photo struct {
	ID              string      `db:"id"`
	DocumentationId string      `db:"documentation_id"`
	Url             null.String `db:"url,omitempty"`
	Name            null.String `db:"name,omitempty"`
	CreatedAt       time.Time   `db:"created_at"`
	UpdatedAt       null.Time   `db:"updated_at,omitempty"`
	DeletedAt       null.Time   `db:"deleted_at,omitempty"`
}
