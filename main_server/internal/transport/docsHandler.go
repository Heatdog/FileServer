package transport

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	Register(router *mux.Router)
}

type DocsHandler struct {
	logger *slog.Logger
}

func NewDocshandler(logger *slog.Logger) Handler {
	return &DocsHandler{
		logger: logger,
	}
}

const (
	getDoc    = "doc/get"
	setDoc    = "doc/set"
	updateDoc = "doc/update"
	deleteDoc = "doc/delete"
)

func (handler *DocsHandler) Register(router *mux.Router) {
	router.HandleFunc(getDoc, handler.getDoc).Methods(http.MethodGet)
	router.HandleFunc(setDoc, handler.setDoc).Methods(http.MethodPost)
	router.HandleFunc(updateDoc, handler.updateDoc).Methods(http.MethodPut)
	router.HandleFunc(deleteDoc, handler.deleteDoc).Methods(http.MethodDelete)
}

func (handler *DocsHandler) getDoc(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("get doc")
}

func (handler *DocsHandler) setDoc(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("set doc")
}

func (handler *DocsHandler) updateDoc(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("update doc")
}

func (handler *DocsHandler) deleteDoc(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("delete doc")
}
