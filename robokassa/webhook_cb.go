package robokassa

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/nats-io/stan.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"strconv"
	"strings"
	"time"
)

func (w Webhook) cb(stanMsg *stan.Msg) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ack := func() {
			e := stanMsg.Ack()
			if e != nil {
				w.logger.Error().Err(e).Send()
			}
		}

		var msg webhookMessage
		err := json.Unmarshal(stanMsg.Data, &msg)
		if err != nil {
			w.logger.Error().Err(err).Msg("json.Unmarshal error. Seems got invalid msg, calling Ack")
			ack()
			return
		}

		w.logger.Info().Str("stanMsg.Data", string(stanMsg.Data)).Msg("got Robokassa webhook message")

		if stanMsg.RedeliveryCount >= 5 {
			w.logger.Error().
				Uint32("redeliveryCount", stanMsg.RedeliveryCount).
				Str("data", string(stanMsg.Data)).
				Msg("seems got dead letter message")

			if stanMsg.RedeliveryCount >= 1000 {
				ack()
				return
			}
		}

		var eg errgroup.Group
		var alreadyProcessed bool
		eg.Go(func() (e error) {
			alreadyProcessed, e = w.eventLogModel.AlreadyProcessed(ctx, msg.ID)
			return
		})

		var hasSuccessStatus bool
		eg.Go(func() (e error) {
			hasSuccessStatus, e = w.invoiceModel.IsDebitSuccessRK(ctx, msg.InvID)
			return
		})

		var userOID primitive.ObjectID
		eg.Go(func() (e error) {
			userOID, e = w.invoiceModel.GetUserIDByPendingRKInvoiceID(ctx, msg.InvID)
			return
		})
		err = eg.Wait()
		if err != nil {
			w.logger.Error().Err(err).Send()
			return
		}

		if alreadyProcessed || hasSuccessStatus {
			ack()
			return
		}

		invID := strconv.Itoa(int(msg.InvID))

		// robokassa has different behavior on test and prod
		var outSum string
		if isTest(w.isTest) {
			outSum = strconv.FormatFloat(float64(msg.OutSum), 'f', -1, 64)
		} else {
			outSum = strconv.Itoa(int(msg.OutSum)) + ".000000"
		}

		sha := sha512.New()
		_, err = sha.Write([]byte(strings.Join([]string{
			outSum,
			invID,
			w.passwordTwo,
		}, ":")))
		if err != nil {
			w.logger.Error().Err(err).Send()
			return
		}

		expectedHash := strings.ToUpper(hex.EncodeToString(sha.Sum(nil)))
		if expectedHash != msg.SignatureValue {
			err = errors.New("invalid SignatureValue. Seems got invalid msg, not calling Ack")
			w.logger.Error().
				Str("expectedHash", expectedHash).
				Str("msg.SignatureValue", msg.SignatureValue).
				Uint64("InvID", msg.InvID).Err(err).Send()
			return
		}

		sess, err := w.mongoStartSession()
		if err != nil {
			w.logger.Error().Err(err).Send()
			return
		}
		defer sess.EndSession(ctx)

		_, err = sess.WithTransaction(ctx, func(sc mongo.SessionContext) (_ interface{}, e error) {
			e = w.eventLogModel.Put(sc, msg.ID)
			if e != nil {
				w.logger.Error().Err(e).Send()
				return
			}

			pennyAmount := uint32(msg.OutSum * 100)

			e = w.invoiceModel.CreateSuccessDebitRK(sc, userOID, msg.InvID, pennyAmount)
			if e != nil {
				w.logger.Error().Err(e).Send()
				return
			}

			e = w.balanceModel.Inc(sc, userOID, pennyAmount)
			if e != nil {
				w.logger.Error().Err(e).Send()
			}
			return
		})
		if err != nil {
			w.logger.Error().Err(err).Send()
			return
		}

		ack()
		return
	}()
}

func isTest(isTest string) bool {
	return isTest == "1"
}
