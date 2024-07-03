package pagination

import (
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

type SortRulesBuilder struct {
	// Sorts fields to be sorted by
	Sorts []string
	// ValidateFields database fields on which the sort can be applied
	ValidateFields []string
	// DBStructAssociation map to associate each DB field to the Struct
	DBStructAssociation map[string]string
	// TableName use this parameter for Join queries (indicates the table to apply the pagination)
	// not required for simple table Queries
	TableName string
}

// ObtainRules returns a list of paginator.Rule objects based on the Sorts field of the SortRulesBuilder struct.
//
// It validates the Sorts field against the ValidateFields field.
// For each value in the Sorts field, it calls the validateSorts function to get the field name and order.
// It then appends a new paginator.Rule object to the rules slice,
// with the Key field set to the corresponding value from the DBStructAssociation map,
// the Order field set to the fieldOrder value,
// and the SQLRepr field set to the concatenation of the TableName and field values.
//
// The function returns the rules slice and any error encountered during the process.
func (r *SortRulesBuilder) ObtainRules() ([]paginator.Rule, error) {
	if len(r.Sorts) > 0 && r.DBStructAssociation == nil {
		return nil, errors.New("table sort cannot be nil")
	}
	if r.TableName != "" && !strings.HasSuffix(r.TableName, ".") {
		r.TableName += "."
	}

	var rules []paginator.Rule
	for _, value := range r.Sorts {
		field, fieldOrder, err := validateSorts(value, r.ValidateFields)
		if err != nil {
			return nil, err
		}
		rules = append(rules, paginator.Rule{Key: r.DBStructAssociation[field], Order: *fieldOrder, SQLRepr: r.TableName + field})
	}
	return rules, nil
}

// ValidateSorts validates the given value against the list of valid fields
// (these fields are those belonging to the corresponding table).
//
// Parameters:
// - value: the value to be validated (string)
// - validateFields: the list of valid fields ([]string)
//
// Returns:
// - finalField: the final validated field (string)
// - finalOrder: the final order (paginator.Order pointer)
// - error: an error if the validation fails (ErrorPagination).
func validateSorts(value string, validateFields []string) (string, *paginator.Order, error) {
	finalField := strings.TrimSpace(value)
	finalOrder := paginator.ASC
	if strings.HasPrefix(finalField, "+") {
		finalField = strings.TrimPrefix(finalField, "+")
	} else if strings.HasPrefix(finalField, "-") {
		finalField = strings.TrimPrefix(finalField, "-")
		finalOrder = paginator.DESC
	}
	if !slices.Contains(validateFields, finalField) {
		return "", nil, ErrorPagination{Status: http.StatusBadRequest, Message: ErrMsgInvalidSort}
	}
	return finalField, &finalOrder, nil
}
