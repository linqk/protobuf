package url

import (
	"net/url"
	"reflect"
	"testing"
)

func TestUserinfo(t *testing.T) {
	tests := []struct {
		name string
		user *UserInfo
		want *url.Userinfo
	}{
		{
			name: "Retrieve user info with username",
			user: &UserInfo{
				Username: "john doe",
			},
			want: url.User("john doe"),
		},
		{
			name: "Retrieve user info with username and password",
			user: &UserInfo{
				Username: "john doe",
				Password: "abc",
			},
			want: url.UserPassword("john doe", "abc"),
		},
		{
			name: "Retrieve user info with empty username and password",
			user: &UserInfo{
				Username: "",
				Password: "",
			},
			want: &url.Userinfo{},
		},
		{
			name: "Retrieve user info with nil",
			user: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userinfo(tt.user)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userinfo(%v) = %v, want %v", tt.user, got, tt.want)
			}
		})
	}
}
