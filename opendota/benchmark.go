package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

// BenchmarkService provides methods for the retrieving
// information about benchmarks of a hero.
type BenchmarkService struct {
	sling *sling.Sling
}

// Benchmark is a collection of benchmarks about a hero.
type Benchmark struct {
	HeroID int             `json:"hero_id"`
	Result benchmarkResult `json:"result"`
}

type benchmarkResult struct {
	GoldPerMin        []benchmark `json:"gold_per_min"`
	XpPerMin          []benchmark `json:"xp_per_min"`
	KillsPerMin       []benchmark `json:"kills_per_min"`
	LastHitsPerMin    []benchmark `json:"last_hits_per_min"`
	HeroDamagePerMin  []benchmark `json:"hero_damage_per_min"`
	HeroHealingPerMin []benchmark `json:"hero_healing_per_min"`
	TowerDamage       []benchmark `json:"tower_damage"`
}

type benchmark struct {
	Percentile float64 `json:"percentile"`
	Value      float64 `json:"value"`
}

// BenchmarkParam is used to specify the hero when retrieving
// benchmarks.
type BenchmarkParam struct {
	HeroID string `url:"hero_id"`
}

func newBenchmarkService(sling *sling.Sling) *BenchmarkService {
	return &BenchmarkService{
		sling: sling.Path("benchmarks"),
	}
}

// Benchmarks returns a collection of benchmarks about a hero.
func (s *BenchmarkService) Benchmarks(param *BenchmarkParam) (Benchmark, *http.Response, error) {
	benchmarks := new(Benchmark)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(param).Receive(benchmarks, apiError)
	return *benchmarks, resp, relevantError(err, *apiError)
}
