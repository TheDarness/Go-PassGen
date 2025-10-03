package generator

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		opts    Options
		wantErr bool
	}{
		{
			name: "Only lowercase",
			opts: Options{Length: 10, Lowercase: true},
		},
		{
			name: "All categories",
			opts: Options{Length: 12, Lowercase: true, Uppercase: true, Digits: true, Specials: true},
		},
		{
			name:    "No categories",
			opts:    Options{Length: 10},
			wantErr: true,
		},
		{
			name:    "Length too short",
			opts:    Options{Length: 2, Lowercase: true, Uppercase: true},
			wantErr: true,
		},
		{
			name:    "Invalid length too small",
			opts:    Options{Length: 3, Lowercase: true},
			wantErr: true,
		},
		{
			name:    "Invalid length too large",
			opts:    Options{Length: 129, Lowercase: true},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := Generate(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if len(password) != tt.opts.Length {
					t.Errorf("Generate() password length = %d, want %d", len(password), tt.opts.Length)
				}
			}
		})
	}
}
