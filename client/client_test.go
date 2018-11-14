package client

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func MockClient(apiURL string) {
	c = &client{
		apiURL:     apiURL,
		apiVersion: APIVersion,
		apiKey:     APIKey,
	}
}

var (
	s = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hi"))
	}))
)

func Test_getClient(t *testing.T) {
	tests := []struct {
		name string
		want *client
	}{
		{
			name: "OK",
			want: &client{
				apiURL:     APIURL,
				apiVersion: APIVersion,
				apiKey:     APIKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		path    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "OK",
			url:     s.URL,
			path:    "/",
			want:    []byte("Hi"),
			wantErr: false,
		},
		{
			name:    "Bad scheme",
			url:     "bad",
			path:    "/",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "No server",
			url:     "http://xxx.xxx.xxx",
			path:    "/",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Bad path",
			url:     s.URL,
			path:    "%s%s",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClientURL(tt.url)
			got, err := Get(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		path    string
		body    interface{}
		want    []byte
		wantErr bool
	}{
		{
			name:    "OK",
			url:     s.URL,
			path:    "/",
			body:    map[string]string{"hakuna": "matata"},
			want:    []byte("Hi"),
			wantErr: false,
		},
		{
			name:    "Invalid body",
			url:     s.URL,
			path:    "/",
			body:    Post,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClientURL(tt.url)
			got, err := Post(tt.path, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		path    string
		body    interface{}
		want    []byte
		wantErr bool
	}{
		{
			name:    "OK",
			url:     s.URL,
			path:    "/",
			body:    map[string]string{"hakuna": "matata"},
			want:    []byte("Hi"),
			wantErr: false,
		},
		{
			name:    "Invalid body",
			url:     s.URL,
			path:    "/",
			body:    Put,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClientURL(tt.url)
			got, err := Put(tt.path, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		path    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "OK",
			url:     s.URL,
			path:    "/",
			want:    []byte("Hi"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClientURL(tt.url)
			got, err := Delete(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
