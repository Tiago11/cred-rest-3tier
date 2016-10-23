package dalinterfaces

import (
  "credits_be/models"
)

type IDALCourse interface {

  Add(c *models.Course) error

  Get(id int) *models.Course

  GetAll() []*models.Course

  Update(c *models.Course) error

  Delete(id int) error

}
