package http

import (
	"fmt"
	"strings"
)

const (
	noteKey                  = "note"
	pageNote                 = "page should be an integer number greater than zero"
	perPageNote              = "perPage should be an integer number equal or greater than zero. use zero as unlimited"
	emailAddressNote         = "example: john@doe.com"
	passwordResetLinkURLNote = "example: https://example.com/reset-password/{TOKEN}" //nolint:gosec
	accessTokenScopeNote     = "token doesn't have access token scope"
)

func includesNote(validIncludes ...string) string {
	return fmt.Sprintf("valid include samples for this route are: %s", strings.Join(validIncludes, ", "))
}

const (
	includesBuilding     = "building"
	includesUnit         = "unit"
	includesOwner        = "owner"
	includesPayer        = "payer"
	includesBill         = "bill"
	includesBulkBill     = "bulkBill"
	includesPreviousBill = "previousBill"
)

const (
	xPaginationCurrentPageHeader = "X-Pagination-Current-Page"
	xPaginationPageCountHeader   = "X-Pagination-Page-Count"
	xPaginationPerPageHeader     = "X-Pagination-Per-Page"
	xPaginationTotalCountHeader  = "X-Pagination-Total-Count"
)
