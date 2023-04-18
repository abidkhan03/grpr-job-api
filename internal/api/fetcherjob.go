package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grpr-job-api/internal/response"
	"github.com/grpr-job-api/internal/service"
)

type FetcherJob struct {
	svc *service.Service
}

func NewCategory(svc *service.Service) *FetcherJob {
	return &FetcherJob{svc: svc}
}

func (api *FetcherJob) Routes(r chi.Router) {
	r.Route("/fetcher", func(r chi.Router) {
		r.Get("/", api.GetAllJobs)
	})
}

func (api *FetcherJob) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	// fetch jobs from fetcherjob table
	jobs, err := api.svc.GetAllJobs(r.Context())
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var jobsResponse []*response.FetcherJob
	for _, job := range jobs {
		jobsResponse = append(jobsResponse, &response.FetcherJob{
			Id:           job.Id,
			Created:      job.Created,
			ReportTitle:  job.ReportTitle,
			SearchEngine: job.SearchEngine,
			Location:     job.Location,
			Language:     job.Language,
		})
	}

	respondOk(w, jobsResponse)

}
