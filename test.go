package main

import (
	"fmt"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 5000, Per: time.Second}
	duration := 4 * time.Second
	header := http.Header{}
	target := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/api/user/me",
		Header: header}
	target.Header.Add("Content-Type", "application/json")
	target.Header.Add("Authorization", "Bearer  eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNTY1MTcxMjI1LCJleHAiOjE1NjU3NzYwMjV9.sdTcC4-996fJ6onXRDYtYXODo29ZJGsW8_N4b6GuXZWPugim5xY44tu-ECjYksrU0mCgtir0v4RrpqubWokMBQ")

    signupTraget := vegeta.Target{
      Method: "GET",
      URL:    "http://localhost:8080/api/auth/signup/me",
      Header: header}
    target.Header.Add("Content-Type", "application/json")
    target.Header.Add("Authorization", "Bearer  eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNTY1MTcxMjI1LCJleHAiOjE1NjU3NzYwMjV9.sdTcC4-996fJ6onXRDYtYXODo29ZJGsW8_N4b6GuXZWPugim5xY44tu-ECjYksrU0mCgtir0v4RrpqubWokMBQ")


  targeter := vegeta.NewStaticTargeter(target, signupTraget)

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
