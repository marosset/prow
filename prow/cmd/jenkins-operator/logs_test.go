/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "testing"

func Test_getRealJenkinsLogPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{
			name: "flatten job",
			args: args{path: "/job/abc/1/consoleText"},
			want: "/job/abc/1/consoleText",
		},
		{
			name: "nested job",
			args: args{path: "job/folder-l1/foleer-L2/the-job/1/consoleText"},
			want: "job/folder-l1/job/foleer-L2/job/the-job/1/consoleText",
		},
		{
			name: "nested job with lead splash",
			args: args{path: "/job/folder-l1/foleer-L2/the-job/1/consoleText"},
			want: "/job/folder-l1/job/foleer-L2/job/the-job/1/consoleText",
		},
		{
			name:    "invalid nested job",
			args:    args{path: "job/.folder-l1/foleer-L2/the-job./1/consoleText"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRealJenkinsLogPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRealJenkinsLogPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRealJenkinsLogPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
