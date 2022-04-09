package http

import (
	"fmt"
	"net/http"

	"github.com/himynamej/api-test.git/lib/problem"
	"github.com/shopspring/decimal"
)

func InvalidEmailOrPasswordError() problem.Problem {
	return problem.CustomError(
		problem.WithTitle("Invalid email address or password provided."),
		problem.WithDetail("One of email address or password is invalid."),
		problem.WithStatus(http.StatusUnauthorized),
	)
}

func UserIsBannedError(detail string) problem.Problem {
	return problem.CustomError(
		problem.WithTitle("User is banned."),
		problem.WithDetail(detail),
		problem.WithStatus(http.StatusUnauthorized),
	)
}

func InvalidChangePasswordTokenError() problem.Problem {
	return problem.CustomError(
		problem.WithTitle("Invalid token provided."),
		problem.WithDetail("The token you provided is invalid or expired."),
		problem.WithStatus(http.StatusUnauthorized),
	)
}

func InsecurePasswordError() problem.Problem {
	return problem.CustomError(
		problem.WithTitle("Insecure password."),
		problem.WithDetail("The password is not strong enough. Try including more special characters, using lowercase letters, using uppercase letters or using a longer password"),
		problem.WithStatus(http.StatusBadRequest),
	)
}

func InvalidOldPasswordError() problem.Problem {
	return problem.CustomError(
		problem.WithTitle("Invalid old password provided."),
		problem.WithDetail("The old password you provided is invalid."),
		problem.WithStatus(http.StatusBadRequest),
	)
}

func NoPasswordSetError() problem.Problem {
	return problem.CustomError(
		problem.WithTitle("No old password has been set."),
		problem.WithDetail("This profile doesn't have any password, so you can't change it. Use forget password method."),
		problem.WithStatus(http.StatusForbidden),
	)
}

func InvalidCardError() problem.Problem {
	return problem.CustomError(
		problem.WithTitle("Invalid card information provided."),
		problem.WithDetail("Invalid card number, expire date, or CVC entered."),
		problem.WithStatus(http.StatusBadRequest),
	)
}

func LowPaymentAmountError(min decimal.Decimal) problem.Problem {
	return problem.CustomError(
		problem.WithTitle("The amount is less than the minimum!"),
		problem.WithDetail(fmt.Sprintf("The minimum payment amount is $%s! Please contact support team for further information.", min.String())),
		problem.WithStatus(http.StatusBadRequest),
	)
}
