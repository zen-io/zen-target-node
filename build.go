package npm

import (
	"os/exec"

	ahoy_targets "gitlab.com/hidothealth/platform/ahoy/src/target"
)

type TypescriptBuildConfig struct {
	Out string `mapstructure:"out"`
	ahoy_targets.BuildFields
}

func (tsc TypescriptBuildConfig) GetTargets(tcc *ahoy_targets.TargetConfigContext) ([]*ahoy_targets.Target, error) {
	return []*ahoy_targets.Target{
		ahoy_targets.NewTarget(
			tsc.Name,
			ahoy_targets.WithSrcs(map[string][]string{"_srcs": tsc.Srcs}),
			ahoy_targets.WithOuts([]string{tsc.Out}),
			ahoy_targets.WithPassEnv(tsc.PassEnv),
			ahoy_targets.WithEnvVars(tsc.Env),
			ahoy_targets.WithVisibility(tsc.Visibility),
			ahoy_targets.WithTargetScript("build", &ahoy_targets.TargetScript{
				Deps: tsc.Deps,
				Run: func(target *ahoy_targets.Target, runCtx *ahoy_targets.RuntimeContext) error {
					cmd := exec.Command("tsc", "-p", ".")
					cmd.Dir = target.Cwd
					cmd.Env = target.GetEnvironmentVariablesList()
					cmd.Stdout = target
					cmd.Stderr = target
					return cmd.Run()
				},
			}),
		),
	}, nil
}
