package isci

import (
	"os"
	"testing"
)

func TestIsCI(t *testing.T) {
	var cc = []struct {
		provider string
		expected bool
	}{
		{provider: "GERRIT_PROJECT", expected: true},
		{provider: "GITLAB_CI", expected: true},
		{provider: "CIRCLECI", expected: true},
		{provider: "SEMAPHORE", expected: true},
		{provider: "DRONE", expected: true},
		{provider: "GITHUB_ACTION", expected: true},
		{provider: "TDDIUM", expected: true},
		{provider: "JENKINS_URL", expected: true},
		{provider: "WERCKER", expected: true},
		{provider: "NETLIFY", expected: true},
		{provider: "NOW_GITHUB_DEPLOYMENT", expected: true},
		{provider: "GITLAB_DEPLOYMENT", expected: true},
		{provider: "BITBUCKET_DEPLOYMENT", expected: true},
		{provider: "bamboo.buildKey", expected: true},
		{provider: "GO_PIPELINE_NAME", expected: true},
		{provider: "CI_NAME", expected: true},
		{provider: "TRAVIS", expected: true},
		{provider: "", expected: false},
		{provider: "APPVEYOR", expected: true},
		{provider: "CODEBUILD_SRC_DIR", expected: true},
		{provider: "BUILDER_OUTPUT", expected: true},
		{provider: "BITBUCKET_BUILD_NUMBER", expected: true},
		{provider: "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI", expected: true},
		{provider: "BITRISE_IO", expected: true},
		{provider: "BUDDY_WORKSPACE_ID", expected: true},
		{provider: "BUILDKITE", expected: true},
		{provider: "CIRRUS_CI", expected: true},
		{provider: "DSARI", expected: true},
		{provider: "STRIDER", expected: true},
		{provider: "TASKCLUSTER_ROOT_URL", expected: true},
		{provider: "HUDSON_URL", expected: true},
		{provider: "NOW_BUILDER", expected: true},
		{provider: "VERCEL_URL", expected: true},
		{provider: "MAGNUM", expected: true},
		{provider: "NEVERCODE", expected: true},
		{provider: "RENDER", expected: true},
		{provider: "SAIL_CI", expected: true},
		{provider: "SHIPPABLE", expected: true},
		{provider: "TEAMCITY_VERSION", expected: true},
		{provider: "NODE", expected: true},
	}

	for _, c := range cc {
		t.Run(c.provider, func(t *testing.T) {
			if c.provider != "" {
				env(t, "CI", "1")
				env(t, c.provider, "1")
			}

			res := Check()
			if res != c.expected {
				t.Errorf("Provider %s got: %v; Expected: %v", c.provider, res, c.expected)
			}
		})
	}
}

func env(t *testing.T, env, val string) {
	orig := os.Getenv(env)
	t.Cleanup(func() {
		if orig != "" {
			os.Setenv(env, orig)
		} else {
			os.Unsetenv(env)
		}
	})

	if val != "" {
		os.Setenv(env, val)
	} else {
		os.Unsetenv(env)
	}
}
