package npm

import (
	zen_targets "github.com/zen-io/zen-core/target"
)

type TypescriptBuildConfig struct {
	Name          string            `mapstructure:"name" zen:"yes" desc:"Name for the target"`
	Description   string            `mapstructure:"desc" zen:"yes" desc:"Target description"`
	Labels        []string          `mapstructure:"labels" zen:"yes" desc:"Labels to apply to the targets"`
	Deps          []string          `mapstructure:"deps" zen:"yes" desc:"Build dependencies"`
	PassEnv       []string          `mapstructure:"pass_env" zen:"yes" desc:"List of environment variable names that will be passed from the OS environment, they are part of the target hash"`
	PassSecretEnv []string          `mapstructure:"secret_env" zen:"yes" desc:"List of environment variable names that will be passed from the OS environment, they are not used to calculate the target hash"`
	Env           map[string]string `mapstructure:"env" zen:"yes" desc:"Key-Value map of static environment variables to be used"`
	Tools         map[string]string `mapstructure:"tools" zen:"yes" desc:"Key-Value map of tools to include when executing this target. Values can be references"`
	Visibility    []string          `mapstructure:"visibility" zen:"yes" desc:"List of visibility for this target"`
	Srcs          []string          `mapstructure:"srcs" desc:"Sources for the build"`
	Out           string            `mapstructure:"out"`
}

func (tsc TypescriptBuildConfig) GetTargets(tcc *zen_targets.TargetConfigContext) ([]*zen_targets.TargetBuilder, error) {

	tb := zen_targets.ToTarget(tsc)
	tb.Srcs = map[string][]string{"_srcs": tsc.Srcs}
	tb.Outs = []string{tsc.Out}

	tb.Scripts["build"] = &zen_targets.TargetBuilderScript{
		Deps: tsc.Deps,
		Run: func(target *zen_targets.Target, runCtx *zen_targets.RuntimeContext) error {
			return target.Exec([]string{"tsc", "-p", "."}, "node build")
		},
	}
	return []*zen_targets.TargetBuilder{tb}, nil
}
