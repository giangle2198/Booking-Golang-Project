package util

import (
	"booking-center-service/internal/common"
	bimClientPB "booking-identity-management/pb"
	"context"
	"regexp"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	pattern = "^Bearer (\\S*)"
)

func ContainInArray(array []string, item string) bool {
	for _, value := range array {
		if strings.EqualFold(value, item) {
			return true
		}
	}
	return false
}

func checkJWT(ctx context.Context,
	bimClient bimClientPB.BIMAPIClient,
	jwt string) (domain, statusCode, errorCode string, err error) {
	verifyTokenReq := &bimClientPB.VerifyTokenRequest{
		Jwt: jwt,
	}
	verifyTokenResp, err := bimClient.DoVerifyJWTToken(ctx, verifyTokenReq)
	if err != nil {
		return domain, statusCode, errorCode, err
	}
	domain = verifyTokenResp.Domain
	errorCode = verifyTokenResp.ReasonCode
	statusCode = verifyTokenResp.StatusCode
	return domain, statusCode, errorCode, nil
}

func JWTAuthentication(bimClient bimClientPB.BIMAPIClient, excludePaths []string) grpc.UnaryServerInterceptor {
	pattern, _ := regexp.Compile(pattern)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}
		methods := strings.Split(info.FullMethod, "/")
		if ContainInArray(excludePaths, methods[len(methods)-1:][0]) {
			return handler(ctx, req)
		}

		if authorizations, ok := md["authorization"]; ok && len(authorizations) != 0 {
			firstItem := authorizations[0]
			if firstItem != "" && pattern.Match([]byte(firstItem)) {
				tokens := pattern.FindSubmatch([]byte(firstItem))
				if len(tokens) == 2 && string(tokens[1]) != "" {
					jwt := string(tokens[1])
					domain, statusCode, _, err := checkJWT(ctx, bimClient, jwt)
					if err != nil {
						return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
					}
					if domain != "" {
						md["domain"] = []string{domain}
						md["token"] = []string{firstItem}
					}
					switch statusCode {
					case common.FAILEDStatus:
						return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
					case common.DONEStatus:
						ctx = metadata.NewIncomingContext(ctx, md)
						return handler(ctx, req)
					}
					return handler(ctx, req)
				}
			}
		}
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}
}
