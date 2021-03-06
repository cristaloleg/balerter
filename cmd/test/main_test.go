package main

import "testing"

func Test_validateLogLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "level ERROR",
			args: args{
				level: "ERROR",
			},
			wantErr: false,
		},
		{
			name: "level WARN",
			args: args{
				level: "WARN",
			},
			wantErr: false,
		},
		{
			name: "level INFO",
			args: args{
				level: "INFO",
			},
			wantErr: false,
		},
		{
			name: "level DEBUG",
			args: args{
				level: "DEBUG",
			},
			wantErr: false,
		},
		{
			name: "level wrong",
			args: args{
				level: "wrong",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateLogLevel(tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("validateLogLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
