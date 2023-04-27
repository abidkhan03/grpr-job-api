package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grpr-job-api/internal/response"
	"github.com/grpr-job-api/internal/service"
)

type GrouperJob struct {
	svc *service.Service
}

func NewGrouper(svc *service.Service) *GrouperJob {
	return &GrouperJob{svc: svc}
}

func (api *GrouperJob) Routes(r chi.Router) {
	r.Route("/grouper", func(r chi.Router) {
		r.Get("/", api.GetGrouperJobs)
	})
}

func (api *GrouperJob) GetGrouperJobs(w http.ResponseWriter, r *http.Request) {
	// fetch jobs from fetcherjob table
	jobs, err := api.svc.GetGrouperJobs(r.Context())
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var jobsResponse []*response.GrouperJob
	for _, job := range jobs {
		jobsResponse = append(jobsResponse, &response.GrouperJob{
			Id:                             job.Id,
			Created:                        job.Created,
			ReportTitle:                    job.ReportTitle,
			GroupingMethod:                 job.GroupingMethod,
			MainKeywordGroupingAccuracy:    job.MainKeywordGroupingAccuracy,
			VariantKeywordGroupingAccuracy: job.VariantKeywordGroupingAccuracy,
		})
	}

	respondOk(w, jobsResponse)

}
