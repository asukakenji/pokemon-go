package pokemon

type TypeSlice []Type

func AllTypes() TypeSlice {
	result := make([]Type, 18)
	for i, t := 0, TypeNormal; t <= TypeFairy; i, t = i+1, t+1 {
		result[i] = t
	}
	return result
}

func (ts TypeSlice) ForEach(consumer func(Type)) {
	for _, t := range ts {
		consumer(t)
	}
}

/*
func (ts TypeSlice) Map(mapper func(Type) interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, t := range ts {
		result = append(result, mapper(t))
	}
	return result
}
*/
