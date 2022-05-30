package routeguides

import (
	"context"
	"user-service/Models"
	notifications "user-service/Rpc/Client/Notifications"
)

func SendUserVerificationCode(user Models.User, verificationCode string) (err error) {
	var userRpc notifications.User

	userRpc.FcmToken = user.FcmToken
	userRpc.Name = user.Name
	userRpc.WhatsappNo = user.WhatsappNo
	userRpc.Email = user.Email
	userRpc.VerificationCode = verificationCode

	_, err = notifications.DialNotificationsAPI(context.Background()).NotificationServiceClient.SendUserVerificationCode(context.Background(), &userRpc)

	if err != nil {
		return err
	}

	return nil
}
