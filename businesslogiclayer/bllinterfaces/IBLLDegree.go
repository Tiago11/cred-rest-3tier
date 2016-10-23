package bllinterfaces

import (
  "credits_be/models"
)

type IBLLDegree interface {

  InitBLLDegree()

  Add(d *models.Degree) error

  Get(id int) *models.Degree

  GetWithCourses(id int) *models.Degree

  GetAll() []*models.Degree

  Update(*models.Degree) error

  Delete(id int) error

}
