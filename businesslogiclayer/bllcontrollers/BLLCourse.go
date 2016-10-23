package bllcontrollers

import (
  "credits_be/models"
  "credits_be/dataaccesslayer/dalinterfaces"
  "credits_be/dataaccesslayer/dalcontrollers"
)

type BLLCourse struct {
  idalCourse dalinterfaces.IDALCourse
}

func (bllC *BLLCourse) InitBLLCourse() {
  bllC.idalCourse = &dalcontrollers.DALCourse{}
}

func (bllC *BLLCourse) Add(co *models.Course) error {
  err := bllC.idalCourse.Add(co)
  if err != nil {
    panic(err)
  }

  return nil
}

func (bllC *BLLCourse) Get(id int) *models.Course {
  return bllC.idalCourse.Get(id)
}

func (bllC *BLLCourse) GetAll() []*models.Course {
  return bllC.idalCourse.GetAll()
}

func (bllC *BLLCourse) Update(co *models.Course) error {
  err := bllC.idalCourse.Update(co)
  if err != nil {
    panic(err)
  }

  return nil
}

func (bllC *BLLCourse) Delete(id int) error {
  err := bllC.idalCourse.Delete(id)
  if err != nil {
    panic(err)
  }

  return nil
}
