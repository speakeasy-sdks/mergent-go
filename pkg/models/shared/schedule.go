// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type ScheduleInput struct {
	// A [cron expression](https://crontab.guru/examples.html) describing the
	// Schedule on which Tasks will run (UTC).
	//
	// Note: execution n + 1 of a Task will not begin until execution n has
	// completed successfully.
	//
	// You must pass either `cron` or `rrule` when creating a new Schedule.
	//
	Cron *string `json:"cron,omitempty"`
	// An optional description of the Schedule. This string must not contain more
	// than 500 characters.
	//
	Description *string `json:"description,omitempty"`
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations)
	// denoting the start of an RRULE schedule. Example: "2021-10-01T15:53:05Z".
	//
	// When not set, it will be set to the current time, and the first Task will
	// be scheduled immediately.
	//
	// Ignored for `cron`-type Schedules.
	//
	Dtstart *time.Time `json:"dtstart,omitempty"`
	// An optional name of the Schedule. This string must not contain more than
	// 100 characters.
	//
	Name *string `json:"name,omitempty"`
	// If `true`, the Schedule will be paused immediately. If `false`, a paused
	// Schedule will be resumed.
	//
	Paused *bool `json:"paused,omitempty"`
	// The name of the of the queue to schedule the Task on. This string must not
	// contain more than 100 characters.
	//
	Queue   *string  `json:"queue,omitempty"`
	Request *Request `json:"request,omitempty"`
	// An [iCal RRule expression](https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html)
	// describing the Schedule on which Tasks will run (UTC). The time of
	// Schedule creation will be used as the start of the recurrence interval
	// (i.e. `DTSTART`).
	//
	// Note: execution n + 1 of a Task will not begin until execution n has
	// completed successfully.
	//
	// You must pass either `cron` or `rrule` when creating a new Schedule.
	//
	Rrule *string `json:"rrule,omitempty"`
}

// Schedule - OK
type Schedule struct {
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations) representing when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A [cron expression](https://crontab.guru/examples.html) describing the
	// Schedule on which Tasks will run (UTC).
	//
	// Note: execution n + 1 of a Task will not begin until execution n has
	// completed successfully.
	//
	// You must pass either `cron` or `rrule` when creating a new Schedule.
	//
	Cron *string `json:"cron,omitempty"`
	// An optional description of the Schedule. This string must not contain more
	// than 500 characters.
	//
	Description *string `json:"description,omitempty"`
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations)
	// denoting the start of an RRULE schedule. Example: "2021-10-01T15:53:05Z".
	//
	// When not set, it will be set to the current time, and the first Task will
	// be scheduled immediately.
	//
	// Ignored for `cron`-type Schedules.
	//
	Dtstart *time.Time `json:"dtstart,omitempty"`
	// A unique ID assigned upon creation.
	ID *string `json:"id,omitempty"`
	// An optional name of the Schedule. This string must not contain more than
	// 100 characters.
	//
	Name *string `json:"name,omitempty"`
	// If `true`, the Schedule will be paused immediately. If `false`, a paused
	// Schedule will be resumed.
	//
	Paused *bool `json:"paused,omitempty"`
	// The name of the of the queue to schedule the Task on. This string must not
	// contain more than 100 characters.
	//
	Queue   *string  `json:"queue,omitempty"`
	Request *Request `json:"request,omitempty"`
	// An [iCal RRule expression](https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html)
	// describing the Schedule on which Tasks will run (UTC). The time of
	// Schedule creation will be used as the start of the recurrence interval
	// (i.e. `DTSTART`).
	//
	// Note: execution n + 1 of a Task will not begin until execution n has
	// completed successfully.
	//
	// You must pass either `cron` or `rrule` when creating a new Schedule.
	//
	Rrule *string `json:"rrule,omitempty"`
}
