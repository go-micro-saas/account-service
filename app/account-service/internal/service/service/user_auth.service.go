package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/service/dto"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// userAuth ...
type userAuth struct {
	servicev1.UnimplementedSrvUserAuthV1Server

	log                  *log.Helper
	userAuthBizRepo      bizrepos.UserAuthBizRepo
	sendEmailCodeBizRepo bizrepos.SendEmailCodeBizRepo
}

// NewUserAuthService ...
func NewUserAuthService(
	logger log.Logger,
	userAuthBizRepo bizrepos.UserAuthBizRepo,
	sendEmailCodeBizRepo bizrepos.SendEmailCodeBizRepo,
) servicev1.SrvUserAuthV1Server {
	return &userAuth{
		log:                  log.NewHelper(log.With(logger, "module", "account-service/service/user_auth")),
		userAuthBizRepo:      userAuthBizRepo,
		sendEmailCodeBizRepo: sendEmailCodeBizRepo,
	}
}

// Ping ping pong
func (s *userAuth) Ping(ctx context.Context, in *resourcev1.PingReq) (out *resourcev1.PingResp, err error) {
	// 可以解析
	authClaims, tokenClaimsOK := authpkg.GetAuthClaimsFromContext(ctx)
	s.log.WithContext(ctx).Infow(
		"tokenClaimsOK", tokenClaimsOK,
		"authClaims", authClaims,
	)

	out = &resourcev1.PingResp{
		Message: in.Message,
	}
	return out, err
}

// LoginByEmail Email登录
func (s *userAuth) LoginByEmail(ctx context.Context, in *resourcev1.LoginByEmailReq) (*resourcev1.LoginResp, error) {
	userModel, signResp, err := s.userAuthBizRepo.LoginByEmail(ctx, in)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

// LoginByPhone 手机登录
func (s *userAuth) LoginByPhone(ctx context.Context, in *resourcev1.LoginByPhoneReq) (*resourcev1.LoginResp, error) {
	userModel, signResp, err := s.userAuthBizRepo.LoginByPhone(ctx, in)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

// RefreshToken 刷新token
func (s *userAuth) RefreshToken(ctx context.Context, in *resourcev1.RefreshTokenReq) (*resourcev1.LoginResp, error) {
	signResp, err := s.userAuthBizRepo.RefreshToken(ctx, in.RefreshToken)
	if err != nil {
		return nil, err
	}

	// user
	userId := signResp.TokenResponse.RefreshTokenItem.Payload.UserID
	userModel, err := s.userAuthBizRepo.CheckAndGetUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) SendPhoneSignupCode(ctx context.Context, req *resourcev1.SendPhoneSignupCodeReq) (*resourcev1.SendSignupCodeResp, error) {
	return nil, errorpkg.WithStack(errorpkg.DefaultErrorMethodNotAllowed())
	//param := dto.AccountDto.ToBoSendVerifyCodeParam(req)
	//dataModel, err := s.userAuthBizRepo.SendVerifyCode(ctx, param)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &resourcev1.SendSignupCodeResp{
	//	Data: dto.AccountDto.ToPbSendSignupCodeRespData(dataModel),
	//}, nil
}

func (s *userAuth) SendEmailSignupCode(ctx context.Context, req *resourcev1.SendEmailSignupCodeReq) (*resourcev1.SendSignupCodeResp, error) {
	param := dto.AccountDto.ToBoSendVerifyCodeParam2(req)
	dataModel, err := s.userAuthBizRepo.SendVerifyCode(ctx, param)
	if err != nil {
		return nil, err
	}

	return &resourcev1.SendSignupCodeResp{
		Data: dto.AccountDto.ToPbSendSignupCodeRespData(dataModel),
	}, nil
}

// SignupByEmail 身份验证-Email注册
func (s *userAuth) SignupByEmail(ctx context.Context, req *resourcev1.SignupByEmailReq) (*resourcev1.LoginResp, error) {
	userModel, signResp, err := s.userAuthBizRepo.SignupByEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

// SignupByPhone 身份验证-手机注册
func (s *userAuth) SignupByPhone(ctx context.Context, req *resourcev1.SignupByPhoneReq) (*resourcev1.LoginResp, error) {
	userModel, signResp, err := s.userAuthBizRepo.SignupByPhone(ctx, req)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) LoginOrSignupByPhone(ctx context.Context, req *resourcev1.LoginOrSignupByPhoneReq) (*resourcev1.LoginResp, error) {
	return s.UnimplementedSrvUserAuthV1Server.LoginOrSignupByPhone(ctx, req)
}

func (s *userAuth) LoginOrSignupByEmail(ctx context.Context, req *resourcev1.LoginOrSignupByEmailReq) (*resourcev1.LoginResp, error) {
	return s.UnimplementedSrvUserAuthV1Server.LoginOrSignupByEmail(ctx, req)
}
