package gee

import (
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *router
	}{
		{
			name: "t1",
			want: NewRouter(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_router_addRoute(t *testing.T) {
	r := NewRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
}

func Test_router_getRoute(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
		want1  map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			got, got1 := r.getRoute(tt.args.method, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
