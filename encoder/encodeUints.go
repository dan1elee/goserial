package encoder

import (
	"bytes"
	"encoding/binary"
	"goserialize/enums"
	"unsafe"
)

func encodeUint(v uint) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	intSize := unsafe.Sizeof(v)
	if intSize == 4 {
		err = binary.Write(buf, binary.LittleEndian, uint32(v))
	} else {
		err = binary.Write(buf, binary.LittleEndian, uint64(v))
	}
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.UINT, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeUint8(v uint8) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.UINT8, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeUint16(v uint16) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.UINT16, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeUint32(v uint32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.UINT32, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeUint64(v uint64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.UINT64, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}
