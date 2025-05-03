package handlers

import (
	"fmt"
	"net/http"

	"main/utils/models"

	"github.com/gin-gonic/gin"
)

// /planエンドポイントのリクエストを処理し、計画案をJSONで返すハンドラ
func HandlePlan(c *gin.Context) {
	var req models.PlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	plan := []string{
		"ステップ1: 目標を確認 - " + req.Goal,
		"ステップ2: 期間を設定 - " + req.Period,
	}

	if req.SessionCount != nil {
		plan = append(plan, "ステップ3: セッション数は "+fmt.Sprintf("%d", *req.SessionCount)+" 回")
	} else {
		plan = append(plan, "ステップ3: セッション数は未指定")
	}

	c.JSON(http.StatusOK, gin.H{
		"goal":          req.Goal,
		"period":        req.Period,
		"session_count": req.SessionCount,
		"plan":          plan,
	})
}
