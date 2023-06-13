package lib

const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteConversationsGetUnreadConversationsCount         = "conversations/unread-status"
	RouteConversationsGetUserConversations                = "conversations"
	RouteConversationsAddPrivateOrGroupConversation       = "conversations"
	RouteConversationsGetConversationForGivenUsernames    = "conversations/by-usernames?userNames=%s"
	RouteConversationsSendMessageToAnExistingConversation = "conversations/%s/messages"
	RouteConversationsGetConversation                     = "conversations/%s"
	RouteConversationsDeleteConversation                  = "conversations/%s"
	RouteConversationsLeaveConversation                   = "conversations/%s/leave"
	RouteConversationsArchiveConversation                 = "conversations/%s/archive"
)

const (
	RouteProfileGetUsersProfile               = "profiles"
	RouteProfileUpdateUsersProfile            = "profiles"
	RouteProfileUpdateUsername                = "profiles/username/%s"
	RouteProfileBlocksUserFromSendingMessages = "users/%s/block"
	RouteProfileUnblocksUser                  = "users/%s/unblock"
)

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = "app/users/reset-password"
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

const (
	RouteUserGetUsers              = "users"
	RouteUserGetUserByUuid         = "users/uuid/%s"
	RouteUserGetUserByEmail        = "users/email/%s"
	RouteUserDeactivateUserAccount = "users/deactivate-account"
)

const (
	RouteVersionGetLatestVersion = "app/latest-version"
	RouteVersionGetAppStatus     = "app/status"
)
