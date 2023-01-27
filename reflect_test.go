package filedescriptionjson_test

import (
	"github.com/cpustejovsky/filedescriptorjson"
	"github.com/cpustejovsky/filedescriptorjson/helloworld"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"testing"
)

var protoname = "HelloRequest"

func TestMessageDescriptorFromFileDescriptor(t *testing.T) {
	fd := helloworld.File_helloworld_helloworld_proto
	md, err := filedescriptionjson.MessageDescriptorFromFileDescriptor(protoname, fd, nil)
	require.Nil(t, err)
	require.Equal(t, protoname, string(md.FullName().Name()))
}

func TestMarshalProtoFromFileDescriptor(t *testing.T) {
	msg := helloworld.HelloRequest{
		Name: "Charles",
	}
	t.Log("Original Message:\n", protojson.Format(&msg))
	bin, err := proto.Marshal(&msg)
	require.Nil(t, err)
	fd := helloworld.File_helloworld_helloworld_proto
	tmp := &helloworld.HelloRequest{}
	err = proto.Unmarshal(bin, tmp)
	require.Nil(t, err)
	require.Equal(t, msg.GetName(), tmp.GetName())
	newMsg, err := filedescriptionjson.FromFileDescriptor(protoname, fd, bin)
	require.Nil(t, err)
	t.Log("Message Converted from FileDescriptor:\n", protojson.Format(newMsg))
}
