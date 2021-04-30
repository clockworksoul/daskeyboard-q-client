package qclient

import "fmt"

type ZoneID string

type Key string

const (
	KeyEsc Key = "KEY_ESC"
	KeyF1  Key = "KEY_F1"
	KeyF2  Key = "KEY_F2"
	KeyF3  Key = "KEY_F3"
	KeyF4  Key = "KEY_F4"
	KeyF5  Key = "KEY_F5"
	KeyF6  Key = "KEY_F6"
	KeyF7  Key = "KEY_F7"
	KeyF8  Key = "KEY_F8"
	KeyF9  Key = "KEY_F9"
	KeyF10 Key = "KEY_F10"
	KeyF11 Key = "KEY_F11"
	KeyF12 Key = "KEY_F12"
	Key0   Key = "KEY_0"
	Key1   Key = "KEY_1"
	Key2   Key = "KEY_2"
	Key3   Key = "KEY_3"
	Key4   Key = "KEY_4"
	Key5   Key = "KEY_5"
	Key6   Key = "KEY_6"
	Key7   Key = "KEY_7"
	Key8   Key = "KEY_8"
	Key9   Key = "KEY_9"
	KeyA   Key = "KEY_A"
	KeyB   Key = "KEY_B"
	KeyC   Key = "KEY_C"
	KeyD   Key = "KEY_D"
	KeyE   Key = "KEY_E"
	KeyF   Key = "KEY_F"
	KeyG   Key = "KEY_G"
	KeyH   Key = "KEY_H"
	KeyI   Key = "KEY_I"
	KeyJ   Key = "KEY_J"
	KeyK   Key = "KEY_K"
	KeyL   Key = "KEY_L"
	KeyM   Key = "KEY_M"
	KeyN   Key = "KEY_N"
	KeyO   Key = "KEY_O"
	KeyP   Key = "KEY_P"
	KeyQ   Key = "KEY_Q"
	KeyR   Key = "KEY_R"
	KeyS   Key = "KEY_S"
	KeyT   Key = "KEY_T"
	KeyU   Key = "KEY_U"
	KeyV   Key = "KEY_V"
	KeyW   Key = "KEY_W"
	KeyX   Key = "KEY_X"
	KeyY   Key = "KEY_Y"
	KeyZ   Key = "KEY_Z"
)

func ZoneIDKey(k Key) ZoneID {
	return ZoneID(k)
}

func Zone2D(x, y int) ZoneID {
	return ZoneID(fmt.Sprintf("%d,%d", x, y))
}

func ZoneLinear(n int) ZoneID {
	return ZoneID(fmt.Sprintf("%d", n))
}

type Effect string

const (
	EffectSetColor   Effect = "SET_COLOR"
	EffectBlink      Effect = "BLINK"
	EffectBreathe    Effect = "BREATHE"
	EffectColorCycle Effect = "COLOR_CYCLE"
)

type SignalRequest struct {
	Name       string `json:"name,omitempty"`
	Message    string `json:"message,omitempty"`
	ZoneID     Key    `json:"zoneId,omitempty"`
	Color      string `json:"color,omitempty"`
	Effect     Effect `json:"effect,omitempty"`
	ProductID  string `json:"pid,omitempty"`
	IsArchived bool   `json:"isArchived,omitempty"`
	IsRead     bool   `json:"isRead,omitempty"`
	IsMuted    bool   `json:"isMuted,omitempty"`
	ClientName string `json:"clientName,omitempty"`
}

const (
	DefaultProductID  = "Q_MATRIX"
	DefaultClientName = "GoClient"
)

func NewSignalRequest(name, color string, zoneID Key) *SignalRequest {
	return &SignalRequest{
		Name:       name,
		ZoneID:     zoneID,
		Color:      color,
		ProductID:  DefaultProductID,
		ClientName: DefaultClientName,
	}
}

func (r *SignalRequest) WithName(s string) *SignalRequest {
	r.Name = s
	return r
}

func (r *SignalRequest) WithMessage(s string) *SignalRequest {
	r.Message = s
	return r
}

func (r *SignalRequest) WithZoneID(k Key) *SignalRequest {
	r.ZoneID = k
	return r
}

func (r *SignalRequest) WithColor(s string) *SignalRequest {
	r.Color = s
	return r
}

func (r *SignalRequest) WithEffect(e Effect) *SignalRequest {
	r.Effect = e
	return r
}

func (r *SignalRequest) WithProductID(s string) *SignalRequest {
	r.ProductID = s
	return r
}

func (r *SignalRequest) WithArchived(b bool) *SignalRequest {
	r.IsArchived = b
	return r
}

func (r *SignalRequest) WithRead(b bool) *SignalRequest {
	r.IsRead = b
	return r
}

func (r *SignalRequest) WithMuted(b bool) *SignalRequest {
	r.IsMuted = b
	return r
}

type SignalResponse struct {
	SignalRequest
	SignalId  int    `json:"id,omitempty"`
	UserId    int    `json:"userId,omitempty"`
	CreatedAt uint64 `json:"createdAt,omitempty"`
	UpdatedAt uint64 `json:"updatedAt,omitempty"`
}

type SignalResponsePage struct {
	Content       []*SignalResponse `json:"content,omitempty"`
	Size          int               `json:"size,omitempty"`
	Sort          string            `json:"sort,omitempty"`
	HasNextPage   bool              `json:"hasNextPage,omitempty"`
	Page          int               `json:"page,omitempty"`
	TotalElements int               `json:"totalElements,omitempty"`
	TotalPages    int               `json:"totalPages,omitempty"`
}
