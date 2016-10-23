package bllinterfaces

import (
  "credits_be/models"
)

type IBLLCourse interface {

  InitBLLCourse()

  Add(co *models.Course) error

  Get(id int) *models.Course

  GetAll() []*models.Course

  Update(co *models.Course) error

  Delete(id int) error

}
