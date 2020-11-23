package call

import (
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"google.golang.org/grpc"
)

func NewClients(
	parserURL,
	userURL string,
) (
	companyClient parser.CompanyClient,
	userClient user.UserClient,
	err error,
) {
	connParser, err := grpc.Dial(parserURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	companyClient = parser.NewCompanyClient(connParser)

	connUser, err := grpc.Dial(userURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	userClient = user.NewUserClient(connUser)

	return
}
