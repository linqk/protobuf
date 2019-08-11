package url_test

import (
	"net/url"
	"reflect"
	"testing"

	pturl "github.com/linqk/protobuft/net/url"
)

func TestFromProto(t *testing.T) {
	tests := []struct {
		name string
		url  pturl.URL
		want url.URL
	}{
		{
			name: "A URL with scheme, host, path and raw path",
			url: pturl.URL{
				Scheme:  "http",
				Host:    "www.google.com",
				Path:    "/file one&two",
				RawPath: "/file%20one%26two",
			},
			want: url.URL{
				Scheme:  "http",
				Host:    "www.google.com",
				Path:    "/file one&two",
				RawPath: "/file%20one%26two",
			},
		},
		{
			name: "A URL with scheme, host, path, raw path and fragment",
			url: pturl.URL{
				Scheme:   "https",
				Host:     "www.google.com",
				Path:     "/",
				RawQuery: "q=go+language",
				Fragment: "foo&bar",
			},
			want: url.URL{
				Scheme:   "https",
				Host:     "www.google.com",
				Path:     "/",
				RawQuery: "q=go+language",
				Fragment: "foo&bar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pturl.FromProto(tt.url)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromProto(%v) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}

func TestToProto(t *testing.T) {
	tests := []struct {
		name string
		url  url.URL
		want pturl.URL
	}{
		{
			name: "A URL with scheme, host, path and raw path",
			url: url.URL{
				Scheme:  "http",
				Host:    "www.google.com",
				Path:    "/file one&two",
				RawPath: "/file%20one%26two",
			},
			want: pturl.URL{
				Scheme:  "http",
				Host:    "www.google.com",
				Path:    "/file one&two",
				RawPath: "/file%20one%26two",
			},
		},
		{
			name: "A URL with scheme, host, path, raw path and fragment",
			url: url.URL{
				Scheme:   "https",
				Host:     "www.google.com",
				Path:     "/",
				RawQuery: "q=go+language",
				Fragment: "foo&bar",
			},
			want: pturl.URL{
				Scheme:   "https",
				Host:     "www.google.com",
				Path:     "/",
				RawQuery: "q=go+language",
				Fragment: "foo&bar",
			},
		},
		{
			name: "A URL with scheme, host, path, and username",
			url: url.URL{
				Scheme: "https",
				Host:   "www.google.com",
				Path:   "/",
				User:   url.User("john doe"),
			},
			want: pturl.URL{
				Scheme: "https",
				Host:   "www.google.com",
				Path:   "/",
				User: &pturl.UserInfo{
					Username:    "john doe",
					Password:    "",
					PasswordSet: false,
				},
			},
		},
		{
			name: "A URL with scheme, host, path, username and password",
			url: url.URL{
				Scheme: "https",
				Host:   "www.google.com",
				Path:   "/",
				User:   url.UserPassword("john doe", "abc"),
			},
			want: pturl.URL{
				Scheme: "https",
				Host:   "www.google.com",
				Path:   "/",
				User: &pturl.UserInfo{
					Username:    "john doe",
					Password:    "abc",
					PasswordSet: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pturl.ToProto(tt.url)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToProto(%v) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}
