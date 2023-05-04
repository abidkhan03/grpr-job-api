package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grpr-job-api/internal/response"
	"github.com/grpr-job-api/internal/service"
)

type CombinedJob struct {
	svc *service.Service
}

func NewCombinedJob(svc *service.Service) *CombinedJob {
	return &CombinedJob{svc: svc}
}

func (api *CombinedJob) Routes(r chi.Router) {
	r.Route("/combined", func(r chi.Router) {
		r.Get("/", api.GetCombinedJobs)
	})
}

func (api *CombinedJob) GetCombinedJobs(w http.ResponseWriter, r *http.Request) {
	// fetch jobs from fetcherjob table
	jobs, err := api.svc.GetCombinedJobs(r.Context())
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var jobsResponse []*response.CombinedJob
	for _, job := range jobs {
		jobsResponse = append(jobsResponse, &response.CombinedJob{
			Id:                             job.Id,
			Created:                        job.Created,
			ReportTitle:                    job.ReportTitle,
			GroupingMethod:                 job.GroupingMethod,
			MainKeywordGroupingAccuracy:    job.MainKeywordGroupingAccuracy,
			VariantKeywordGroupingAccuracy: job.VariantKeywordGroupingAccuracy,
			SearchEngine:                   job.SearchEngine,
			Location:                       job.Location,
			Language:                       job.Language,
		})
	}

	respondOk(w, jobsResponse)

}
