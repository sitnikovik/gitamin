package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFrom(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				path: fromRoot("lintcommit.example.yaml"),
			},
			want: Config{
				Branch: Branch{
					Default: "master",
				},
				Type: Type{
					List: []string{
						"feat",
						"patch",
						"fix",
						"chore",
						"refactor",
						"perf",
						"test",
						"docs",
						"build",
						"ci",
					},
					Required: true,
				},
				Scope: Scope{
					Pattern:  "^[A-Za-z _-]+$",
					Required: true,
				},
				Subject: Subject{
					MinLength: 10,
					MaxLength: 72,
				},
				Body: Body{
					RequiredForTypes: []string{
						"feat",
						"fix",
						"refactor",
						"perf",
						"docs",
					},
					MinLength:                 20,
					MaxLength:                 1000,
					RequiredForBreakingChange: true,
					RequireBlankLine:          true,
				},
				Task: Task{
					Pattern:       `(TASK|PROJ|BUG)-[0-9]+`,
					BranchPattern: `feature/(TASK|PROJ|BUG)-[0-9]+`,
					Required:      true,
					Location: []string{
						"subject",
						"branch",
					},
				},
				Forbidden: Forbidden{
					Words: []string{
						"WIP",
						"temp",
						"test commit",
						"debug",
						"fix bug",
						"quick fix",
						"update code",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error on unexisting file",
			args: args{
				path: "unknown_file.yaml",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "error on marshaling failed by invalid yaml",
			args: args{
				path: fromRoot("README.md"),
			},
			want:    Config{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFrom(tt.args.path)

			require.Equal(
				t,
				tt.wantErr,
				err != nil,
				"unexpected error: %v",
				err,
			)
			require.Equal(
				t,
				tt.want,
				got,
				"unexpected result",
			)
		})
	}
}

// fromRoot returns the path to the root directory of the project.
func fromRoot(path string) string {
	return "../../" + path
}
