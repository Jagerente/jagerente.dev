package main

import (
	"log"
	"net/http"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/jagerente/jagerente.xyz/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Members struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

type Member struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	CanAccessClosed bool   `json:"can_access_closed"`
	IsClosed        bool   `json:"is_closed"`
}

func handleRequest() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/api/getDons", getDons)

	//HTTP
	r.Run(config.ENV.Host)
}

func main() {
	config.Init()
	handleRequest()
}

func getDons(c *gin.Context) {
	groupId := c.Query("group_id")
	token := c.Query("access_token")

	vk := api.NewVK(token)

	p := params.NewGroupsGetMembersBuilder()
	p.GroupID(groupId)
	p.Filter("donut")
	p.Lang(object.LangRU)

	donsList, err := vk.GroupsGetMembersFields(api.Params(p.Params))
	if err != nil {
		log.Fatal(err)
	}

	var dons []string

	for _, don := range donsList.Items {
		dons = append(dons, don.FirstName+" "+don.LastName)
	}
	c.JSON(http.StatusOK, dons)
}
