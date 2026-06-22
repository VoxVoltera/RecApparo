package models

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Author       uuid.UUID // user
	Creation     time.Time
	Tags         []Tag
	MotherTicket *uuid.UUID // ticket
	Tickets      []Ticket
}

type Ticket struct {
	ID          uuid.UUID
	Name        string
	Description string
	Comments    []Comment
	Author      uuid.UUID // user
	Creation    time.Time
	DueDate     time.Time
	Tags        []Tag
	BoardID     uuid.UUID  // mother board
	SubBoard    *uuid.UUID // board
	Assigned    []User
	ColumnID uuid.UUID
}

type Comment struct {
	ID       uuid.UUID
	Author   uuid.UUID // user
	Creation time.Time
	Content  string
	Replies  []Comment
	TicketID uuid.UUID  // mother ticket
	ParentID *uuid.UUID // parent comment
}

type User struct {
	ID       uuid.UUID
	Name     string
	Teams    []Team
	Assigned []Ticket
}

type Tag struct {
	ID          uuid.UUID
	Name        string
	Description string
	Colour      string // hex
	Special     bool
	BoardID     *uuid.UUID // board, nil = global
}

type Team struct {
	ID      uuid.UUID
	Name    string
	Members []User
}

type Column struct {
	ID       uuid.UUID
	BoardID  uuid.UUID // board
	Name     string
	Position int
}

type TicketTag struct {
	TicketID uuid.UUID // ticket
	TagID    uuid.UUID // tag
}

type TicketAssignee struct {
	TicketID uuid.UUID // ticket
	UserID   uuid.UUID // user
}

type TeamMember struct {
	TeamID uuid.UUID // team
	UserID uuid.UUID // user
	Role   string
}