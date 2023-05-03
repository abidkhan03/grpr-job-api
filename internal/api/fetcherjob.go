package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/grpr-job-api/internal/dao"
	"github.com/grpr-job-api/internal/request"
	"github.com/grpr-job-api/internal/response"
	"github.com/grpr-job-api/internal/service"
)

type FetcherJob struct {
	svc *service.Service
}

func NewFetcher(svc *service.Service) *FetcherJob {
	return &FetcherJob{svc: svc}
}

func (api *FetcherJob) Routes(r chi.Router) {
	r.Route("/fetcher", func(r chi.Router) {
		r.Get("/", api.GetFetcherJobs)
		r.Post("/job", api.AddFetcherJob)
	})
}

func (api *FetcherJob) GetFetcherJobs(w http.ResponseWriter, r *http.Request) {
	// fetch jobs from fetcherjob table
	jobs, err := api.svc.GetFetcherJobs(r.Context())
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

func (api *FetcherJob) AddFetcherJob(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var req request.FetcherJob
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// parse created time
	createdTime, err := time.Parse("2006-01-02", req.Created.Format("2006-01-02"))
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, "invalid created time")
		return
	}

	// create new job
	job := &dao.FetcherJob{
		Created:      createdTime,
		ReportTitle:  req.ReportTitle,
		SearchEngine: req.SearchEngine,
		Location:     req.Location,
		Language:     req.Language,
	}

	err = api.svc.AddFetcherJob(r.Context(), job)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondOk(w, "job added successfully")
}
