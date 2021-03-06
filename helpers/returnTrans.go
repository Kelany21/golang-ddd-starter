package helpers

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

func DoneUpdate(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_update_item")
}

func DoneDelete(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_delete_item")
}

func DoneActivate(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g),"activated")
}

func DoneTrash(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g),"trashed")
}

func DoneDeactivate(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "deactivated")
}

func DoneGetItem(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_get_item")
}

func DoneCreateItem(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_created_item")
}

func DoneGetAllItems(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "get_all_items")
}

func ItemNotFound(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "item_not_found")
}

func Wrong(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "wrong")
}
