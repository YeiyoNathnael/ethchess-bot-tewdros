package lichess

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLichessUser(username string) int64 {

	var _user User

	url := "https://lichess.org/api/user/" + username

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	err := json.Unmarshal(body, &_user)
	if err != nil {
		fmt.Errorf("Error unmarshaling:", err)
	}

	return _user.Perfs.Bullet.Rating

}
