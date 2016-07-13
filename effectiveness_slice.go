package pokemon

import (
	"sort"
)

type EffectivenessIterable interface {
	ForEach(consumer func(Effectiveness))
	Filter(predicate func(Effectiveness) bool) EffectivenessIterable
	Sort(less func(Effectiveness, Effectiveness) bool) EffectivenessIterable
}

type EffectivenessSlice []Effectiveness

func (es EffectivenessSlice) ForEach(consumer func(Effectiveness)) {
	for _, e := range es {
		consumer(e)
	}
}

func (es EffectivenessSlice) Filter(predicate func(Effectiveness) bool) EffectivenessIterable {
	result := make(EffectivenessSlice, 0)
	for _, e := range es {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

func (es EffectivenessSlice) Sort(less func(Effectiveness, Effectiveness) bool) EffectivenessIterable {
	result := make(EffectivenessSlice, len(es))
	copy(result, es)
	sort.Stable(sortableEffectivenessSlice{result, less})
	return result
}

func AllEffectivenesses() EffectivenessIterable {
	return virtualEffectivenessSlice(0)
}

type virtualEffectivenessSlice int

func (es virtualEffectivenessSlice) ForEach(consumer func(Effectiveness)) {
	for e := EX; e <= EO; e++ {
		consumer(e)
	}
}

func (es virtualEffectivenessSlice) Filter(predicate func(Effectiveness) bool) EffectivenessIterable {
	result := make(EffectivenessSlice, 0)
	for e := EX; e <= EO; e++ {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

func (es virtualEffectivenessSlice) Sort(less func(Effectiveness, Effectiveness) bool) EffectivenessIterable {
	result := make(EffectivenessSlice, 4)
	for i, e := 0, EX; e <= EO; i, e = i+1, e+1 {
		result[i] = e
	}
	sort.Stable(sortableEffectivenessSlice{result, less})
	return result
}

type sortableEffectivenessSlice struct {
	slice EffectivenessSlice
	less  func(Effectiveness, Effectiveness) bool
}

func (es sortableEffectivenessSlice) Len() int {
	return len(es.slice)
}

func (es sortableEffectivenessSlice) Less(i, j int) bool {
	return es.less(es.slice[i], es.slice[j])
}

func (es sortableEffectivenessSlice) Swap(i, j int) {
	es.slice[i], es.slice[j] = es.slice[j], es.slice[i]
}
