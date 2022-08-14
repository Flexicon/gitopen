package main

import (
	"errors"
	"testing"
)

func TestRepoService_GetRepositoryURL(t *testing.T) {
	type fields struct {
		cmd Commander
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "https remote",
			fields: fields{
				cmd: TestCommander{Output: "https://github.com/Flexicon/gitopen.git", Error: nil},
			},
			want:    "https://github.com/Flexicon/gitopen",
			wantErr: false,
		},
		{
			name: "git remote",
			fields: fields{
				cmd: TestCommander{Output: "git@github.com:Flexicon/gitopen.git", Error: nil},
			},
			want:    "https://github.com/Flexicon/gitopen",
			wantErr: false,
		},
		{
			name: "non-github git remote",
			fields: fields{
				cmd: TestCommander{Output: "git@gitlab.com:gitlab-org/ruby/gems/gitlab-styles.git", Error: nil},
			},
			want:    "https://gitlab.com/gitlab-org/ruby/gems/gitlab-styles",
			wantErr: false,
		},
		{
			name: "ssh remote",
			fields: fields{
				cmd: TestCommander{Output: "ssh://git@github.com/Flexicon/gitopen", Error: nil},
			},
			want:    "https://github.com/Flexicon/gitopen",
			wantErr: false,
		},
		{
			name: "ssh remote - longer",
			fields: fields{
				cmd: TestCommander{Output: "ssh://git@github.com/Homebrew/homebrew-core", Error: nil},
			},
			want:    "https://github.com/Homebrew/homebrew-core",
			wantErr: false,
		},
		{
			name: "invalid remote",
			fields: fields{
				cmd: TestCommander{Output: "i-am-invalid", Error: nil},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "command error",
			fields: fields{
				cmd: TestCommander{Output: "", Error: errors.New("command error")},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RepoService{
				cmd: tt.fields.cmd,
			}
			got, err := s.GetRepositoryURL("origin")
			if (err != nil) != tt.wantErr {
				t.Errorf("RepoService.GetRepositoryURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RepoService.GetRepositoryURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
