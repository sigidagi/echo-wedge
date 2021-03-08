package client

import "time"

// Error is struct for unmarshalling error message from API gtw
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSONRPCResponseNetwork is struct for Sending API gtw's response
type JSONRPCResponseNetwork struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value struct {
			Name   string `json:"name"`
			Device []struct {
				Value []struct {
					State []struct {
						Data      string    `json:"data"`
						Type      string    `json:"type"`
						Timestamp time.Time `json:"timestamp"`
						Meta      struct {
							ID      string `json:"id"`
							Type    string `json:"type"`
							Version string `json:"version"`
						} `json:"meta"`
					} `json:"state"`
					Name       string `json:"name"`
					Permission string `json:"permission"`
					Status     string `json:"status"`
					Number     struct {
						Min  float64 `json:"min"`
						Max  int     `json:"max"`
						Step int     `json:"step"`
						Unit string  `json:"unit"`
					} `json:"number"`
					Meta struct {
						ID      string `json:"id"`
						Type    string `json:"type"`
						Version string `json:"version"`
					} `json:"meta"`
				} `json:"value"`
				Name          string `json:"name"`
				Manufacturer  string `json:"manufacturer"`
				Product       string `json:"product"`
				Serial        string `json:"serial"`
				Description   string `json:"description"`
				Protocol      string `json:"protocol"`
				Communication string `json:"communication"`
				Meta          struct {
					ID      string `json:"id"`
					Type    string `json:"type"`
					Version string `json:"version"`
				} `json:"meta"`
			} `json:"device"`
			Meta struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Version string `json:"version"`
			} `json:"meta"`
		} `json:"value"`
		Meta struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseValueDevice is struct for unmarshalling value and device data into got from API gtw
type JSONResponseValueDevice struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value struct {
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
		} `json:"value"`
		Meta struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}


// JSONResponseOneDevice is struct for unmarshalling data into got from API gtw
type JSONResponseOneDevice struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value struct {
			Value []struct {
				State []struct {
					Data      string    `json:"data"`
					Type      string    `json:"type"`
					Timestamp time.Time `json:"timestamp"`
					Meta      struct {
						ID      string `json:"id"`
						Type    string `json:"type"`
						Version string `json:"version"`
					} `json:"meta"`
				} `json:"state"`
				Name       string `json:"name"`
				Permission string `json:"permission"`
				Status     string `json:"status"`
				Number     struct {
					Min  float64 `json:"min"`
					Max  int     `json:"max"`
					Step int     `json:"step"`
					Unit string  `json:"unit"`
				} `json:"number"`
				Meta struct {
					ID      string `json:"id"`
					Type    string `json:"type"`
					Version string `json:"version"`
				} `json:"meta"`
			} `json:"value"`
			Name          string `json:"name"`
			Manufacturer  string `json:"manufacturer"`
			Product       string `json:"product"`
			Serial        string `json:"serial"`
			Description   string `json:"description"`
			Protocol      string `json:"protocol"`
			Communication string `json:"communication"`
			Meta          struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Version string `json:"version"`
			} `json:"meta"`
		} `json:"value"`
		Meta struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}
// JSONResponseOneValue is struct for unmarshalling data into got from API gtw
type JSONResponseOneValue struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value struct {
			State []struct {
				Data      string    `json:"data"`
				Type      string    `json:"type"`
				Timestamp time.Time `json:"timestamp"`
				Meta      struct {
					ID      string `json:"id"`
					Type    string `json:"type"`
					Version string `json:"version"`
				} `json:"meta"`
			} `json:"state"`
			Name       string `json:"name"`
			Permission string `json:"permission"`
			Status     string `json:"status"`
			Number     struct {
				Min  float64 `json:"min"`
				Max  int     `json:"max"`
				Step int     `json:"step"`
				Unit string  `json:"unit"`
			} `json:"number"`
			Meta struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Version string `json:"version"`
			} `json:"meta"`
		} `json:"value"`
		Meta struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseState is struct for unmarshalling data into got from API gtw
type JSONResponseState struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value struct {
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
		} `json:"value"`
		Meta struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}

// JSONResponseOneState is struct for unmarshalling data into got from API gtw
type JSONResponseOneState struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Value struct {
			Data      string    `json:"data"`
			Type      string    `json:"type"`
			Timestamp time.Time `json:"timestamp"`
			Meta      struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Version string `json:"version"`
			} `json:"meta"`
		} `json:"value"`
		Meta struct {
			ServerSendTime string `json:"server_send_time"`
		} `json:"meta"`
	} `json:"result"`
	Error Error `json:"error"`
}