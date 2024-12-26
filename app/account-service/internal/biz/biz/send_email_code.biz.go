package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	errorv1 "github.com/go-micro-saas/account-service/api/account-service/v1/errors"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	emailpkg "github.com/ikaiguang/go-srv-kit/kit/email"
	regexpkg "github.com/ikaiguang/go-srv-kit/kit/regex"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"strings"
)

type sendEmailCodeBiz struct {
	log            *log.Helper
	emailClient    emailpkg.Client
	messageExample emailpkg.Message
}

func NewSendEmailCodeBiz(
	logger log.Logger,
	cfg *bo.SendEmailCodeConfig,
) (bizrepos.SendEmailCodeBizRepo, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "account-service/biz/send_email_code"))
	emailClient, err := emailpkg.NewClient(*cfg.Sender)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	cleanup := func() {
		cErr := emailClient.Close()
		if cErr != nil {
			logHelper.Errorw("msg", "close email client error", "error", cErr)
		}
	}
	return &sendEmailCodeBiz{
		log:            logHelper,
		emailClient:    emailClient,
		messageExample: cfg.Message,
	}, cleanup, nil
}

func (s *sendEmailCodeBiz) SendEmailCode(ctx context.Context, param *bo.SendEmailCodeParam) (*bo.SendEmailCodeReply, error) {
	param.VerifyCode = strings.TrimSpace(param.VerifyCode)
	if param.VerifyCode == "" {
		e := errorpkg.ErrorInvalidParameter("verify_code")
		return nil, errorpkg.WithStack(e)
	}
	if !regexpkg.IsEmail(param.VerifyAccount) {
		e := errorv1.DefaultErrorS103InvalidEmail()
		return nil, errorpkg.WithStack(e)
	}

	// send
	msg := s.messageExample
	msg.To = []string{param.VerifyAccount}
	message := &emailpkg.CodeMessage{
		Message: &msg,
		Code:    param.VerifyCode,
	}
	err := s.emailClient.SendCode(message)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return &bo.SendEmailCodeReply{
		IsSendToServer: true,
		Code:           param.VerifyCode,
	}, nil
}