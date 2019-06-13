package packet

import (
	"bytes"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ItemFrameDropItem is sent by the server to make the item in an item frame drop out of the frame
// client-side.
type ItemFrameDropItem struct {
	// Position is the position of the item frame that had its item dropped. There must be a 'block entity'
	// present at this position.
	Position protocol.BlockPos
}

// ID ...
func (*ItemFrameDropItem) ID() uint32 {
	return IDItemFrameDropItem
}

// Marshal ...
func (pk *ItemFrameDropItem) Marshal(buf *bytes.Buffer) {
	_ = protocol.WriteUBlockPosition(buf, pk.Position)
}

// Unmarshal ...
func (pk *ItemFrameDropItem) Unmarshal(buf *bytes.Buffer) error {
	return protocol.UBlockPosition(buf, &pk.Position)
}
