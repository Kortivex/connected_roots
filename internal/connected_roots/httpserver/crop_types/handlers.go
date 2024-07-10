package crop_types

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots/crop_types"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingCropTypesHandlers = "http-handler.crop-types"

	tracingPostCropTypesHandlers   = "http-handler.crop-types.post-crop-type"
	tracingPutCropTypesHandlers    = "http-handler.crop-types.put-crop-type"
	tracingGetCropTypesHandlers    = "http-handler.crop-types.get-crop-type"
	tracingListCropTypesHandlers   = "http-handler.crop-types.list-crop-types"
	tracingDeleteCropTypesHandlers = "http-handler.crop-types.delete-crop-type"

	cropTypeIDParam = "crop_type_id"
)

type CropTypesHandlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	// Services.
	cropTypeSvc *crop_types.Service
}

func NewCropTypesHandlers(appCtx *connected_roots.Context) *CropTypesHandlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingCropTypesHandlers)

	return &CropTypesHandlers{
		gorm:        appCtx.Gorm,
		logger:      log,
		conf:        appCtx.Conf,
		cropTypeSvc: crop_types.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *CropTypesHandlers) ListCropTypesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListCropTypesHandlers)
	defer span.End()

	filters := connected_roots.CropTypePaginationFilters{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListCropTypesHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	rolesRes, err := h.cropTypeSvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListCropTypesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, rolesRes)
}
