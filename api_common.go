package openapi

import "github.com/antihax/optional"

type CompactOpts struct {
	Authoritative optional.Bool
}

type CompactionStatusOpts struct {
	Authoritative optional.Bool
}

type CreateNonPartitionedTopicOpts struct {
	Authoritative optional.Bool
}

type CreateSubscriptionOpts struct {
	Authoritative optional.Bool
	Replicated    optional.Bool
}

type DeletePartitionedTopicOpts struct {
	Force         optional.Bool
	Authoritative optional.Bool
}

type DeleteSubscriptionOpts struct {
	Authoritative optional.Bool
}

type DeleteTopicOpts struct {
	Force         optional.Bool
	Authoritative optional.Bool
}

type ExpireMessagesForAllSubscriptionsOpts struct {
	Authoritative optional.Bool
}

type ExpireTopicMessagesOpts struct {
	Authoritative optional.Bool
}

type GetBacklogOpts struct {
	Authoritative optional.Bool
}

type GetInternalStatsOpts struct {
	Authoritative optional.Bool
}

type GetLastMessageIdOpts struct {
	Authoritative optional.Bool
}

type GetPartitionedMetadataOpts struct {
	Authoritative optional.Bool
}

type GetPartitionedStatsOpts struct {
	Authoritative optional.Bool
}

type GetStatsOpts struct {
	Authoritative optional.Bool
}

type GetSubscriptionsOpts struct {
	Authoritative optional.Bool
}

type GrantPermissionsOnTopicOpts struct {
	Body optional.Interface
}

type OffloadStatusOpts struct {
	Authoritative optional.Bool
}

type PeekNthMessageOpts struct {
	Authoritative optional.Bool
}

type ResetCursorOpts struct {
	Authoritative optional.Bool
}

type ResetCursorOnPositionOpts struct {
	Authoritative optional.Bool
	MessageId     optional.Interface
}

type SkipAllMessagesOpts struct {
	Authoritative optional.Bool
}

type SkipMessagesOpts struct {
	Authoritative optional.Bool
}

type TerminateOpts struct {
	Authoritative optional.Bool
}

type TriggerOffloadOpts struct {
	Authoritative optional.Bool
}

type UnloadTopicOpts struct {
	Authoritative optional.Bool
}
