package npm

import (
	"os/exec"

	zen_targets "github.com/zen-io/zen-core/target"
)

type TypescriptBuildConfig struct {
	Out string `mapstructure:"out"`
	zen_targets.BuildFields
}

func (tsc TypescriptBuildConfig) GetTargets(tcc *zen_targets.TargetConfigContext) ([]*zen_targets.Target, error) {
	return []*zen_targets.Target{
		zen_targets.NewTarget(
			tsc.Name,
			zen_targets.WithSrcs(map[string][]string{"_srcs": tsc.Srcs}),
			zen_targets.WithOuts([]string{tsc.Out}),
			zen_targets.WithPassEnv(tsc.PassEnv),
			zen_targets.WithEnvVars(tsc.Env),
			zen_targets.WithVisibility(tsc.Visibility),
			zen_targets.WithTargetScript("build", &zen_targets.TargetScript{
				Deps: tsc.Deps,
				Run: func(target *zen_targets.Target, runCtx *zen_targets.RuntimeContext) error {
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
