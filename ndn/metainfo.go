/* YaNFD - Yet another NDN Forwarding Daemon
 *
 * Copyright (C) 2020 Eric Newberry.
 *
 * This file is licensed under the terms of the MIT License, as found in LICENSE.md.
 */

package ndn

import (
	"errors"
	"strconv"
	"time"

	"github.com/eric135/YaNFD/ndn/tlv"
	"github.com/eric135/YaNFD/ndn/util"
)

// MetaInfo represents the MetaInfo in a Data packet.
type MetaInfo struct {
	contentType     *uint64
	freshnessPeriod *time.Duration
	finalBlockID    NameComponent
	wire            *tlv.Block
}

// NewMetaInfo creates a new MetaInfo structure.
func NewMetaInfo() *MetaInfo {
	m := new(MetaInfo)
	return m
}

// DecodeMetaInfo decodes a MetaInfo from a block.
func DecodeMetaInfo(wire *tlv.Block) (*MetaInfo, error) {
	if wire == nil {
		return nil, util.ErrNonExistent
	}
	wire.Parse()

	m := new(MetaInfo)
	m.wire = wire.DeepCopy()
	mostRecentElem := 0
	for _, elem := range wire.Subelements() {
		switch elem.Type() {
		case tlv.ContentType:
			if mostRecentElem >= 1 {
				return nil, errors.New("ContentType is duplicate or out-of-order")
			}
			mostRecentElem = 1
			contentType, err := tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("Error decoding ContentType")
			}
			m.contentType = new(uint64)
			*m.contentType = contentType
		case tlv.FreshnessPeriod:
			if mostRecentElem >= 2 {
				return nil, errors.New("FreshnessPeriod is duplicate or out-of-order")
			}
			mostRecentElem = 2
			freshnessPeriod, err := tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("Error decoding FreshnessPeriod")
			}
			m.freshnessPeriod = new(time.Duration)
			*m.freshnessPeriod = time.Duration(freshnessPeriod) * time.Millisecond
		case tlv.FinalBlockID:
			if mostRecentElem == 3 {
				return nil, errors.New("FinalBlockId is duplicate or out-or-order")
			}
			mostRecentElem = 3
			elem.Parse()
			if len(elem.Subelements()) != 1 {
				return nil, errors.New("FinalBlockId must contain exactly one name component")
			}
			finalBlockID, err := DecodeNameComponent(elem.Subelements()[0])
			if err != nil {
				return nil, errors.New("Error decoding FinalBlockId")
			}
			m.finalBlockID = finalBlockID
		default:
			if tlv.IsCritical(elem.Type()) {
				return nil, tlv.ErrUnrecognizedCritical
			}
			// If non-critical, ignore
		}
	}
	return m, nil
}

func (m *MetaInfo) String() string {
	str := "MetaInfo("

	isFirst := true
	if m.contentType != nil {
		if !isFirst {
			str += ", "
		}
		str += "ContentType=" + strconv.FormatUint(*m.contentType, 10)
		isFirst = false
	}
	if m.freshnessPeriod != nil {
		if !isFirst {
			str += ", "
		}
		str += "FreshnessPeriod=" + strconv.FormatInt(m.freshnessPeriod.Milliseconds(), 10) + "ms"
	}
	if m.finalBlockID != nil {
		if !isFirst {
			str += ", "
		}
		str += "FinalBlockId=" + m.finalBlockID.String()
	}

	str += ")"
	return str
}

// DeepCopy returns a deep copy of the MetaInfo.
func (m *MetaInfo) DeepCopy() *MetaInfo {
	copyM := new(MetaInfo)
	if m.contentType != nil {
		copyM.contentType = new(uint64)
		*copyM.contentType = *m.contentType
	}
	if m.freshnessPeriod != nil {
		copyM.freshnessPeriod = new(time.Duration)
		*copyM.freshnessPeriod = *m.freshnessPeriod
	}
	if m.finalBlockID != nil {
		copyM.finalBlockID = m.finalBlockID.DeepCopy()
	}
	if m.wire != nil {
		copyM.wire = m.wire.DeepCopy()
	}
	return copyM
}

// ContentType returns the ContentType set in the MetaInfo.
func (m *MetaInfo) ContentType() *uint64 {
	if m.contentType == nil {
		return nil
	}

	contentType := new(uint64)
	*contentType = *m.contentType
	return contentType
}

// SetContentType sets the ContentType in the MetaInfo.
func (m *MetaInfo) SetContentType(contentType uint64) {
	m.contentType = new(uint64)
	*m.contentType = contentType
	m.wire = nil
}

// UnsetContentType unsets the ContentType in the MetaInfo.
func (m *MetaInfo) UnsetContentType() {
	m.contentType = nil
	m.wire = nil
}

// FreshnessPeriod returns the FreshnessPeriod set in the MetaInfo.
func (m *MetaInfo) FreshnessPeriod() *time.Duration {
	if m.freshnessPeriod == nil {
		return nil
	}

	freshnessPeriod := new(time.Duration)
	*freshnessPeriod = *m.freshnessPeriod
	return freshnessPeriod
}

// SetFreshnessPeriod sets the FreshnessPeriod in the MetaInfo.
func (m *MetaInfo) SetFreshnessPeriod(freshnessPeriod time.Duration) {
	m.freshnessPeriod = new(time.Duration)
	*m.freshnessPeriod = freshnessPeriod
	m.wire = nil
}

// UnsetFreshnessPeriod unsets the FreshnessPeriod in the MetaInfo.
func (m *MetaInfo) UnsetFreshnessPeriod() {
	m.freshnessPeriod = nil
	m.wire = nil
}

// FinalBlockID returns the FinalBlockId set in the MetaInfo.
func (m *MetaInfo) FinalBlockID() NameComponent {
	if m.finalBlockID == nil {
		return nil
	}

	return m.finalBlockID.DeepCopy()
}

// SetFinalBlockID sets the FinalBlockId in the MetaInfo.
func (m *MetaInfo) SetFinalBlockID(finalBlockID NameComponent) {
	m.finalBlockID = finalBlockID.DeepCopy()
	m.wire = nil
}

// UnsetFinalBlockID unsets the FinalBlockId in the MetaInfo.
func (m *MetaInfo) UnsetFinalBlockID() {
	m.finalBlockID = nil
	m.wire = nil
}

// Encode encodes the MetaInfo into a block.
func (m *MetaInfo) Encode() (*tlv.Block, error) {
	if m.wire != nil {
		return m.wire.DeepCopy(), nil
	}

	m.wire = new(tlv.Block)
	m.wire.SetType(tlv.MetaInfo)

	// ContentType
	if m.contentType != nil {
		m.wire.Append(tlv.EncodeNNIBlock(tlv.ContentType, *m.contentType))
	}

	// FreshnessPeriod
	if m.freshnessPeriod != nil {
		m.wire.Append(tlv.EncodeNNIBlock(tlv.FreshnessPeriod, uint64(m.freshnessPeriod.Milliseconds())))
	}

	// FinalBlockId
	if m.finalBlockID != nil {
		encodedComponent, err := m.finalBlockID.Encode().Wire()
		if err != nil {
			return nil, errors.New("Unable to encode FinalBlockId")
		}
		m.wire.Append(tlv.NewBlock(tlv.FinalBlockID, encodedComponent))
	}

	m.wire.Wire()
	return m.wire.DeepCopy(), nil
}

// HasWire returns whether the MetaInfo has an existing valid wire encoding.
func (m *MetaInfo) HasWire() bool {
	return m.wire != nil
}
