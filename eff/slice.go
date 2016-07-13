package eff

type EffectivenessSlice []Effectiveness

func All() EffectivenessSlice {
	result := make([]Effectiveness, 4)
	for i, e := 0, EX; e <= EO; i, e = i+1, e+1 {
		result[i] = e
	}
	return result
}

func (es EffectivenessSlice) ForEach(consumer func(Effectiveness)) {
	for _, e := range es {
		consumer(e)
	}
}
