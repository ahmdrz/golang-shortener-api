package googleshortener

type Response struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
	Status  string `json:"status"`
}

type Error struct {
}
