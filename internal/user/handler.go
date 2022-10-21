package user

import (
	"avitoTechInternship/internal/handlers"
	"avitoTechInternship/internal/order"
	"avitoTechInternship/pkg/logging"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	userURL = "/user/"
)

type handler struct {
	logger       *logging.Logger
	serviceUser  *Service
	serviceOrder *order.Service
}

func NewHandler(logger *logging.Logger, serviceUser *Service, serviceOrder *order.Service) handlers.Handler {
	return &handler{
		logger:       logger,
		serviceUser:  serviceUser,
		serviceOrder: serviceOrder,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(fmt.Sprintf("%saccrual", userURL), h.Accrual)
	router.POST(fmt.Sprintf("%sreserv", userURL), h.Reservation)
	router.POST(fmt.Sprintf("%srecogn", userURL), h.Recognition)
	router.GET(fmt.Sprintf("%sbalance", userURL), h.GetBalance)

}

func (h *handler) Accrual(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.errorResponse(writer, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var reqData UserDTO
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.serviceUser.Accrual(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	return
}

func (h *handler) Reservation(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.errorResponse(writer, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var reqData order.OrderDTO
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.serviceOrder.Reservation(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	return
}

func (h *handler) Recognition(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.errorResponse(writer, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var reqData order.OrderDTO
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.serviceOrder.Recognition(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	return
}

func (h *handler) GetBalance(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.errorResponse(writer, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var reqData UserDTO
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	balance, err := h.serviceUser.GetBalance(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	respData, err := json.Marshal(balance)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	h.logger.Error(writer.Write(respData))
	return
}

func (h *handler) errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	h.logger.Error(w.Write(jsonResp))
}
