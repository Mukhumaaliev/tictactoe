package web

// GameRequest представляет запрос клиента с обновленным игровым полем
type GameRequest struct {
	ID    string    `json:"id,omitempty"`
	Board [3][3]int `json:"board"`
}

// GameResponse представляет ответ сервера с текущим состоянием игры
type GameResponse struct {
	ID     string    `json:"id"`
	Board  [3][3]int `json:"board"`
	Status string    `json:"status"`
}

// ErrorResponse представляет формат ответа при возникновении ошибки
type ErrorResponse struct {
	Error string `json:"error"`
}
