package sarama

import "testing"

var (
	offsetCommitRequestNoBlocksV0 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x00}

	offsetCommitRequestNoBlocksV1 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x00}

	offsetCommitRequestNoBlocksV2 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x33,
		0x00, 0x00, 0x00, 0x00}

	offsetCommitRequestOneBlockV0 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x52, 0x21,
		0x00, 0x00, 0x00, 0x00, 0xDE, 0xAD, 0xBE, 0xEF,
		0x00, 0x08, 'm', 'e', 't', 'a', 'd', 'a', 't', 'a'}

	offsetCommitRequestOneBlockV1 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x52, 0x21,
		0x00, 0x00, 0x00, 0x00, 0xDE, 0xAD, 0xBE, 0xEF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x08, 'm', 'e', 't', 'a', 'd', 'a', 't', 'a'}

	offsetCommitRequestOneBlockV2 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x33,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x52, 0x21,
		0x00, 0x00, 0x00, 0x00, 0xDE, 0xAD, 0xBE, 0xEF,
		0x00, 0x08, 'm', 'e', 't', 'a', 'd', 'a', 't', 'a'}
)

func TestOffsetCommitRequest(t *testing.T) {
	request := new(OffsetCommitRequest)

	request.ConsumerGroup = "foobar"
	testEncodable(t, "no blocks v0", request, offsetCommitRequestNoBlocksV0)

	request.ConsumerGroupGeneration = 0x1122
	request.ConsumerID = "cons"
	request.Version = 1
	testEncodable(t, "no blocks v1", request, offsetCommitRequestNoBlocksV1)

	request.RetentionTime = 0x4433
	request.Version = 2
	testEncodable(t, "no blocks v2", request, offsetCommitRequestNoBlocksV2)

	request.AddBlock("topic", 0x5221, 0xDEADBEEF, ReceiveTime, "metadata")
	request.Version = 0
	testEncodable(t, "one block v0", request, offsetCommitRequestOneBlockV0)

	request.Version = 1
	testEncodable(t, "one block v1", request, offsetCommitRequestOneBlockV1)

	request.Version = 2
	testEncodable(t, "one block v2", request, offsetCommitRequestOneBlockV2)
}
