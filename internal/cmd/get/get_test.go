package get

import (
	"github.com/g2a-com/klio/internal/context"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_initialiseProjectInCurrentDir(t *testing.T) {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("can't get current directory: %s", err)
	}

	projetConfigFileName := "test-config-name.yaml"
	installDirName := "test-dir"

	type args struct {
		ctx context.CLIContext
	}
	tests := []struct {
		name    string
		args    args
		want    context.CLIContext
		wantErr bool
	}{
		{
			name: "should initialise default klio config file and update context paths",
			args: args{
				ctx: struct {
					Config context.CLIConfig
					Paths  context.Paths
				}{
					Config: context.CLIConfig{
						ProjectConfigFileName: projetConfigFileName,
						InstallDirName:        installDirName,
					},
					Paths: struct {
						ProjectConfigFile string
						ProjectInstallDir string
						GlobalInstallDir  string
					}{}},
			},
			want: context.CLIContext{
				Config: context.CLIConfig{
					ProjectConfigFileName: projetConfigFileName,
					InstallDirName:        installDirName,
				},
				Paths: context.Paths{
					ProjectConfigFile: filepath.Join(currentWorkingDirectory, projetConfigFileName),
					ProjectInstallDir: filepath.Join(currentWorkingDirectory, installDirName),
					GlobalInstallDir:  "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initialiseProjectInCurrentDir(tt.args.ctx)
			defer os.RemoveAll(got.Paths.GlobalInstallDir)
			defer os.RemoveAll(got.Paths.ProjectConfigFile)

			if (err != nil) != tt.wantErr {
				t.Errorf("initialiseProjectInCurrentDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
