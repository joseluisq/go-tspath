package tsconfig

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		configPath string
	}

	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "new instance with a ts config path",
			args: args{configPath: "./tsconfig2.json"},
			want: &Config{
				ConfigPath: "./tsconfig2.json",
				Data: &ConfigData{
					CompilerOptions: CompilerOptions{
						BaseURL: "./",
					},
				},
			},
		},
		{
			name: "new instance without a ts config path",
			want: &Config{
				ConfigPath: "./tsconfig.json",
				Data: &ConfigData{
					CompilerOptions: CompilerOptions{
						BaseURL: "./",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.configPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Read(t *testing.T) {
	tsconf := New("./../../sample/tsconfig.replace.json")

	tests := []struct {
		name   string
		tsconf *Config
		want   ConfigData
	}{
		{
			name:   "read tsconfig.replace.json file",
			tsconf: tsconf,
			want: ConfigData{
				CompilerOptions: CompilerOptions{
					BaseURL: "./",
					Paths: map[string][]string{
						"~/*": {"src/*"},
					},
					OutDir: "./dist",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tsconf.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
