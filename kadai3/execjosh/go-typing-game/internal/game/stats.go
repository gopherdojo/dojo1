package game

// Stats represents game stats
type Stats struct {
	successCount int
	failureCount int
}

// SuccessCount returns the success count
func (s *Stats) SuccessCount() int {
	return s.successCount
}

// FailureCount returns the failure count
func (s *Stats) FailureCount() int {
	return s.failureCount
}

// LogSuccess logs a success
func (s *Stats) LogSuccess() {
	s.successCount++
}

// LogFailure logs a failure
func (s *Stats) LogFailure() {
	s.failureCount++
}
