package splitbillapi_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/desdulianto/splitbill"
	"github.com/desdulianto/splitbillapi/handler"
)

const URL string = "http://localhost:8080"

func TestPostMethod(t *testing.T) {
	cases := []struct {
		in   splitbill.Bill
		want splitbillapi.BillResponse
	}{
		{
			in: splitbill.Bill{
				Amount: 100000,
				PaidBy: "A",
				People: splitbill.People{"A", "B", "C", "D", "E"},
			},
			want: splitbillapi.BillResponse{
				AmountPay: 100000 / 5,
				PayFrom:   splitbill.People{"B", "C", "D", "E"},
				PayTo:     "A",
			},
		},
		{
			in: splitbill.Bill{
				Amount: 100000,
				PaidBy: "A",
				People: splitbill.People{"A", "B", "C"},
			},
			want: splitbillapi.BillResponse{
				AmountPay: 100000 / 3,
				PayFrom:   splitbill.People{"B", "C"},
				PayTo:     "A",
			},
		},
		{
			in: splitbill.Bill{
				Amount: 200000,
				PaidBy: "",
				People: splitbill.People{"A", "B", "C"},
			},
			want: splitbillapi.BillResponse{
				AmountPay: 200000 / 3,
				PayFrom:   splitbill.People{"A", "B", "C"},
				PayTo:     "",
			},
		},
	}

	for _, c := range cases {
		bill, _ := json.Marshal(c.in)

		req := httptest.NewRequest("POST", URL, bytes.NewBuffer(bill))
		res := httptest.NewRecorder()
		splitbillapi.SplitBillHandler(res, req)

		result := res.Result()
		if result.StatusCode != 200 {
			t.Fatalf("POST %v expected %v status code, got %v", req, 200, result.StatusCode)
		}

		var response splitbillapi.BillResponse
		body := res.Body.Bytes()
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("POST %v responds with %v expected %v", bill, response, c.want)
		}
	}
}

func TestWithInvalidData(t *testing.T) {
	cases := []splitbill.Bill{
		splitbill.Bill{
			Amount: 0,
			PaidBy: "A",
			People: splitbill.People{"A", "B", "C", "D", "E"},
		},
		splitbill.Bill{
			Amount: -1,
			PaidBy: "A",
			People: splitbill.People{"A", "B", "C", "D", "E"},
		},
		splitbill.Bill{
			Amount: 100000,
			PaidBy: "A",
			People: splitbill.People{},
		},
		splitbill.Bill{
			Amount: 100000,
			PaidBy: "",
			People: splitbill.People{},
		},
	}

	for _, c := range cases {
		payload, _ := json.Marshal(c)
		req := httptest.NewRequest("POST", URL, bytes.NewBuffer(payload))
		res := httptest.NewRecorder()
		splitbillapi.SplitBillHandler(res, req)

		result := res.Result()
		if result.StatusCode != 400 {
			t.Fatalf("Request with %v should get bad request status code, but get %v",
				payload, result.StatusCode)
		}
	}
}

func TestWithInvalidMethod(t *testing.T) {
	methods := []string{"GET", "DELETE", "PUT"}
	payload, _ := json.Marshal(splitbill.Bill{
		Amount: 100000,
		PaidBy: "A",
		People: splitbill.People{"A", "B", "C", "D", "E"},
	})

	for _, method := range methods {
		req := httptest.NewRequest(method, URL, bytes.NewBuffer(payload))
		res := httptest.NewRecorder()
		splitbillapi.SplitBillHandler(res, req)

		result := res.Result()
		if result.StatusCode != 405 {
			t.Fatalf("Request with %v method should get response with %d status code, but it returns with %v",
				method, 405, result.StatusCode)
		}
	}
}
