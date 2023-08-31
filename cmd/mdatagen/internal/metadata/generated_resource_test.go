// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetMapResourceAttr(map[string]any{"key1": "map.resource.attr-val1", "key2": "map.resource.attr-val2"})
			rb.SetOptionalResourceAttr("optional.resource.attr-val")
			rb.SetSliceResourceAttr([]any{"slice.resource.attr-item1", "slice.resource.attr-item2"})
			rb.SetStringEnumResourceAttrOne()
			rb.SetStringResourceAttr("string.resource.attr-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 4, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 5, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("map.resource.attr")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, map[string]any{"key1": "map.resource.attr-val1", "key2": "map.resource.attr-val2"}, val.Map().AsRaw())
			}
			val, ok = res.Attributes().Get("optional.resource.attr")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "optional.resource.attr-val", val.Str())
			}
			val, ok = res.Attributes().Get("slice.resource.attr")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, []any{"slice.resource.attr-item1", "slice.resource.attr-item2"}, val.Slice().AsRaw())
			}
			val, ok = res.Attributes().Get("string.enum.resource.attr")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "one", val.Str())
			}
			val, ok = res.Attributes().Get("string.resource.attr")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "string.resource.attr-val", val.Str())
			}
		})
	}
}