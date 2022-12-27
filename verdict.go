package testlib

type Verdict string

const (
	NT Verdict = "NOT TESTED"
	OK Verdict = "ACCEPTED"
	WA Verdict = "WRONG ANSWER"
	PE Verdict = "PRESENTATION ERROR"
	EF Verdict = "EPIC FAIL"
	CF Verdict = "CHECK FAILED"
	TL Verdict = "TIME LIMIT EXCEED"
	ML Verdict = "MEMORY LIMIT EXCEED"
	SV Verdict = "SECURITY VIOLATION"
	CE Verdict = "COMPILATION ERROR"
	RE Verdict = "RUNTIME ERROR"
	IO Verdict = "INVALID IO"
	TT Verdict = "TESTED"
	WT Verdict = "WALL TIME LIMIT"
)
