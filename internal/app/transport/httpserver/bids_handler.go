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

// Helper function to extract username
func (h HttpServer) getUsername(r *http.Request, usernameKey string) (string, error) {
	username := r.URL.Query().Get(usernameKey)
	if username == "" {
		return "", errors.New("missing-username")
	}
	return username, nil
}

// Helper function to extract pagination
func extractPagination(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	limit := 5
	offset := (page - 1) * limit
	return limit, offset
}

// Helper function to get and validate user
func (h HttpServer) getUser(r *http.Request, w http.ResponseWriter, usernameKey string) (*domain.User, bool) {
	username, err := h.getUsername(r, usernameKey)
	if err != nil {
		server.BadRequest("missing-username", err, w, r)
		return nil, false
	}
	user, err := h.userService.GetUser(r.Context(), username)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
		} else {
			server.RespondWithError(err, w, r)
		}
		return nil, false
	}
	return user, true
}

// CreateBID creates a new bid
func (h HttpServer) CreateBid(w http.ResponseWriter, r *http.Request) {
	var bidRequest BidRequest
	if err := json.NewDecoder(r.Body).Decode(&bidRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := bidRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	// Get user
	_, err := h.userService.GetUserByID(r.Context(), bidRequest.AuthorId)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NonAuthorised("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	if _, err := h.tenderService.GetTenderByID(r.Context(), bidRequest.TenderId); err != nil {
		server.NotFound("tender-not-found", err, w, r)
		return
	}

	bid, err := toDomainBid(bidRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	createdBid, err := h.bidService.CreateBid(r.Context(), bid)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBid(createdBid)
	server.RespondOK(response, w, r)
}

func (h HttpServer) GetTenderBids(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenderId := vars["tenderId"]
	if _, err := h.tenderService.GetTenderByID(r.Context(), tenderId); err != nil {
		server.NotFound("tender-not-found", err, w, r)
		return
	}

	user, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}
	limit, offset := extractPagination(r)

	bids, err := h.bidService.GetTenderBids(r.Context(), tenderId, user.ID().String(), limit, offset)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]BidResponse, 0, len(bids))
	for _, tender := range bids {
		response = append(response, toResponseBid(tender))
	}

	server.RespondOK(response, w, r)
}

func (h HttpServer) GetUserBids(w http.ResponseWriter, r *http.Request) {
	user, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}
	limit, offset := extractPagination(r)

	bids, err := h.bidService.GetUserBids(r.Context(), user.ID().String(), limit, offset)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]BidResponse, 0, len(bids))
	for _, tender := range bids {
		response = append(response, toResponseBid(tender))
	}

	server.RespondOK(response, w, r)
}

func (h HttpServer) GetBidStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}

	bidID := vars["bidId"]
	if _, err := h.bidService.GetBidByID(r.Context(), bidID); err != nil {
		server.NotFound("bid-not-found", err, w, r)
		return
	}

	status, err := h.bidService.GetBidStatus(r.Context(), bidID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("bid-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(status, w, r)
}

func (h HttpServer) UpdateBid(w http.ResponseWriter, r *http.Request) {
	var updateRequest UpdateBidRequest

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}
	if err := updateRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	vars := mux.Vars(r)

	bidID := vars["bidId"]
	if _, err := h.bidService.GetBidByID(r.Context(), bidID); err != nil {
		server.NotFound("bid-not-found", err, w, r)
		return
	}

	_, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}

	updateRequest.ID = bidID
	bid, err := toDomainUpdateBid(updateRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}
	updatedBid, err := h.bidService.UpdateBid(r.Context(), bid)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBid(updatedBid)
	server.RespondOK(response, w, r)
}

func (h HttpServer) BidFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bidID := vars["bidId"]

	if _, err := h.bidService.GetBidByID(r.Context(), bidID); err != nil {
		server.NotFound("bid-not-found", err, w, r)
		return
	}

	_, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}

	feedback := r.URL.Query().Get("bidFeedback")
	if feedback == "" || len(feedback) > 100 {
		server.BadRequest("missing-feedback", errors.New("missing feedback"), w, r)
		return
	}

	updatedBid, err := h.bidService.UpdateBidDescription(r.Context(), bidID, feedback)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBid(updatedBid)
	server.RespondOK(response, w, r)
}

func (h HttpServer) UpdateBidStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}

	bidID := vars["bidId"]
	if _, err := h.bidService.GetBidByID(r.Context(), bidID); err != nil {
		server.NotFound("bid-not-found", err, w, r)
		return
	}

	status := r.URL.Query()["status"]
	err := validateBidStatus(status[0])
	if err != nil {
		server.BadRequest("invalid-status", err, w, r)
		return
	}

	updatedBid, err := h.bidService.UpdateBidStatus(r.Context(), bidID, status[0])
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBid(updatedBid)
	server.RespondOK(response, w, r)
}

func (h HttpServer) SubmitDecision(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bidID := vars["bidId"]
	bid, err := h.bidService.GetBidByID(r.Context(), bidID)
	if err != nil {
		server.NotFound("bid-not-found", err, w, r)
		return
	}

	user, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}

	author, _ := h.userService.GetUserByID(r.Context(), bid.AuthorId().String())

	if author.OrganizationID() != user.OrganizationID() {
		server.Forbidden("user don't have permissions ", err, w, r)
		return
	}

	decision := r.URL.Query()["decision"]
	err = validateDecision(decision[0])
	if err != nil {
		server.BadRequest("invalid-decision", err, w, r)
	}

	updatedBid, err := h.bidService.UpdateBidStatus(r.Context(), bidID, decision[0])
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBid(updatedBid)
	server.RespondOK(response, w, r)
}

func (h HttpServer) RollbackBidVersion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bidID := vars["bidId"]
	if _, err := h.bidService.GetBidByID(r.Context(), bidID); err != nil {
		server.NotFound("bid-not-found", err, w, r)
		return
	}

	version, err := strconv.Atoi(vars["version"])
	if err != nil || version <= 0 {
		server.BadRequest("invalid version", err, w, r)
		return
	}

	_, ok := h.getUser(r, w, "username")
	if !ok {
		return
	}

	updatedBid, err := h.bidService.RollbackBidVersion(r.Context(), bidID, version)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBid(updatedBid)
	server.RespondOK(response, w, r)
}

func (h HttpServer) GetReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenderID := vars["tenderId"]
	if _, err := h.tenderService.GetTenderByID(r.Context(), tenderID); err != nil {
		server.NotFound("tender-not-found", err, w, r)
		return
	}

	authorUsername, ok := h.getUser(r, w, "authorUsername")
	if !ok {
		return
	}
	requesterUsername, ok := h.getUser(r, w, "requesterUsername")
	if !ok {
		return
	}
	_ = requesterUsername

	limit, offset := extractPagination(r)

	bids, err := h.bidService.GetTenderBids(r.Context(), tenderID, authorUsername.ID().String(), limit, offset)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]BidResponse, 0, len(bids))
	for _, tender := range bids {
		response = append(response, toResponseBid(tender))
	}

	server.RespondOK(response, w, r)
}
