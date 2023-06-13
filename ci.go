package npm

import (
	"os/exec"

	zen_targets "github.com/zen-io/zen-core/target"
)

type NpmCiConfig struct {
	zen_targets.BuildFields
}

func (ncc NpmCiConfig) GetTargets(tcc *zen_targets.TargetConfigContext) ([]*zen_targets.Target, error) {
	return []*zen_targets.Target{
		zen_targets.NewTarget(
			ncc.Name,
			zen_targets.WithSrcs(map[string][]string{"_srcs": ncc.Srcs}),
			zen_targets.WithOuts([]string{"node_modules"}),
			zen_targets.WithPassEnv(ncc.PassEnv),
			zen_targets.WithEnvVars(ncc.Env),
			zen_targets.WithVisibility(ncc.Visibility),
			zen_targets.WithTargetScript("build", &zen_targets.TargetScript{
				Deps: ncc.Deps,
				Run: func(target *zen_targets.Target, runCtx *zen_targets.RuntimeContext) error {
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
