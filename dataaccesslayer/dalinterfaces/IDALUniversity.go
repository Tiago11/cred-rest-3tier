package dalinterfaces

import (
  "credits_be/models"
)

type IDALUniversity interface {

    Add(u *models.University) error

    Get(id int) *models.University

    GetWithDegrees(id int) *models.University

    GetAll() []*models.University

    Update(u *models.University) error

    Delete(id int) error

}
