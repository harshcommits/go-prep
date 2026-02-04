package misc

import (
	"strings"
)

type ReleaseType struct {
	Web     Version
	Desktop Version
	Agent   Version
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func CalculateNextVersions(currentVersion ReleaseType, commits []string) ReleaseType {

	var releaseType string
	var platform string

	for _, commit := range commits {

		// Parse "type(platform): message" format
		typeEnd := strings.Index(commit, "(")
		platformEnd := strings.Index(commit, ")")

		if typeEnd != -1 && platformEnd != -1 && typeEnd < platformEnd {
			releaseType = commit[:typeEnd]
			platform = commit[typeEnd+1 : platformEnd]
		}

		switch releaseType {
		case "feat":
			if platform == "web" {
				currentVersion.Web.Minor += 1
			} else if platform == "desktop" {
				currentVersion.Desktop.Minor += 1
			} else if platform == "core" {
				currentVersion.Desktop.Minor += 1
				currentVersion.Agent.Minor += 1
				currentVersion.Web.Minor += 1
			}
		case "fix":
			if platform == "web" {
				currentVersion.Web.Patch += 1
			} else if platform == "desktop" {
				currentVersion.Desktop.Patch += 1
			} else if platform == "core" {
				currentVersion.Desktop.Patch += 1
				currentVersion.Agent.Patch += 1
				currentVersion.Web.Patch += 1
			}
		}

	}

	return currentVersion

}
