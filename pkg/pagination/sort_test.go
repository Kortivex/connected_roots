package pagination

import (
	"testing"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/stretchr/testify/assert"
)

// function returns field name and ASC order when input has no prefix.
func TestValidateSortsNoPrefix(t *testing.T) {
	// Given
	input := "name"
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "name", field)
	assert.Equal(t, paginator.ASC, *order)
}

// function returns field name and ASC order when input has '+' prefix.
func TestValidateSortsPlusPrefix(t *testing.T) {
	// Given
	input := "+name"
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "name", field)
	assert.Equal(t, paginator.ASC, *order)
}

// function returns field name and DESC order when input has '-' prefix.
func TestValidateSortsMinusPrefix(t *testing.T) {
	// Given
	input := "-name"
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "name", field)
	assert.Equal(t, paginator.DESC, *order)
}

// function trims whitespace around the input field name before processing.
func TestValidateSortsTrimWhitespace(t *testing.T) {
	// Given
	input := "  name  "
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "name", field)
	assert.Equal(t, paginator.ASC, *order)
}

// function accepts and correctly processes a valid field name from the list of validateFields.
func TestValidateSortsValidField(t *testing.T) {
	// Given
	input := "age"
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "age", field)
	assert.Equal(t, paginator.ASC, *order)
}

// function returns an error when input field name is not in validateFields.
func TestValidateSortsInvalidField(t *testing.T) {
	// Given
	input := "height"
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Empty(t, field)
}

// function returns an error when input field name is empty or whitespace only.
func TestValidateSortsEmptyField(t *testing.T) {
	// Given
	input := "   "
	validFields := []string{"name", "age"}

	// When
	field, order, err := validateSorts(input, validFields)

	// Then
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Empty(t, field)
}

// function correctly identifies and processes field names with mixed case sensitivity if present in validateFields.
func TestValidateSortsMixedCaseSensitivity(t *testing.T) {
	// Given
	validateFields := []string{"field1", "Field2", "fIeLd3"}
	value := "+Field2"

	// When
	field, order, err := validateSorts(value, validateFields)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "Field2", field)
	assert.Equal(t, paginator.ASC, *order)
}

////

// Sorts are validated and transformed into paginator rules when all inputs are valid.
func TestValidSortsTransformedToRules(t *testing.T) {
	// Given
	builder := SortRulesBuilder{
		Sorts:               []string{"+name", "-date"},
		ValidateFields:      []string{"name", "date"},
		DBStructAssociation: map[string]string{"name": "Name", "date": "Date"},
		TableName:           "users",
	}

	// When
	rules, err := builder.ObtainRules()

	// Then
	assert.NoError(t, err)
	assert.Len(t, rules, 2)
	assert.Equal(t, "Name", rules[0].Key)
	assert.Equal(t, paginator.ASC, rules[0].Order)
	assert.Equal(t, "Date", rules[1].Key)
	assert.Equal(t, paginator.DESC, rules[1].Order)
}

// TableName suffix is correctly adjusted to include a dot if missing.
func TestTableNameSuffixDotAdjustment(t *testing.T) {
	// Given
	builder := SortRulesBuilder{
		Sorts:               []string{"name"},
		ValidateFields:      []string{"name"},
		DBStructAssociation: map[string]string{"name": "Name"},
		TableName:           "users",
	}

	// When
	rules, err := builder.ObtainRules()

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "users.", rules[0].SQLRepr[:6])
}

// Empty Sorts list results in an empty rules list without error.
func TestEmptySortsListResultsInEmptyRulesList(t *testing.T) {
	// Given
	builder := SortRulesBuilder{
		Sorts:               []string{},
		ValidateFields:      []string{"name", "date"},
		DBStructAssociation: map[string]string{"name": "Name", "date": "Date"},
		TableName:           "users",
	}

	// When
	rules, err := builder.ObtainRules()

	// Then
	assert.NoError(t, err)
	assert.Empty(t, rules)
}

// TestNilDBStructAssociationNotAllow a nill DBStructAssociation result in an error.
func TestNilDBStructAssociationNotAllow(t *testing.T) {
	// Given
	builder := SortRulesBuilder{
		Sorts:               []string{"-name", "+date"},
		ValidateFields:      []string{"name", "date"},
		DBStructAssociation: nil,
		TableName:           "users",
	}

	// When
	rules, err := builder.ObtainRules()

	// Then
	assert.Error(t, err)
	assert.Empty(t, rules)
}
