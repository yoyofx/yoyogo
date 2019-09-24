package Utils

func MergeMap(source map[string][]string, tagret map[string][]string) {
	for k, v := range tagret {
		source[k] = v
	}
}
