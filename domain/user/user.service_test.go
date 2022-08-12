package user

import "testing"

func TestUser_Login(t *testing.T) {
	type fields struct {
		Id       int64
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
				Id: 100,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTAwIiwiaXNfYWRtaW4iOmZhbHNlLCJlbWFpbCI6InRlc3RAYXNrZS5jby5rciJ9.6ZcNPBX0kGfmHTCndbbt63ICX_mxyW7nCcXw3fxDo7A",
		},
		{
			"101",
			fields{
				Id: 123,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTIzIiwiaXNfYWRtaW4iOmZhbHNlLCJlbWFpbCI6InRlc3RAYXNrZS5jby5rciJ9.kS0LaBans8jHojQ6PAfXgloIReBPFQVxOOBLwGq18p0",
		},
		{
			"102",
			fields{
				Id: 11,
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTEiLCJpc19hZG1pbiI6ZmFsc2UsImVtYWlsIjoidGVzdEBhc2tlLmNvLmtyIn0.PTXUo32VfBNI0QX71vHyVklXpeO0uAvK_dEKtsxGTnk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Password: tt.fields.Password,
			}
			if got := u.Login(); got != tt.want {
				t.Errorf("\nLogin() = %v,\nwant %v", got, tt.want)
			}
		})
	}
}

func Test_getId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"1",
			"test@aske.co.kr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getId(); got != tt.want {
				t.Errorf("getId() = %v, want %v", got, tt.want)
			}
		})
	}
}
