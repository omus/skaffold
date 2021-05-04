/*
Copyright 2021 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inspect

import (
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/parser"
	latest_v1 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest/v1"
)

// Options holds flag values for the various `skaffold inspect` commands
type Options struct {
	// Filename is the `skaffold.yaml` file path
	Filename string
	// OutFormat is the output format. One of: json
	OutFormat string
	// Modules is the module filter for specific commands
	Modules []string

	ProfilesOptions
}

// ProfilesOptions holds flag values for various `skaffold inspect profiles` commands
type ProfilesOptions struct {
	// BuildEnv is the build-env filter for command output. One of: local, googleCloudBuild, cluster.
	BuildEnv BuildEnv
}

type BuildEnv string

var (
	getConfigSetFunc = parser.GetConfigSet
	BuildEnvs        = struct {
		Unspecified      BuildEnv
		Local            BuildEnv
		GoogleCloudBuild BuildEnv
		Cluster          BuildEnv
	}{
		Local: "local", GoogleCloudBuild: "googleCloudBuild", Cluster: "cluster",
	}
)

func (b BuildEnv) MatchesConfig(t *latest_v1.BuildType) bool {
	switch b {
	case BuildEnvs.Local:
		return t.LocalBuild != nil
	case BuildEnvs.GoogleCloudBuild:
		return t.GoogleCloudBuild != nil
	case BuildEnvs.Cluster:
		return t.Cluster != nil
	case BuildEnvs.Unspecified:
		return true
	default:
		return false
	}
}
