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
			"100",
			fields{
				ID: 100,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTAwIiwiaXNfYWRtaW4iOmZhbHNlLCJlbWFpbCI6InRlc3RAYXNrZS5jby5rciJ9.6ZcNPBX0kGfmHTCndbbt63ICX_mxyW7nCcXw3fxDo7A",
		},
		{
			"101",
			fields{
				ID: 123,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTIzIiwiaXNfYWRtaW4iOmZhbHNlLCJlbWFpbCI6InRlc3RAYXNrZS5jby5rciJ9.kS0LaBans8jHojQ6PAfXgloIReBPFQVxOOBLwGq18p0",
		},
		{
			"102",
			fields{
				ID: 11,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTEiLCJpc19hZG1pbiI6ZmFsc2UsImVtYWlsIjoidGVzdEBhc2tlLmNvLmtyIn0.PTXUo32VfBNI0QX71vHyVklXpeO0uAvK_dEKtsxGTnk",
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
			if got := getName(u); got != tt.want {
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
		want string
	}{
		{
			"S / token get Id",
			args{strToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTEiLCJpc19hZG1pbiI6ZmFsc2UsImVtYWlsIjoidGVzdEBhc2tlLmNvLmtyIn0.PTXUo32VfBNI0QX71vHyVklXpeO0uAvK_dEKtsxGTnk"},
			"test@aske.co.kr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEmail(tt.args.strToken); got != tt.want {
				t.Errorf("getEmail() = %v, want %v", got, tt.want)
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
