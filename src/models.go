package src

type IletiMerkeziSms struct {
	Request IMRequest `json:"request"`
}

type IMRequest struct {
	Authentication IMAuthentication `json:"authentication"`
	Order          IMOrder          `json:"order"`
}
type IMAuthentication struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

type IMOrder struct {
	Sender string `json:"sender"`
	// sendDateTime map[string]interface{}
	Iys     string    `json:"iys"`
	IysList string    `json:"iysList"`
	Message IMMessage `json:"message"`
}
type IMMessage struct {
	Text       string       `json:"text"`
	Receipents IMReceipents `json:"receipents"`
}
type IMReceipents struct {
	Number []string `json:"number"`
}

type Message struct {
	Addresses []string `json:"Addresses"`
	Messages  []string `json:"Messages"`
}
