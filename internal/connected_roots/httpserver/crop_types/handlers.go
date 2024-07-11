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

func (h *CropTypesHandlers) PostCropTypeHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostCropTypesHandlers)
	defer span.End()

	cropTypeBody := connected_roots.CropTypes{}
	if err := c.Bind(&cropTypeBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostCropTypesHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	cropTypeRes, err := h.cropTypeSvc.Save(ctx, &cropTypeBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPostCropTypesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, cropTypeRes)
}

func (h *CropTypesHandlers) PutCropTypeHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPutCropTypesHandlers)
	defer span.End()

	cropTypeID := c.Param(cropTypeIDParam)
	if cropTypeID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	cropTypeBody := connected_roots.CropTypes{}
	if err := c.Bind(&cropTypeBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPutCropTypesHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	cropTypeBody.ID = cropTypeID

	cropTypeRes, err := h.cropTypeSvc.Update(ctx, &cropTypeBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPutCropTypesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, cropTypeRes)
}

func (h *CropTypesHandlers) GetCropTypeHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetCropTypesHandlers)
	defer span.End()

	cropTypeID := c.Param(cropTypeIDParam)
	if cropTypeID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	cropTypeRes, err := h.cropTypeSvc.Obtain(ctx, cropTypeID)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingGetCropTypesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, cropTypeRes)
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

func (h *CropTypesHandlers) DeleteCropTypeHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingDeleteCropTypesHandlers)
	defer span.End()

	cropTypeID := c.Param(cropTypeIDParam)
	if cropTypeID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	if err := h.cropTypeSvc.Remove(ctx, cropTypeID); err != nil {
		err = fmt.Errorf("%s: %w", tracingDeleteCropTypesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
