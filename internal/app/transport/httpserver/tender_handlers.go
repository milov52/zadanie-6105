package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/common/server"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

// CreateTender creates a new tender
func (h HttpServer) CreateTender(w http.ResponseWriter, r *http.Request) {
	var tenderRequest TenderRequest
	if err := json.NewDecoder(r.Body).Decode(&tenderRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := tenderRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	// Get user
	user, err := h.userService.GetUser(r.Context(), tenderRequest.CreatorUsername)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	if user.OrganizationID().String() != tenderRequest.OrganizationId {
		server.Forbidden("user is not work in this organization", err, w, r)
		return
	}

	tenderRequest.UserID = user.ID().String()
	tender, err := toDomainTender(tenderRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	createdTender, err := h.tenderService.CreateTender(r.Context(), tender)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseTender(createdTender)
	server.RespondOK(response, w, r)
}

func (h HttpServer) GetTenders(w http.ResponseWriter, r *http.Request) {
	// filter by service_type
	queryServiceTypes := r.URL.Query()["service_type"]

	var serviceTypes []string
	for _, serviceType := range queryServiceTypes {
		err := validateTenderServiceType(serviceType)
		if err != nil {
			if errors.Is(err, domain.ErrNegative) {
				server.BadRequest("wrong service type", err, w, r)
				return
			}
		}
		serviceTypes = append(serviceTypes, serviceType)
	}

	// limit
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit > 50 {
		server.BadRequest("wrong limit", err, w, r)
		return
	}
	if limit == 0 {
		limit = 5
	}

	// offset
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		server.BadRequest("wrong offset", err, w, r)
		return
	}

	tenders, err := h.tenderService.GetTenders(r.Context(), serviceTypes, limit, offset)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]TenderResponse, 0, len(tenders))
	for _, tender := range tenders {
		response = append(response, toResponseTender(tender))
	}

	server.RespondOK(response, w, r)
}

func (h HttpServer) GetUserTenders(w http.ResponseWriter, r *http.Request) {
	// filter by username
	username := r.URL.Query()["username"][0]
	// Get user
	user, err := h.userService.GetUser(r.Context(), username)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	// limit
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit > 50 {
		server.BadRequest("wrong limit", err, w, r)
		return
	}
	if limit == 0 {
		limit = 5
	}

	// offset
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		server.BadRequest("wrong offset", err, w, r)
		return
	}

	tenders, err := h.tenderService.GetUserTenders(r.Context(), user.ID().String(), limit, offset)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]TenderResponse, 0, len(tenders))
	for _, tender := range tenders {
		response = append(response, toResponseTender(tender))
	}

	server.RespondOK(response, w, r)
}

func (h HttpServer) GetTenderStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenderID := vars["tenderId"]

	// Get user
	username := r.URL.Query()["username"]
	if username == nil || len(username) == 0 {
		server.BadRequest("missing-username", errors.New("missing_username"), w, r)
		return
	}

	user, err := h.userService.GetUser(r.Context(), username[0])
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	tender, err := h.tenderService.GetTenderByID(r.Context(), tenderID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("tender-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	if tender.UserID() != user.ID() {
		server.Forbidden("not user tender", err, w, r)
		return
	}

	server.RespondOK(tender.Status(), w, r)
}

func (h HttpServer) UpdateTender(w http.ResponseWriter, r *http.Request) {
	var updateRequest UpdateTenderRequest

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}
	if err := updateRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	vars := mux.Vars(r)
	tenderID := vars["tenderId"]

	tender, err := h.tenderService.GetTenderByID(r.Context(), tenderID)
	if err != nil {
		server.NotFound("tender-not-found", err, w, r)
		return
	}

	username := r.URL.Query()["username"]
	if username == nil || len(username) == 0 {
		server.BadRequest("missing-username", errors.New("missing_username"), w, r)
		return
	}

	user, err := h.userService.GetUser(r.Context(), username[0])
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	if tender.UserID() != user.ID() {
		server.Forbidden("not user tender", err, w, r)
		return
	}

	updateRequest.ID = tender.ID()
	domainTender, err := toDomainUpdateTender(updateRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	updatedTender, err := h.tenderService.UpdateTender(r.Context(), domainTender)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseTender(updatedTender)
	server.RespondOK(response, w, r)
}

func (h HttpServer) UpdateTenderStatus(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query()["username"]
	if username == nil || len(username) == 0 {
		server.BadRequest("missing-username", errors.New("missing username"), w, r)
		return
	}

	user, err := h.userService.GetUser(r.Context(), username[0])
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	vars := mux.Vars(r)
	tenderID := vars["tenderId"]

	tender, err := h.tenderService.GetTenderByID(r.Context(), tenderID)
	if err != nil {
		server.NotFound("tender-not-found", err, w, r)
		return
	}

	if tender.UserID() != user.ID() {
		server.Forbidden("not user tenders", err, w, r)
		return
	}

	status := r.URL.Query()["status"]
	if status == nil || len(status) == 0 {
		server.BadRequest("missing-status", errors.New("missing status"), w, r)
		return
	}
	err = validateTenderStatus(status[0])
	if err != nil {
		server.BadRequest("invalid-status", err, w, r)
		return
	}

	updatedTender, err := h.tenderService.UpdateTenderStatus(r.Context(), tenderID, status[0])
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseTender(updatedTender)
	server.RespondOK(response, w, r)
}

func (h HttpServer) RollbackVersion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenderID := vars["tenderId"]
	tender, err := h.tenderService.GetTenderByID(r.Context(), tenderID)
	if err != nil {
		server.NotFound("tender-not-found", err, w, r)
		return
	}

	username := r.URL.Query()["username"]
	if username == nil || len(username) == 0 {
		server.BadRequest("missing-username", err, w, r)
		return
	}

	user, err := h.userService.GetUser(r.Context(), username[0])
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	if tender.UserID() != user.ID() {
		server.Forbidden("not user tenders", err, w, r)
		return
	}

	version, err := strconv.Atoi(vars["version"])
	if err != nil || version <= 0 {
		server.BadRequest("invalid version", err, w, r)
		return
	}

	updatedTender, err := h.tenderService.RollbackVersion(r.Context(), tenderID, version)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseTender(updatedTender)
	server.RespondOK(response, w, r)
}
