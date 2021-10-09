package disgo

type ClientOptions struct {
	Intents []Intent
}

func (o ClientOptions) GetIntentInt() int {
	intentInt := 0

	for _, i := range o.Intents {
		intentInt |= i.BitCode
	}
	return intentInt
}