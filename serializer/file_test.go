package serializer

import (
	"pcbook/pb"
	"pcbook/sample"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFileName := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFileName)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = ReadProtobufFromBinaryFile(binaryFileName, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

	jsonFileName := "../tmp/laptop.json"

	err = WriteProtobufToJSONFile(laptop1, jsonFileName)
	require.NoError(t, err)

	laptop3 := &pb.Laptop{}

	err = ReadProtobufFromJSONFile(jsonFileName, laptop3)

	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop3))
}
