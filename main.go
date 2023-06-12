package npm

import (
	ahoy_targets "gitlab.com/hidothealth/platform/ahoy/src/target"
)

var KnownTargets = ahoy_targets.TargetCreatorMap{
	"npm_ci":           NpmCiConfig{},
	"typescript_build": TypescriptBuildConfig{},
}
