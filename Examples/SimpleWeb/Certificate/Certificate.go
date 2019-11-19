package Certificate

import (
	"path"
	"runtime"
)

func GetApplicationPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current frame")
	}

	return path.Dir(filename)

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	panic("Failed to get dir")
	//}
	//return dir
}

// GetCertificatePaths returns the paths to certificate and key
func GetCertificatePaths() (string, string) {
	certPath := GetApplicationPath()
	return path.Join(certPath, "server.crt"), path.Join(certPath, "server.key")
}
