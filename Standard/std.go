package Std

type M = map[string]string

func PadLeft(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

//
