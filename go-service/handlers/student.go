package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go-service/services"
)

type StudentHandler struct {
	api *services.APIClient
	pdf *services.PDFGenerator
}

func NewStudentHandler(api *services.APIClient, pdf *services.PDFGenerator) *StudentHandler {
	return &StudentHandler{api: api, pdf: pdf}
}

func (h *StudentHandler) GetStudentReport(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing student id", http.StatusBadRequest)
		return
	}

	student, err := h.api.GetStudent(id)
	if err != nil {
		log.Printf("failed to fetch student %s: %v", id, err)
		http.Error(w, "failed to fetch student data", http.StatusInternalServerError)
		return
	}

	pdfData, err := h.pdf.GenerateStudentReport(student)
	if err != nil {
		log.Printf("pdf generation failed for student %s: %v", id, err)
		http.Error(w, "failed to generate pdf", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=student_%s_report.pdf", id))
	w.Write(pdfData)
}
