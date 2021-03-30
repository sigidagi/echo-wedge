package models

// Params is struct for params inside JSONRPC
type Params struct {
	URL  string      `json:"url"`
	Data interface{} `json:"data"`
}

// JSONRPC is main struct for exchanging data
type JSONRPC struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

// Error is struct for unmarshalling error message from API gtw
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ItemList struct {
	Child []struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"child"`
	ID    []string `json:"id"`
	Mode  bool     `json:"mode"`
	Limit int      `json:"limit"`
	Count int      `json:"count"`
	Meta  struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"meta"`
}

// JSONRPCResponseNetwork is struct for Sending API gtw's response
type JSONResponseNetwork struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value Network `json:"value"`
		Meta  struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseValueDevice is struct for unmarshalling value and device data into got from API gtw
type JSONResponseItemList struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value ItemList `json:"value"`
		Meta  struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseOneDevice is struct for unmarshalling data into got from API gtw
type JSONResponseDevice struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value Device `json:"value"`
		Meta  struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseOneValue is struct for unmarshalling data into got from API gtw
type JSONResponseValue struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value Value `json:"value"`
		Meta  struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseOneState is struct for unmarshalling data into got from API gtw
type JSONResponseState struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value State `json:"value"`
		Meta  struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}
