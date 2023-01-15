package generator

func stringPtr(v string) *string {
	return &v
}

func mapArray[I any, O any](in []I, f func(v I) O) []O {
	out := make([]O, len(in))
	for i, v := range in {
		out[i] = f(v)
	}

	return out
}

func prepend(x []byte, y []byte) []byte {
	return append(y, x...) // todo: maybe there is a better performant way?
}
