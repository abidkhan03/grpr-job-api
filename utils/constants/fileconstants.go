package constants

import "os"

func FileConstant() {
	ProjectBaseDir := os.Getenv("AG_GROUPER_HOME")
	if ProjectBaseDir == "" {
		ProjectBaseDir = "/usr/local/ag_grouper"
	}

}
