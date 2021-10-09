package disgo

type Client struct {
	Token string
	Bot bool
	Options ClientOptions
}

func NewClient(token string, bot bool, options ClientOptions) *Client {
	return &Client{
		token,
		bot,
		options,
	}
}