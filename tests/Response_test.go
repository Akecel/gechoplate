package tests

import (
	"gechoplate/controllers"
	"reflect"
	"testing"
)

func TestSetResponse(t *testing.T) {
	type args struct {
		code    int
		message string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
		want *controllers.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controllers.SetResponse(tt.args.code, tt.args.message, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
