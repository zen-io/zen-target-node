package npm

import (
	zen_targets "github.com/zen-io/zen-core/target"
)

var KnownTargets = zen_targets.TargetCreatorMap{
	"npm_ci":           NpmCiConfig{},
	"typescript_build": TypescriptBuildConfig{},
}
