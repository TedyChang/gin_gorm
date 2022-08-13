package user

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUser_Login(t *testing.T) {
	type fields struct {
		ID       uint
		Name     string
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"S > login 100",
			fields{
				ID: 100,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTAwIiwiaXNfYWRtaW4iOmZhbHNlLCJ1c2VyX2lkIjoxMDB9.cT5r-erbnVfxu0MoB6bAXcpFh3gCifnt-KQIaGhk9GM",
		},
		{
			"S > login 123",
			fields{
				ID: 123,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTIzIiwiaXNfYWRtaW4iOmZhbHNlLCJ1c2VyX2lkIjoxMjN9.tztEiPbPyUapPllQDfnD3_hg6Yxr9nYJUebVwQ5UN3Y",
		},
		{
			"S > login 11",
			fields{
				ID: 11,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTEiLCJpc19hZG1pbiI6ZmFsc2UsInVzZXJfaWQiOjExfQ.ByYFfzhAQ6Y-XwNY9BP4RXns17eSo8zty9V9qU1OgzA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Model: gorm.Model{
					ID: tt.fields.ID,
				},
				Name:     tt.fields.Name,
				Password: tt.fields.Password,
			}
			if got := createToken(u); got != tt.want {
				t.Errorf("\nLogin() = %v,\nwant %v", got, tt.want)
			}
		})
	}
}

func Test_getId(t *testing.T) {
	type args struct {
		strToken string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"S / token get Id",
			args{strToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTAwIiwiaXNfYWRtaW4iOmZhbHNlLCJ1c2VyX2lkIjoxMDB9.cT5r-erbnVfxu0MoB6bAXcpFh3gCifnt-KQIaGhk9GM"},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserID(tt.args.strToken); got != tt.want {
				t.Errorf("getUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateSave(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 error
	}{
		{
			"S / validate user save",
			args{u: User{
				Name:     "test",
				Password: "pw1",
			}},
			nil,
			nil,
		},
		{
			"F / empty name",
			args{u: User{
				Name:     "",
				Password: "pw1",
			}},
			errors.New("empty name"),
			nil,
		},
		{
			"F / empty pw",
			args{u: User{
				Name:     "test",
				Password: "",
			}},
			nil,
			errors.New("empty pw"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateSave(tt.args.u)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateSave() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateSave() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
