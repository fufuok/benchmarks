//go:generate msgp
package msgp

type MessageCorpus int32

const (
	Message_UNIVERSAL MessageCorpus = 0
	Message_WEB       MessageCorpus = 1
	Message_IMAGES    MessageCorpus = 2
	Message_LOCAL     MessageCorpus = 3
	Message_NEWS      MessageCorpus = 4
	Message_PRODUCTS  MessageCorpus = 5
	Message_VIDEO     MessageCorpus = 6
)

type Message struct {
	Query         string
	PageNumber    int32
	ResultPerPage int32
	Comment       []byte
	Corpus        MessageCorpus
}
