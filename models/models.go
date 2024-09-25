package models

type MatchData struct {
	username string
	domain   string
	name     string
}

type PollResponse struct {
	Data struct {
		Id        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`

	Support struct {
		Url  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

// {
//   "data": {
//     "id": 2,
//     "email": "janet.weaver@reqres.in",
//     "first_name": "Janet",
//     "last_name": "Weaver",
//     "avatar": "https://reqres.in/img/faces/2-image.jpg"
//   },
//   "support": {
//     "url": "https://reqres.in/#support-heading",
//     "text": "To keep ReqRes free, contributions towards server costs are appreciated!"
//   }
// }
