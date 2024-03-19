package common

import (
	"fmt"
	"net/url"
	"strings"
)

type FilterOperator string

const (
	FilterOperatorContains       FilterOperator = "contains"
	FilterOperatorContainsAny    FilterOperator = "contains-any"
	FilterOperatorContainsAll    FilterOperator = "contains-all"
	FilterOperatorStartsWith     FilterOperator = "starts-with"
	FilterOperatorEndsWith       FilterOperator = "ends-with"
	FilterOperatorEquals         FilterOperator = "equals"
	FilterOperatorGreaterOrEqual FilterOperator = "greater-or-equal"
	FilterOperatorGreaterThan    FilterOperator = "greater-than"
	FilterOperatorLessOrEqual    FilterOperator = "less-or-equal"
	FilterOperatorLessThan       FilterOperator = "less-than"
	FilterOperatorAny            FilterOperator = "any"

	//Boolean operators
	FilterOperatorAND FilterOperator = "and"
	FilterOperatorOR  FilterOperator = "or"
	FilterOperatorNOT FilterOperator = "not"
)

type FilterBuilder struct {
	filters []string
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{}
}

func (builder *FilterBuilder) Equal(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%q", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorEquals, operation))
	return builder
}

func (builder *FilterBuilder) LessThan(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%q", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorLessThan, operation))

	return builder
}

func (builder *FilterBuilder) LessOrEqual(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%q", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorLessOrEqual, operation))

	return builder
}

func (builder *FilterBuilder) GreaterThan(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%q", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorGreaterThan, operation))

	return builder
}

func (builder *FilterBuilder) GreaterOrEqual(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%q", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorGreaterOrEqual, operation))

	return builder
}

func (builder *FilterBuilder) Contains(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%q", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorContains, operation))

	return builder
}

func (builder *FilterBuilder) ContainsAny(field string, values []string) *FilterBuilder {
	var valuesStr = make([]string, 0)
	for _, value := range values {
		valuesStr = append(valuesStr, fmt.Sprintf("%q", value))
	}
	operation := url.QueryEscape(fmt.Sprintf("%s,%v", field, valuesStr))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorContainsAny, operation))

	return builder
}

func (builder *FilterBuilder) ContainsAll(field string, values []string) *FilterBuilder {
	var valuesStr = make([]string, 0)
	for _, value := range values {
		valuesStr = append(valuesStr, fmt.Sprintf("%q", value))
	}

	operation := url.QueryEscape(fmt.Sprintf("%s,%v", field, valuesStr))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorContainsAll, operation))

	return builder
}

func (builder *FilterBuilder) EndsWith(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%s", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorEndsWith, operation))

	return builder
}

func (builder *FilterBuilder) StartsWith(field string, value string) *FilterBuilder {
	operation := url.QueryEscape(fmt.Sprintf("%s,%s", field, value))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorStartsWith, operation))

	return builder
}

func (builder *FilterBuilder) Any(field string, values []string) *FilterBuilder {
	var valuesStr = make([]string, 0)
	for _, value := range values {
		valuesStr = append(valuesStr, fmt.Sprintf("%q", value))
	}

	operation := url.QueryEscape(fmt.Sprintf("%s,%v", field, valuesStr))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorAny, operation))

	return builder
}

func (builder *FilterBuilder) And(fb1 FilterBuilder, fb2 FilterBuilder) *FilterBuilder {
	if len(fb1.filters) > 1 || len(fb2.filters) > 1 {
		panic("And operator must contain two operands")
	}

	operation := url.QueryEscape(fmt.Sprintf("%s,%s", fb1.build(), fb2.build()))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorAND, operation))

	return builder
}

func (builder *FilterBuilder) Or(fb1 FilterBuilder, fb2 FilterBuilder) *FilterBuilder {
	if len(fb1.filters) > 1 || len(fb2.filters) > 1 {
		panic("Or operator must contain two operands")
	}

	operation := url.QueryEscape(fmt.Sprintf("%s,%s", fb1.build(), fb2.build()))
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorOR, operation))

	return builder
}

func (builder *FilterBuilder) Not(fb FilterBuilder) *FilterBuilder {
	if len(fb.filters) > 1 {
		panic("Not operator must contain one operands")
	}

	filter := fb.build()
	builder.filters = append(builder.filters, fmt.Sprintf("%s(%s)", FilterOperatorNOT, filter))

	return builder
}

func (fb *FilterBuilder) Build() string {
	return fmt.Sprintf("filter=%v", fb.build())
}

func (fb *FilterBuilder) build() string {
	return strings.Join(fb.filters, ",")
}
