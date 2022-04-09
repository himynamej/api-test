package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/himynamej/api-test.git/lib/http/middlewares"
	"github.com/himynamej/api-test.git/submetering"
)

type handler struct {
	router      *mux.Router
	accountsSvc submetering.AccountsService
	//buildingSvc    submetering.BuildingService
	//bulkBillSvc    submetering.BulkBillService
	//unitSvc        submetering.UnitService
	//powerRecordSvc submetering.PowerRecordService
	//billingSvc     submetering.BillingService
	//paymentSvc     submetering.PaymentService
	//settingSvc     submetering.SettingService
	//formSvc        submetering.FormService
	//openAPIYaml []byte
	loc *time.Location
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewHandler(
	//openAPIYaml []byte,
	accountsSvc submetering.AccountsService,
	// buildingSvc submetering.BuildingService,
	// unitSvc submetering.UnitService,
	// powerRecordSvc submetering.PowerRecordService,
	// billingSvc submetering.BillingService,
	// bulkBillSvc submetering.BulkBillService,
	// paymentSvc submetering.PaymentService,
	// settingSvc submetering.SettingService,
	// formSvc submetering.FormService,
	loc *time.Location,
) http.Handler {
	h := &handler{
		router:      mux.NewRouter(),
		accountsSvc: accountsSvc,
		// buildingSvc:    buildingSvc,
		// unitSvc:        unitSvc,
		// powerRecordSvc: powerRecordSvc,
		// billingSvc:     billingSvc,
		// bulkBillSvc:    bulkBillSvc,
		// paymentSvc:     paymentSvc,
		// settingSvc:     settingSvc,
		// formSvc:        formSvc,
		//openAPIYaml: openAPIYaml,
		loc: loc,
	}

	h.router.StrictSlash(true)
	h.router.Use(middlewares.CORS())
	h.router.Use(middlewares.Logger(log.New(os.Stdout, fmt.Sprintln(), 0)))
	h.router.Use(middlewares.RecoverPanic())

	h.registerRoutes()

	return h
}
