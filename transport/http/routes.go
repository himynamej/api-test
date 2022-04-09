package http

import (
	gohttp "net/http"
)

func (h *handler) registerRoutes() {
	//	h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/health").HandlerFunc(h.HandleHealth())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/openapi.yaml").
	// 	HandlerFunc(h.HandleOpenAPIYaml())
	//h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/login").HandlerFunc(h.HandleLogin())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/refresh-token").
	// 	HandlerFunc(h.HandleRefreshToken())

	h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/users").HandlerFunc(h.HandleCreateUser())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users").
	// 	HandlerFunc(h.HandleListUsers())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/me").
	// 	HandlerFunc(h.HandleGetCurrentUser())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/{userUuid}").
	// 	HandlerFunc(h.HandleGetUser())
	// h.router.Methods(gohttp.MethodPut, gohttp.MethodOptions).Path("/users/{userUuid}").
	// 	HandlerFunc(h.HandleUpdateUser())
	// h.router.Methods(gohttp.MethodPut, gohttp.MethodOptions).Path("/users/{userUuid}/ban").
	// 	HandlerFunc(h.HandleBanUser())
	// h.router.Methods(gohttp.MethodPut, gohttp.MethodOptions).Path("/users/{userUuid}/unban").
	// 	HandlerFunc(h.HandleUnbanUser())

	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/request-password-reset").
	// 	HandlerFunc(h.HandleRequestPasswordReset())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/reset-password").
	// 	HandlerFunc(h.HandleResetPassword())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/change-password").
	// 	HandlerFunc(h.HandleChangePassword())

	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/buildings").
	// 	HandlerFunc(h.HandleListBuildings())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/buildings").
	// 	HandlerFunc(h.HandleCreateBuilding())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/buildings/{buildingUuid}").
	// 	HandlerFunc(h.HandleGetBuilding())
	// h.router.Methods(gohttp.MethodPut, gohttp.MethodOptions).Path("/buildings/{buildingUuid}").
	// 	HandlerFunc(h.HandleUpdateBuilding())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/buildings/{buildingUuid}/bulk-bills").
	// 	HandlerFunc(h.HandleListBulkBillsByBuilding())

	// h.router.
	// 	Methods(gohttp.MethodGet, gohttp.MethodOptions).
	// 	Path("/bulk-bills").
	// 	HandlerFunc(h.HandleListBulkBills())
	// h.router.
	// 	Methods(gohttp.MethodPost, gohttp.MethodOptions).
	// 	Path("/bulk-bills").
	// 	HandlerFunc(h.HandleSubmitBulkBill())
	// h.router.
	// 	Methods(gohttp.MethodGet, gohttp.MethodOptions).
	// 	Path("/bulk-bills/{bulkBillUuid}").
	// 	HandlerFunc(h.HandleGetBulkBill())
	// h.router.
	// 	Methods(gohttp.MethodPut, gohttp.MethodOptions).
	// 	Path("/bulk-bills/{bulkBillUuid}").
	// 	HandlerFunc(h.HandleUpdateBulkBill())
	// h.router.
	// 	Methods(gohttp.MethodDelete, gohttp.MethodOptions).
	// 	Path("/bulk-bills/{bulkBillUuid}").
	// 	HandlerFunc(h.HandleDeleteBulkBill())
	// h.router.
	// 	Methods(gohttp.MethodPost, gohttp.MethodOptions).
	// 	Path("/bulk-bills/{bulkBillUuid}/generate-bills").
	// 	HandlerFunc(h.HandleGenerateBillsForBulkBill())
	// h.router.
	// 	Methods(gohttp.MethodGet, gohttp.MethodOptions).
	// 	Path("/bulk-bills/{bulkBillUuid}/bulk-pdf").
	// 	HandlerFunc(h.HandleGenerateBulkPDFURL())

	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/units").
	// 	HandlerFunc(h.HandleListUnits())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/me/units").
	// 	HandlerFunc(h.HandleListCurrentUserUnits())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/units").
	// 	HandlerFunc(h.HandleCreateUnit())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/units/{unitUuid}").
	// 	HandlerFunc(h.HandleGetUnit())
	// h.router.Methods(gohttp.MethodPut, gohttp.MethodOptions).Path("/units/{unitUuid}").
	// 	HandlerFunc(h.HandleUpdateUnit())
	// h.router.
	// 	Methods(gohttp.MethodGet, gohttp.MethodOptions).
	// 	Path("/units/{unitUuid}/monthly-consumption/{resource}").
	// 	HandlerFunc(h.HandleListUnitMonthlyConsumption())

	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/units/{unitUuid}/auto-payment").
	// 	HandlerFunc(h.HandleCreateAutoPaymentForUnit())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/units/{unitUuid}/auto-payment").
	// 	HandlerFunc(h.HandleGetAutoPaymentStatusForUnit())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/units/{unitUuid}/auto-payment/confirm").
	// 	HandlerFunc(h.HandleConfirmAutoPaymentForUnit())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/units/{unitUuid}/auto-payment/cancel").
	// 	HandlerFunc(h.HandleCancelAutoPaymentForUnit())

	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/power-records").
	// 	HandlerFunc(h.HandleSubmitPowerRecord())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/power-records").
	// 	HandlerFunc(h.HandleListPowerRecords())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/units/{unitUuid}/power-records").
	// 	HandlerFunc(h.HandleListPowerRecordsByUnit())

	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/bills").
	// 	HandlerFunc(h.HandleListBills())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/units/{unitUuid}/bills").
	// 	HandlerFunc(h.HandleListBillsByUnit())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/me/bills").
	// 	HandlerFunc(h.HandleListCurrentUserBills())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/me/units/{unitUuid}/bills").
	// 	HandlerFunc(h.HandleListCurrentUserUnitBills())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/bills/{billUuid}").
	// 	HandlerFunc(h.HandleGetBill())
	// h.router.
	// 	Methods(gohttp.MethodPost, gohttp.MethodOptions).
	// 	Path("/bills/{billUuid}/pdf").
	// 	HandlerFunc(h.HandleGenerateBillPDFURL())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/me/units/{unitUuid}/current-billing-rates").
	// 	HandlerFunc(h.HandleGetCurrentUserUnitCurrentBillingRates())

	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/transactions").
	// 	HandlerFunc(h.HandleListTransactions())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/users/me/transactions").
	// 	HandlerFunc(h.HandleListCurrentUserTransactions())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/transactions").
	// 	HandlerFunc(h.HandleCreateTransaction())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/transactions/{transactionUuid}/confirm").
	// 	HandlerFunc(h.HandleConfirmTransaction())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/transactions/{transactionUuid}/cancel").
	// 	HandlerFunc(h.HandleCancelTransaction())

	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/settings").
	// 	HandlerFunc(h.HandleListSettings())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/settings").
	// 	HandlerFunc(h.HandleCreateSetting())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/settings/{settingUuid}").
	// 	HandlerFunc(h.HandleGetSetting())
	// h.router.Methods(gohttp.MethodPut, gohttp.MethodOptions).Path("/settings/{settingUuid}").
	// 	HandlerFunc(h.HandleUpdateSetting())

	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/forms/{formUUID}").
	// 	HandlerFunc(h.HandleGetForm())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/forms").
	// 	HandlerFunc(h.HandlePostForm())
	// h.router.Methods(gohttp.MethodGet, gohttp.MethodOptions).Path("/forms").
	// 	HandlerFunc(h.HandleListForm())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/forms/{formUUID}/accept").
	// 	HandlerFunc(h.HandleAcceptForm())
	// h.router.Methods(gohttp.MethodPost, gohttp.MethodOptions).Path("/forms/{formUUID}/reject").
	// 	HandlerFunc(h.HandleRejectForm())

	// h.router.
	// 	Methods(gohttp.MethodGet, gohttp.MethodOptions).
	// 	Path("/apple-app-site-association").
	// 	HandlerFunc(h.HandleAppleAppSiteAssociation())
}
