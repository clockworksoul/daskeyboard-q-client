package qclient

type Signal struct {
	Name       string `json:"name,omitempty"`
	Message    string `json:"message,omitempty"`
	ZoneId     string `json:"zoneId,omitempty"`
	Color      string `json:"color,omitempty"`
	Effect     string `json:"effect,omitempty"`
	Pid        string `json:"pid,omitempty"`
	IsArchived bool   `json:"isArchived,omitempty"`
	IsRead     bool   `json:"isRead,omitempty"`
	IsMuted    bool   `json:"isMuted,omitempty"`
	ClientName string `json:"clientName,omitempty"`
}

type SignalResponse struct {
	UserId    uint   `json:"name,omitempty"`
	CreatedAt uint64 `json:"name,omitempty"`
	UpdatedAt uint64 `json:"name,omitempty"`
}
