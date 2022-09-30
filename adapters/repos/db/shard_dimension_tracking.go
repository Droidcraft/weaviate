//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package db

import (
	"encoding/binary"
	"time"

	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
)

func (s *Shard) Dimensions() int {
	b := s.store.Bucket(helpers.DimensionsBucketLSM)
	if b == nil {
		return 0
	}

	c := b.MapCursor()
	sum := 0
	for k, v := c.First(); k != nil; k, v = c.Next() {
		dimLength := binary.LittleEndian.Uint32(k)
		sum += int(dimLength) * len(v)
	}
	c.Close()

	return sum
}

func (s *Shard) initDimensionTracking() {
	if !s.index.Config.TrackVectorDimensions {
		return
	}

	go func() {
		t := time.NewTicker(5 * time.Minute)

		for {
			if s.stopMetrics {
				return
			}

			dimCount := s.Dimensions()
			if s.promMetrics != nil {
				metric, err := s.promMetrics.VectorDimensionsSum.
					GetMetricWithLabelValues(s.index.Config.ClassName.String(), s.name)
				if err == nil {
					metric.Set(float64(dimCount))
				}
			}
			<-t.C
		}
	}()
}