package ram

import (
	"math"

	"github.com/shirou/gopsutil/v3/mem"
)

type Usage struct {
	Total          uint64  `json:"total"`
	Available      uint64  `json:"available"`
	Used           uint64  `json:"used"`
	Free           uint64  `json:"free"`
	Percent        float64 `json:"percent"`
	RoundedPercent int     `json:"rounded_percent"`
}

func GetUsage() (*Usage, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	roundedPercent := int(math.Round(v.UsedPercent))

	return &Usage{
		Total:          v.Total,
		Available:      v.Available,
		Used:           v.Used,
		Free:           v.Free,
		Percent:        v.UsedPercent,
		RoundedPercent: roundedPercent,
	}, nil
}
