package types

import "time"

type Status int64

const (
	Failed Status = iota
	Passed
	Skipped
	NonConclusive
	Aborted
	Error
)

type TestResult struct {
	Id       string
	Class    string
	Suite    string
	TestCase string
	Outcome  TestOutcome
	Duration time.Duration
}
type TestOutcome struct {
	Id       string
	Status   Status
	TestCase string
}
