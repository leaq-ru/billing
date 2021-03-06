package billingimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/leaq-ru/proto/codegen/go/billing"
	"net/http"
	"strconv"
)

func (s *server) RobokassaWebhook(
	_ context.Context,
	req *billing.RobokassaWebhookRequest,
) (
	res *wrappers.StringValue,
	err error,
) {
	invID, err := strconv.Atoi(req.GetInvId())
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = errors.New(http.StatusText(http.StatusBadRequest))
		return
	}

	err = s.robokassaWebhook.ProcessAsync(req.GetSecret(), uint64(invID), req.GetOutSum(), req.GetSignatureValue())
	if err != nil {
		s.logger.Error().Err(err).Send()
		err = errors.New(http.StatusText(http.StatusInternalServerError))
		return
	}

	res = &wrappers.StringValue{
		Value: "OK" + strconv.Itoa(invID),
	}
	return
}
