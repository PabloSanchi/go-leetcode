package commons

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"io"
	"log/slog"
)

func Marshal(msg *[]byte, protoMsg proto.Message) error {
	encodedMsg, err := proto.Marshal(protoMsg)
	if err != nil {
		slog.Error("could not marshal message", "error", err)
		return err
	}

	var buffer bytes.Buffer
	msgLength := uint32(len(encodedMsg))

	if err = binary.Write(&buffer, binary.BigEndian, msgLength); err != nil {
		slog.Error("could not encode message length", "error", err)
		return err
	}

	if _, err = buffer.Write(encodedMsg); err != nil {
		slog.Error("could not write message", "error", err)
		return err
	}

	*msg = buffer.Bytes()
	return nil
}

func ReadMsg(reader *bufio.Reader, dest *[]byte) error {
	lengthBytes := make([]byte, 4)
	if _, err := io.ReadFull(reader, lengthBytes); err != nil {
		slog.Error("could not read message length", "error", err)
		return err
	}

	msgLength := binary.BigEndian.Uint32(lengthBytes)
	msg := make([]byte, msgLength)
	if _, err := io.ReadFull(reader, msg); err != nil {
		slog.Error("could not read message", "error", err)
		return err
	}

	*dest = msg
	return nil
}
