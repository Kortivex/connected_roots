package orchard

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/bars"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	sessionServ "github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"github.com/Kortivex/connected_roots/pkg/sdk"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
)

const (
	tracingOrchardHandlers = "http-handler.orchard"

	getCreateOrchardHandlerName  = "http-handler.orchard.get-create-orchard"
	postCreateOrchardHandlerName = "http-handler.orchard.post-create-orchard"
	getUpdateOrchardHandlerName  = "http-handler.orchard.get-update-orchard"
	postUpdateOrchardHandlerName = "http-handler.orchard.post-update-orchard"
	getViewOrchardHandlerName    = "http-handler.orchard.get-view-orchard"
	getListOrchardHandlerName    = "http-handler.orchard.get-list-orchards"
	getDeleteOrchardHandlerName  = "http-handler.orchard.get-delete-orchard"
	postDeleteOrchardHandlerName = "http-handler.orchard.post-delete-orchard"

	orchardIDParam = "orchard_id"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewOrchardsHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingOrchardHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetOrchardsListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getListOrchardHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getListOrchardHandlerName)

	message, err := h.sessionSvc.ObtainMessage(ctx, c, "message")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	log.Debug(fmt.Sprintf("message: %s", message))

	notifications := map[string]interface{}{}
	if len(message) > 0 && message[0] == "success" {
		notifications = map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminOrchardsDeleteSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminOrchardsDeleteSuccessMessage),
		}
	}

	nextCursor := ""
	if c.QueryParam("next_cursor") != "" {
		nextCursor = c.QueryParam("next_cursor")
		log.Debug(fmt.Sprintf("next_cursor: %s", nextCursor))
	}
	prevCursor := ""
	if c.QueryParam("previous_cursor") != "" {
		prevCursor = c.QueryParam("previous_cursor")
		log.Debug(fmt.Sprintf("previous_cursor: %s", prevCursor))
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	orchards, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "20", nextCursor, prevCursor, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-orchards-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonOrchardListPageI18N(c)), map[string]interface{}{
			"orchards":   orchards,
			"pagination": pagination,
		}), notifications))
}