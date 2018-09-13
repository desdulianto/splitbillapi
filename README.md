# SplitBillApi

## Description

Web API for split bill service.

## Owner

SRE

## Architecture Diagram

Add Architecture diagram here.

## SLO and SLI

- Availability: 99%
- Rata-rata response time: < 20ms

## Contact and On-Call Information

- Des Dulianto (des.dulianto@bukalapak.com) SRE

## Prerequisite

- Go
- Go Dep
- [Splitbill library](https://github.com/desdulianto/splitbill)

## Running Test

Run test using `make`

## Running Application

1. Install dependency `dep ensure`
2. Run web service `go run app/web/main.go`
3. Access web service endpoint on http://localhost:8080

## Onboarding and Development Guide

Add Onboarding and Development Guide here.

Some documentation about onboading and development process.

## Request Flows, Endpoints, and Dependencies

Endpoints:

- POST /
  
  POST with JSON payload:

  ```json
  {
      "Amount": 10000,
      "PaidBy": "Abang",
      "People": ["Abang", "Adek"]
  }
  ```

  Returns JSON payload:

  ```json
  {
      "AmountPay": 5000,
      "PayFrom": ["Adek"],
      "PayTo": "Abang"
  }
  ```

- GET /healthz

  For getting service health status. Should returns 200 and "ok".

## On-Call Runbooks

Add On-Call Runbooks here.

## FAQ

Add FAQ here.

## Links

- [Golang](https://tour.golang.org/basics/1)
- [Write Go Code](https://golang.org/doc/code.html)
- [Go Dep](https://github.com/golang/dep)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review](https://github.com/golang/go/wiki/CodeReviewComments)
- [Splitbill Library](https://github.com/desdulianto/splitbill)
