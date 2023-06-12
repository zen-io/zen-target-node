package npm

import (
	"os/exec"

	ahoy_targets "gitlab.com/hidothealth/platform/ahoy/src/target"
)

type NpmCiConfig struct {
	ahoy_targets.BuildFields
}

func (ncc NpmCiConfig) GetTargets(tcc *ahoy_targets.TargetConfigContext) ([]*ahoy_targets.Target, error) {
	return []*ahoy_targets.Target{
		ahoy_targets.NewTarget(
			ncc.Name,
			ahoy_targets.WithSrcs(map[string][]string{"_srcs": ncc.Srcs}),
			ahoy_targets.WithOuts([]string{"node_modules"}),
			ahoy_targets.WithPassEnv(ncc.PassEnv),
			ahoy_targets.WithEnvVars(ncc.Env),
			ahoy_targets.WithVisibility(ncc.Visibility),
			ahoy_targets.WithTargetScript("build", &ahoy_targets.TargetScript{
				Deps: ncc.Deps,
				Run: func(target *ahoy_targets.Target, runCtx *ahoy_targets.RuntimeContext) error {
					cmd := exec.Command("npm", "ci")
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
