// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: job/job.proto

package job

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _job_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PutTimerJobRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PutTimerJobRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAppId()) < 1 {
		return PutTimerJobRequestValidationError{
			field:  "AppId",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetJob() == nil {
		return PutTimerJobRequestValidationError{
			field:  "Job",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetJob()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutTimerJobRequestValidationError{
				field:  "Job",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCallback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutTimerJobRequestValidationError{
				field:  "Callback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PutTimerJobRequestValidationError is the validation error returned by
// PutTimerJobRequest.Validate if the designated constraints aren't met.
type PutTimerJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutTimerJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutTimerJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutTimerJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutTimerJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutTimerJobRequestValidationError) ErrorName() string {
	return "PutTimerJobRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PutTimerJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutTimerJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutTimerJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutTimerJobRequestValidationError{}

// Validate checks the field values on Job with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Job) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetJobId()) < 1 {
		return JobValidationError{
			field:  "JobId",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetExpire()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JobValidationError{
				field:  "Expire",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetTasks()) < 1 {
		return JobValidationError{
			field:  "Tasks",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTasks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobValidationError{
					field:  fmt.Sprintf("Tasks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if _, ok := Job_TaskStrategy_name[int32(m.GetTaskStrategy())]; !ok {
		return JobValidationError{
			field:  "TaskStrategy",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// JobValidationError is the validation error returned by Job.Validate if the
// designated constraints aren't met.
type JobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobValidationError) ErrorName() string { return "JobValidationError" }

// Error satisfies the builtin error interface
func (e JobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobValidationError{}

// Validate checks the field values on PutTimerJobResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PutTimerJobResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Status

	return nil
}

// PutTimerJobResponseValidationError is the validation error returned by
// PutTimerJobResponse.Validate if the designated constraints aren't met.
type PutTimerJobResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutTimerJobResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutTimerJobResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutTimerJobResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutTimerJobResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutTimerJobResponseValidationError) ErrorName() string {
	return "PutTimerJobResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PutTimerJobResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutTimerJobResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutTimerJobResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutTimerJobResponseValidationError{}

// Validate checks the field values on Task with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Task) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		return TaskValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "Target",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTaskContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "TaskContent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Reason

	// no validation rules for JobId

	// no validation rules for Id

	// no validation rules for TaskName

	return nil
}

// TaskValidationError is the validation error returned by Task.Validate if the
// designated constraints aren't met.
type TaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValidationError) ErrorName() string { return "TaskValidationError" }

// Error satisfies the builtin error interface
func (e TaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValidationError{}

// Validate checks the field values on InquireJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *InquireJobRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAppId()) < 1 {
		return InquireJobRequestValidationError{
			field:  "AppId",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetJobId()) < 1 {
		return InquireJobRequestValidationError{
			field:  "JobId",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// InquireJobRequestValidationError is the validation error returned by
// InquireJobRequest.Validate if the designated constraints aren't met.
type InquireJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InquireJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InquireJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InquireJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InquireJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InquireJobRequestValidationError) ErrorName() string {
	return "InquireJobRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InquireJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInquireJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InquireJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InquireJobRequestValidationError{}

// Validate checks the field values on InquireJobResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InquireJobResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetJob()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InquireJobResponseValidationError{
				field:  "Job",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	return nil
}

// InquireJobResponseValidationError is the validation error returned by
// InquireJobResponse.Validate if the designated constraints aren't met.
type InquireJobResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InquireJobResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InquireJobResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InquireJobResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InquireJobResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InquireJobResponseValidationError) ErrorName() string {
	return "InquireJobResponseValidationError"
}

// Error satisfies the builtin error interface
func (e InquireJobResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInquireJobResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InquireJobResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InquireJobResponseValidationError{}

// Validate checks the field values on InquireTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InquireTaskRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAppId()) < 1 {
		return InquireTaskRequestValidationError{
			field:  "AppId",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		return InquireTaskRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// InquireTaskRequestValidationError is the validation error returned by
// InquireTaskRequest.Validate if the designated constraints aren't met.
type InquireTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InquireTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InquireTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InquireTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InquireTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InquireTaskRequestValidationError) ErrorName() string {
	return "InquireTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InquireTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInquireTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InquireTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InquireTaskRequestValidationError{}

// Validate checks the field values on InquireTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InquireTaskResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InquireTaskResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	return nil
}

// InquireTaskResponseValidationError is the validation error returned by
// InquireTaskResponse.Validate if the designated constraints aren't met.
type InquireTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InquireTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InquireTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InquireTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InquireTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InquireTaskResponseValidationError) ErrorName() string {
	return "InquireTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e InquireTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInquireTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InquireTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InquireTaskResponseValidationError{}
