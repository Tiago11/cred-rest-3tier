package bllcontrollers

import (
  "credits_be/models"
  "credits_be/dataaccesslayer/dalinterfaces"
  "credits_be/dataaccesslayer/dalcontrollers"
)

type BLLDegree struct {
  idalDegree dalinterfaces.IDALDegree
}

func (bllD *BLLDegree) InitBLLDegree() {
  bllD.idalDegree = &dalcontrollers.DALDegree{}
}

func (bllD *BLLDegree) Add(d *models.Degree) error {
  err := bllD.idalDegree.Add(d)
  if err != nil {
    panic(err)
  }

  return nil
}

func (bllD *BLLDegree) Get(id int) *models.Degree {
  return bllD.idalDegree.Get(id)
}

func (bllD *BLLDegree) GetWithCourses(id int) *models.Degree {
  return bllD.idalDegree.GetWithCourses(id)
}

func (bllD *BLLDegree) GetAll() []*models.Degree {
  return bllD.idalDegree.GetAll()
}

func (bllD *BLLDegree) Update(d *models.Degree) error {
  err := bllD.idalDegree.Update(d)
  if err != nil {
    panic(err)
  }

  return nil
}

func (bllD *BLLDegree) Delete(id int) error {
  err := bllD.idalDegree.Delete(id)
  if err != nil {
    panic(err)
  }

  return nil
}
