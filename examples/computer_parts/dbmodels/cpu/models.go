// Code generated by sqlc. DO NOT EDIT.

package cpu

import (
	"time"
)

type CPU struct {
	ID            int64
	MakeID        int64
	Name          string
	Cores         int32
	ClockSpeedGhz string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
