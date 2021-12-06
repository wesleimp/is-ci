package isci

import (
	"os"
	"strings"
)

// Check returns true if the current environment is found to probably be a CI environment or service
func Check() bool {
	ciVar := os.Getenv("CI")
	return ciVar == "true" ||
		ciVar == "1" ||
		check("CI_NAME") ||
		check("GITHUB_ACTION") ||
		check("GITLAB_CI") ||
		check("NETLIFY") ||
		check("TRAVIS") ||
		check("CODEBUILD_SRC_DIR") ||
		check("BUILDER_OUTPUT") ||
		check("GITLAB_DEPLOYMENT") ||
		check("NOW_GITHUB_DEPLOYMENT") ||
		check("NOW_BUILDER") ||
		check("BITBUCKET_DEPLOYMENT") ||
		check("GERRIT_PROJECT") ||
		check("SYSTEM_TEAMFOUNDATIONCOLLECTIONURI") ||
		check("BITRISE_IO") ||
		check("BUDDY_WORKSPACE_ID") ||
		check("BUILDKITE") ||
		check("CIRRUS_CI") ||
		check("APPVEYOR") ||
		check("CIRCLECI") ||
		check("SEMAPHORE") ||
		check("DRONE") ||
		check("DSARI") ||
		check("TDDIUM") ||
		check("STRIDER") ||
		check("TASKCLUSTER_ROOT_URL") ||
		check("JENKINS_URL") ||
		check("bamboo.buildKey") ||
		check("GO_PIPELINE_NAME") ||
		check("HUDSON_URL") ||
		check("WERCKER") ||
		check("MAGNUM") ||
		check("NEVERCODE") ||
		check("RENDER") ||
		check("SAIL_CI") ||
		check("SHIPPABLE") ||
		strings.HasSuffix(os.Getenv("NODE"), "//heroku/node/bin/node")
}

func check(provider string) bool {
	if _, ok := os.LookupEnv(provider); ok {
		return true
	}

	return false
}
