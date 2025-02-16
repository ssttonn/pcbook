package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)

	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to json: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("cannot write json data to file: %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)

	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}

func ReadProtobufFromJSONFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("cannot read json data from file: %w", err)
	}

	err = JSONToProtobuf(data, message)

	if err != nil {
		return fmt.Errorf("cannot unmarshal json to proto message: %w", err)
	}

	return nil
}
