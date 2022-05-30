package notifications

import (
	"context"
	"fmt"
	"log"
	"sync"

	grpc "google.golang.org/grpc"
)

type dialNotificationsAPI struct {
	NotificationServiceClient NotificationServiceClient
}

var initDefaulNotificationsAPI *dialNotificationsAPI = nil
var onceDefaulNotificationsAPI sync.Once

func DialNotificationsAPI(context context.Context) *dialNotificationsAPI {
	onceDefaulNotificationsAPI.Do(func() {
		conn, err := grpc.Dial("localhost:6005", grpc.WithInsecure())

		if err != nil {
			log.Println(fmt.Printf("did not connect: %s", err))
		}

		initDefaulNotificationsAPI = &dialNotificationsAPI{
			NotificationServiceClient: NewNotificationServiceClient(conn),
		}
	})

	return initDefaulNotificationsAPI
}
