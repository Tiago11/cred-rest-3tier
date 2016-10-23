package dalinterfaces

import (
  "credits_be/models"
)

type IDALDegree interface {

    Add(d *models.Degree) error

    Get(id int) *models.Degree

    GetWithCourses(id int) *models.Degree

    GetAll() []*models.Degree

    Update(d *models.Degree) error

    Delete(id int) error

}
