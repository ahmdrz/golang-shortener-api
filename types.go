package googleshortener

type Response struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type Error struct {
}
