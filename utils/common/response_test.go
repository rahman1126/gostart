package common

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetErrorMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				err: nil,
			},
			want: "Success",
		},
		{
			args: args{
				err: ErrNotFound,
			},
			want: "Not Found",
		},
		{
			args: args{
				err: ErrUnauthorized,
			},
			want: "Unauthorized",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetErrorMessage(tt.args.err); got != tt.want {
				t.Errorf("GetErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				err: errors.New("Not Found"),
			},
			want: 404,
		},
		{
			args: args{
				err: ErrUnauthorized,
			},
			want: 401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStatusCode(tt.args.err); got != tt.want {
				t.Errorf("GetStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		data interface{}
		err  error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				data: nil,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func BenchmarkGetStatusCode(b *testing.B) {
	for i:=0; i<b.N; i++ {
		GetStatusCode(errors.New("Test"))
	}
}

func BenchmarkGetErrorMessage(b *testing.B) {
	for i:=0; i<b.N; i++ {
		GetErrorMessage(errors.New("Halo"))
	}
}