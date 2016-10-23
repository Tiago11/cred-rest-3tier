package bllcontrollers

import (
  "credits_be/dataaccesslayer/dalcontrollers"
)

func BLLInitDB() {
  dalcontrollers.InitDB()
}
