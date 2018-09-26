package amigo

// Originate holds information used to originate outgoing calls.
//		Channel - Channel name to call.
//		Exten - Extension to use (requires Context and Priority)
//		Context - Context to use (requires Exten and Priority)
//		Priority - Priority to use (requires Exten and Context)
//		Application - Application to execute.
//		Data - Data to use (requires Application).
//		Timeout - How long to wait for call to be answered (in ms.).
//		CallerID - Caller ID to be set on the outgoing channel.
//		Variable - Channel variable to set, multiple Variable: headers are allowed.
//		Account - Account code.
//		Async - Set to true for fast origination.
//		Codecs - Comma-separated list of codecs to use for this call.
type Originate struct {
	Channel     string `ami:"Channel"`
	Exten       string `ami:"Exten"`
	Context     string `ami:"Context"`
	Priority    int    `ami:"Priority"`
	Application string `ami:"Application"`
	Data        string `ami:"Data"`
	Timeout     int    `ami:"Timeout"`
	CallerID    string `ami:"CallerID"`
	Variable    string `ami:"Variable"`
	Account     string `ami:"Account"`
	Async       string `ami:"Async"`
	Codecs      string `ami:"Codecs"`
}
