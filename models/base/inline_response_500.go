package base

type InlineResponse500 struct {
	// Описание ошибки
	Message string `json:"message"`
	// Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId string `json:"request_id,omitempty"`
	// Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code int32 `json:"code,omitempty"`
}
