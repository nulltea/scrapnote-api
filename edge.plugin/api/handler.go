package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/core/meta"

	"github.com/timoth-y/scrapnote-api/data.records/core/model"

	"github.com/timoth-y/scrapnote-api/edge.plugin/core/service"
	"github.com/timoth-y/scrapnote-api/edge.plugin/usecase/serializer/json"
)

type Handler struct {
	service     service.RecordService
	auth        *rest.AuthMiddleware
	contentType string
}

func NewHandler(service service.RecordService, auth core.AuthService, config config.CommonConfig) *Handler {
	return &Handler{
		service,
		rest.NewAuthMiddleware(auth),
		config.ContentType,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r,"recordID")
	order, err := h.service.GetOne(r.Context(), code); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, order, http.StatusOK)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	order, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.Add(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, order, http.StatusOK)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	order, err := h.getRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.Update(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", h.contentType)
	w.WriteHeader(statusCode)
	if body != nil {
		raw, err := h.serializer(h.contentType).Encode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(raw); err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) getRequestQuery(r *http.Request) (meta.RequestQuery, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).DecodeMap(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *Handler) getRequestBody(r *http.Request) (*model.Record, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}


func (h *Handler) serializer(contentType string) service.RecordSerializer {
	return json.NewSerializer()
}