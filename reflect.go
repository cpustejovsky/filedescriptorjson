package filedescriptorjson

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

func MessageDescriptorFromFileDescriptor(name string, fd protoreflect.FileDescriptor, resolver protodesc.Resolver) (protoreflect.MessageDescriptor, error) {
	fdp := protodesc.ToFileDescriptorProto(fd)
	f, err := protodesc.NewFile(fdp, resolver)
	if err != nil {
		return nil, fmt.Errorf("error creating NewFile: %w", err)
	}

	md := f.Messages().ByName(protoreflect.Name(name))

	if md == nil {
		return nil, errors.New("message descriptor was nil")
	}

	return md, nil
}

func ToJSON(name string, fd protoreflect.FileDescriptor, data []byte) (string, error) {
	md, err := MessageDescriptorFromFileDescriptor(name, fd, nil)
	if err != nil {
		return "", err
	}
	msg := dynamicpb.NewMessage(md)
	err = proto.Unmarshal(data, msg)
	if err != nil {
		return "", err
	}
	return protojson.Format(msg), nil
}
