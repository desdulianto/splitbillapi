package splitbillapi

import (
	"encoding/json"
	"net/http"

	"github.com/desdulianto/splitbill"
)

// ErrorMessage encapsulates ErrorMessage reponse payload
type ErrorMessage struct {
	Code    int
	Message string
}

// BillResponse encapsulates Bill response payload
type BillResponse struct {
	AmountPay splitbill.Money
	PayFrom   splitbill.People
	PayTo     splitbill.Person
}

// Write response as JSON
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(data)
	if err == nil {
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		body, _ = json.Marshal(ErrorMessage{http.StatusInternalServerError, err.Error()})
		w.Write(body)
	}
}

// Create and return http response of ErrorMessage object
func errorResponse(w http.ResponseWriter, code int, message string) ErrorMessage {
	w.WriteHeader(code)
	return ErrorMessage{
		Code:    code,
		Message: message,
	}
}

// BadRequest response
func badRequest(w http.ResponseWriter, message string) ErrorMessage {
	return errorResponse(w, http.StatusBadRequest, message)
}

// handle POST request
func handlePost(w http.ResponseWriter, r *http.Request) interface{} {
	var response interface{}

	// decode request body
	var bill splitbill.Bill
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bill)
	if err != nil {
		response = badRequest(w, err.Error())
	} else {
		// generate response
		amountPay, err := bill.SplitEvenly()
		if err != nil {
			response = badRequest(w, err.Error())
		} else {
			response = BillResponse{
				AmountPay: amountPay,
				PayFrom:   bill.GetPeople(),
				PayTo:     bill.PaidBy,
			}
		}
	}
	return response
}

// SplitBillHandler handle split bill endpoint
func SplitBillHandler(w http.ResponseWriter, r *http.Request) {
	var response interface{}

	switch r.Method {
	case "POST":
		response = handlePost(w, r)
	default:
		response = errorResponse(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
	}
	jsonResponse(w, response)
}
