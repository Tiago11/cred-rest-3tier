package bllcontrollers

import (
  "credits_be/models"
  "credits_be/dataaccesslayer/dalinterfaces"
  "credits_be/dataaccesslayer/dalcontrollers"
)

type BLLUniversity struct {
  idalUniversity dalinterfaces.IDALUniversity
}

func (bllU *BLLUniversity) InitBLLUniversity() {
   bllU.idalUniversity = &dalcontrollers.DALUniversity{}
}

func (bllU *BLLUniversity) Add(u *models.University) error {
  err := bllU.idalUniversity.Add(u)
  if err != nil {
    panic(err)
  }

  return nil
}

func (bllU *BLLUniversity) Get(id int) *models.University {
  return bllU.idalUniversity.Get(id)
}

func (bllU *BLLUniversity) GetWithDegrees(id int) *models.University {
  return bllU.idalUniversity.GetWithDegrees(id)
}

func (bllU *BLLUniversity) GetAll() []*models.University {
  return bllU.idalUniversity.GetAll()
}

func (bllU *BLLUniversity) Update(u *models.University) error {
  err := bllU.idalUniversity.Update(u)
  if err != nil {
    panic(err)
  }

  return nil
}

func (bllU *BLLUniversity) Delete(id int) error {
  err := bllU.idalUniversity.Delete(id)
  if err != nil {
    panic(err)
  }

  return nil
}
