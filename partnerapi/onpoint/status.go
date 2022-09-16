package onpoint

const (
	OrderStatusNew              = "new"
	OrderStatusPendingWarehouse = "pending_warehouse"
	OrderStatusWhProcessing     = "wh_processing"
	OrderStatusWhCompleted      = "wh_completed"
	OrderStatusDlPending        = "dl_pending"
	OrderStatusDlIntransit      = "dl_intransit"
	OrderStatusDLDelivered      = "dl_delivered"
	OrderStatusDLReturning      = "dl_returning"
	OrderStatusReturned         = "returned"
	OrderStatusPartialCancelled = "partial_cancelled"
	OrderStatusCancelled        = "cancelled"
	OrderStatusCompleted        = "completed"
	OrderStatusUnknown          = "unknown"
)
