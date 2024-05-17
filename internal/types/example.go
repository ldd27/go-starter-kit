package types

type GetExamplePageListReq struct {
	PageReq
}

type GetExampleCursorPageListReq struct {
	CursorPageReq
}

type CreateExampleReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateExampleReq struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type DeleteExampleReq struct {
	ID int `json:"id" validate:"required"`
}

type Example struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
