package controllers

import (
	"github.com/LavGo/cportal/utils"
	"github.com/LavGo/cportal/models"
)

type ApiMainController struct {
	baseController
}

func (this *ApiMainController)Get()  {
	var jsondata models.PortalData
	utils.ParseJson("conf/urls.json",&jsondata)
	this.Data["json"]=jsondata
	this.ServeJSON()
}