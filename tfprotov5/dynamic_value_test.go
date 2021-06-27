package tfprotov5

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestNewDynamicValue(t *testing.T) {
	tests := map[string]struct {
		t       tftypes.Type
		v       interface{}
		want    DynamicValue
		wantErr bool
	}{
		"optional-attributes": {
			t: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"id":   tftypes.Number,
					"name": tftypes.String,
				},
				OptionalAttributes: map[string]struct{}{
					"name": {},
				},
			},
			v: map[string]tftypes.Value{
				"id": tftypes.NewValue(tftypes.Number, 1),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			v := tftypes.NewValue(tt.t, tt.v)
			got, err := NewDynamicValue(tt.t, v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDynamicValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			unmarshaled, err := got.Unmarshal(tt.t)
			if err != nil {
				t.Errorf("Unmarshal() error = %v", err)
				return
			}
			if !unmarshaled.Equal(v) {
				t.Errorf("Unmarshal() = %v, want %v", unmarshaled, v)
			}
		})
	}
}
