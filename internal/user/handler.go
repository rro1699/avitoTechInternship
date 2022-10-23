package user

import (
	"avitoTechInternship/internal/handlers"
	"avitoTechInternship/internal/order"
	"avitoTechInternship/internal/report"
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
	logger      *logging.Logger
	serviceUser *Service
}

func NewHandler(logger *logging.Logger, serviceUser *Service) handlers.Handler {
	return &handler{
		logger:      logger,
		serviceUser: serviceUser,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(fmt.Sprintf("%saccrual", userURL), h.Accrual)
	router.POST(fmt.Sprintf("%sreserv", userURL), h.Reservation)
	router.POST(fmt.Sprintf("%srecogn", userURL), h.Recognition)
	router.POST(fmt.Sprintf("%sbalance", userURL), h.GetBalance)
	router.POST(fmt.Sprintf("%sreport", userURL), h.GetReport)
}

// @Summary Accrual
// @Description The method of accruing funds to the balance
// @Accept  json
// @Produce  json
// @Param input body UserDTO true "info about accruing balance"
// @Router /accrual [post]
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
	id, err := h.serviceUser.Accrual(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Accrual for user with id:%s succesfully ended", id)
	h.successResponse(writer, message, http.StatusOK)
	return
}

// @Summary Reservation
// @Description Method of reserving funds of the main balance in a separate account
// @Accept  json
// @Produce  json
// @Param input body order.OrderDTO true "reservation information"
// @Router /reserv [post]
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
	message, err := h.serviceUser.Reservation(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	h.successResponse(writer, message, http.StatusOK)
	return
}

// @Summary Recognition
// @Description Revenue recognition method
// @Accept  json
// @Produce  json
// @Param input body order.OrderDTO true "recognition information"
// @Router /recogn [post]
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

	message, err := h.serviceUser.Recognition(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	h.successResponse(writer, message, http.StatusOK)
	return
}

// @Summary GetBalance
// @Description User balance receipt method
// @Accept  json
// @Produce  json
// @Param input body UserDTO true "user information"
// @Router /balance [post]
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
	writer.Write(respData)
	return
}

// @Summary GetReport
// @Description method to get monthly report
// @Accept  json
// @Produce  json
// @Param input body report.ReportDTO true "information about date"
// @Router /report [post]
func (h *handler) GetReport(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.errorResponse(writer, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var reqData report.ReportDTO
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	filePath, err := h.serviceUser.GetReport(&reqData)
	if err != nil {
		h.errorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	h.successResponse(writer, filePath, http.StatusOK)
	return
}

func (h *handler) errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	h.logger.Info(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func (h *handler) successResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
