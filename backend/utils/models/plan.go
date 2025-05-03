package models

// PlanRequest は /plan エンドポイントで受け取るリクエストの構造体
type PlanRequest struct {
	Goal           string `json:"goal"`             // 例: "英語を話せるようになる"
	Period         string `json:"period"`           // 例: "3ヶ月"
	SessionCount   *int   `json:"session_count"`    // 任意：20など。未入力ならnil
}
