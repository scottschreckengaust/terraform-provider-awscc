// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

var (
	_ basetypes.ListTypable                    = (*multisetTypeOf[basetypes.StringType])(nil)
	_ basetypes.ListValuable                   = (*MultisetValueOf[basetypes.StringType])(nil)
	_ basetypes.ListValuableWithSemanticEquals = (*MultisetValueOf[basetypes.StringType])(nil)
)

// A multiset is an array allowing non-unique items with insertion order not significant.
// Multisets do not correspond directly with either Terraform Lists (insertion order is significant) or Sets (unique items).
// Multiset Attributes are declared as Lists with a custom type.

type multisetTypeOf[T attr.Type] struct {
	basetypes.ListType
}

func NewMultisetTypeOf[T attr.Type](ctx context.Context) multisetTypeOf[T] {
	var zero T
	return multisetTypeOf[T]{basetypes.ListType{ElemType: zero}}
}

func (t multisetTypeOf[T]) Equal(o attr.Type) bool {
	other, ok := o.(multisetTypeOf[T])

	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (multisetTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("MultisetTypeOf[%T]", zero)
}

func (t multisetTypeOf[T]) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics
	var zero T

	if in.IsNull() {
		return NewMultisetValueOfNull[T](ctx), diags
	}

	if in.IsUnknown() {
		return NewMultisetValueOfUnknown[T](ctx), diags
	}

	v, d := basetypes.NewListValue(zero, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewMultisetValueOfUnknown[T](ctx), diags
	}

	return MultisetValueOf[T]{ListValue: v}, diags
}

func (t multisetTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, listValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

func (multisetTypeOf[T]) ValueType(context.Context) attr.Value {
	return MultisetValueOf[T]{}
}

type MultisetValueOf[T attr.Type] struct {
	basetypes.ListValue
}

func (v MultisetValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(MultisetValueOf[T])

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (MultisetValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewMultisetTypeOf[T](ctx)
}

func (v MultisetValueOf[T]) ListSemanticEquals(ctx context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(MultisetValueOf[T])
	if !ok {
		return false, diags
	}

	old, d := v.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	new, d := newValue.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	oldElems, newElems := old.Elements(), new.Elements()

	if len(oldElems) != len(newElems) {
		return false, diags
	}

	for _, newElem := range newElems {
		found := false
		for i, oldElem := range oldElems {
			if oldElem.Equal(newElem) {
				oldElems = slices.Delete(oldElems, i, i+1)
				found = true
				break
			}
		}
		if !found {
			return false, diags
		}
	}

	return len(oldElems) == 0, diags
}

func NewMultisetValueOfNull[T attr.Type](ctx context.Context) MultisetValueOf[T] {
	var zero T
	return MultisetValueOf[T]{ListValue: basetypes.NewListNull(zero)}
}

func NewMultisetValueOfUnknown[T attr.Type](ctx context.Context) MultisetValueOf[T] {
	var zero T
	return MultisetValueOf[T]{ListValue: basetypes.NewListUnknown(zero)}
}

func NewMultisetValueOf[T attr.Type](ctx context.Context, elements []attr.Value) (MultisetValueOf[T], diag.Diagnostics) {
	var diags diag.Diagnostics
	var zero T

	v, d := basetypes.NewListValue(zero, elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewMultisetValueOfUnknown[T](ctx), diags
	}

	return MultisetValueOf[T]{ListValue: v}, diags
}

func NewMultisetValueOfMust[T attr.Type](ctx context.Context, elements []attr.Value) MultisetValueOf[T] {
	return ccdiag.Must(NewMultisetValueOf[T](ctx, elements))
}
