package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/geeeeorge/Go-book-review/pkg/cognito"
)

// CognitoMiddleware Cognitoを利用した認証
// ignorePathsは認証が必要のないAPIのPathの配列
func CognitoMiddleware(region, poolID, iss string, ignorePaths []string) func(next echo.HandlerFunc) echo.HandlerFunc {
	pathMap := map[string]struct{}{}
	for _, path := range ignorePaths {
		pathMap[path] = struct{}{}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			// 認証が不要なapiの場合はそのまま返す
			if _, ok := pathMap[ec.Path()]; ok {
				return next(ec)
			}

			authorization, ok := ec.Request().Header[http.CanonicalHeaderKey("Authorization")]
			if !ok {
				return ec.JSON(http.StatusBadRequest, map[string]string{
					"msg": "no authorization header",
				})
			}
			if len(authorization) != 1 {
				return ec.JSON(http.StatusBadRequest, map[string]string{
					"msg": "only one authorization header is required",
				})
			}

			tok := strings.TrimPrefix(authorization[0], "Bearer ")
			jwt, err := cognito.NewValid(region, poolID, iss, tok, time.Now())
			if err != nil {
				return ec.JSON(http.StatusUnauthorized, map[string]string{
					"msg":          "unauthorized",
					"error_detail": err.Error(),
				})
			}

			ctx := ec.Request().Context()
			ctx = context.WithValue(ctx, cognito.ContextAuthorizationKey, jwt)
			ec.SetRequest(ec.Request().WithContext(ctx))

			return next(ec)
		}
	}
}
