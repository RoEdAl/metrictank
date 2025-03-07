package mdata

import (
	"fmt"
	"math"

	"github.com/grafana/metrictank/consolidation"
	log "github.com/sirupsen/logrus"
)

// Aggregation is a container for all summary statistics / aggregated data for 1 metric, in 1 time frame
// if the Cnt is 0, the numbers don't necessarily make sense.
type Aggregation struct {
	Min float64
	Max float64
	Sum float64
	Cnt float64
	Lst float64
}

func NewAggregation() *Aggregation {
	return &Aggregation{
		Min: math.MaxFloat64,
		Max: -math.MaxFloat64,
	}
}

func (a *Aggregation) GetValueFor(consolidator consolidation.Consolidator) (float64, error) {
	switch consolidator {
	case consolidation.Cnt:
		return a.Cnt, nil
	case consolidation.Lst:
		return a.Lst, nil
	case consolidation.Min:
		return a.Min, nil
	case consolidation.Max:
		return a.Max, nil
	case consolidation.Sum:
		return a.Sum, nil
	}
	err := fmt.Errorf("internal error: AggMetric.GetAggregated(): unknown consolidator %q", consolidator)
	log.Errorf("AGG: %s", err.Error())
	return 0, err
}

func (a *Aggregation) Add(val float64) {
	a.Min = math.Min(val, a.Min)
	a.Max = math.Max(val, a.Max)
	a.Sum += val
	a.Cnt += 1
	a.Lst = val
}

func (a *Aggregation) Reset() {
	a.Min = math.MaxFloat64
	a.Max = -math.MaxFloat64
	a.Sum = 0
	a.Cnt = 0
	// no need to set a.Lst, for a to be valid (Cnt > 1), a.Lst will always be set properly
}
