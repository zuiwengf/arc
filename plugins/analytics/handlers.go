package analytics

import (
	"log"
	"net/http"
	"strconv"

	"github.com/appbaseio-confidential/arc/internal/util"
)

func (a *analytics) getOverview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("click_analytics")

		var clickAnalytics bool
		if q != "" {
			if v, err := strconv.ParseBool(q); err == nil {
				clickAnalytics = v
			}
		}
		from, to, size := queryParams(r.URL.Query())

		raw, err := a.es.analyticsOverview(from, to, size, clickAnalytics)
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getAdvanced() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("click_analytics")

		var clickAnalytics bool
		if q != "" {
			if v, err := strconv.ParseBool(q); err == nil {
				clickAnalytics = v
			}
		}
		from, to, size := queryParams(r.URL.Query())

		raw, err := a.es.advancedAnalytics(from, to, size, clickAnalytics)
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getPopularSearches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("click_analytics")

		var clickAnalytics bool
		if q != "" {
			if v, err := strconv.ParseBool(q); err == nil {
				clickAnalytics = v
			}
		}
		from, to, size := queryParams(r.URL.Query())

		raw, err := a.es.popularSearchesRaw(from, to, size, clickAnalytics)
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getNoResultsSearches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := a.es.noResultsSearchesRaw(queryParams(r.URL.Query()))
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getPopularFilters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("click_analytics")

		var clickAnalytics bool
		if q != "" {
			if v, err := strconv.ParseBool(q); err == nil {
				clickAnalytics = v
			}
		}
		from, to, size := queryParams(r.URL.Query())

		raw, err := a.es.popularFiltersRaw(from, to, size, clickAnalytics)
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getPopularResults() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("click_analytics")

		var clickAnalytics bool
		if q != "" {
			if v, err := strconv.ParseBool(q); err == nil {
				clickAnalytics = v
			}
		}
		from, to, size := queryParams(r.URL.Query())

		raw, err := a.es.popularResultsRaw(from, to, size, clickAnalytics)
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getGeoRequestsDistribution() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := a.es.geoRequestsDistribution(queryParams(r.URL.Query()))
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getLatencies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := a.es.latencies(queryParams(r.URL.Query()))
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}

func (a *analytics) getSummary() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := a.es.summary(queryParams(r.URL.Query()))
		if err != nil {
			log.Printf("%s: %v", logTag, err)
			util.WriteBackMessage(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		util.WriteBackRaw(w, raw, http.StatusOK)
	}
}