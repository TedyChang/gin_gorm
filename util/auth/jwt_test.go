package auth

import "testing"

func TestCreateToken(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"S > create token",
			args{id: 100},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTAwIiwiaXNfYWRtaW4iOmZhbHNlLCJ1c2VyX2lkIjoxMDB9.cT5r-erbnVfxu0MoB6bAXcpFh3gCifnt-KQIaGhk9GM",
			false,
		},
		{
			"S > create token",
			args{id: 123},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTIzIiwiaXNfYWRtaW4iOmZhbHNlLCJ1c2VyX2lkIjoxMjN9.tztEiPbPyUapPllQDfnD3_hg6Yxr9nYJUebVwQ5UN3Y",
			false,
		},
		{
			"S > create token",
			args{id: 11},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTEiLCJpc19hZG1pbiI6ZmFsc2UsInVzZXJfaWQiOjExfQ.ByYFfzhAQ6Y-XwNY9BP4RXns17eSo8zty9V9qU1OgzA",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
