package controllers

import (
  "github.com/revel/revel"
  "credits_be/businesslogiclayer/bllcontrollers"
)


func init() {
  revel.OnAppStart(bllcontrollers.BLLInitDB)
}
