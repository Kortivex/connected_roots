package ferrors

import (
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"github.com/Kortivex/connected_roots/pkg/utils"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"

	"github.com/labstack/echo/v4"
)

func MatchError(err error) *commons.ErrorS {
	switch e := utils.UnwrapErr(err).(type) {
	default:
		if value, ok := errorAPIMap[e.Error()]; ok {
			return commons.NewErrorS(value.Status, value.Message, value.Details, err).(*commons.ErrorS)
		}

		return commons.NewErrorS(http.StatusInternalServerError, ErrSomethingWentWrong.Error(), nil, err).(*commons.ErrorS)
	}
}

func Error401(c echo.Context) error {
	return c.Render(http.StatusOK, "401.gohtml", translator.AddDataKeys(CommonErrorsPageI18N(c), map[string]interface{}{
		"title_401":  translator.T(c, translator.PagesErrorsTitle401),
		"body_401":   translator.T(c, translator.PagesErrorsBody401),
		"button_401": translator.T(c, translator.PagesErrorsButton401),
	}))
}

func Error403(c echo.Context) error {
	return c.Render(http.StatusOK, "403.gohtml", translator.AddDataKeys(CommonErrorsPageI18N(c), map[string]interface{}{
		"title_403":  translator.T(c, translator.PagesErrorsTitle403),
		"body_403":   translator.T(c, translator.PagesErrorsBody403),
		"button_403": translator.T(c, translator.PagesErrorsButton403),
	}))
}

func Error404(c echo.Context) error {
	return c.Render(http.StatusOK, "404.gohtml", translator.AddDataKeys(CommonErrorsPageI18N(c), map[string]interface{}{
		"title_404":  translator.T(c, translator.PagesErrorsTitle404),
		"body_404":   translator.T(c, translator.PagesErrorsBody404),
		"button_404": translator.T(c, translator.PagesErrorsButton404),
	}))
}

func Error500(c echo.Context) error {
	return c.Render(http.StatusOK, "500.gohtml", translator.AddDataKeys(CommonErrorsPageI18N(c), map[string]interface{}{
		"title_500":  translator.T(c, translator.PagesErrorsTitle500),
		"body_500":   translator.T(c, translator.PagesErrorsBody500),
		"button_500": translator.T(c, translator.PagesErrorsButton500),
	}))
}
