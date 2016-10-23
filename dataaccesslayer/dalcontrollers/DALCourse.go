package dalcontrollers

import (
  "database/sql"
  "credits_be/models"
)

type DALCourse struct {

}

func (dalC *DALCourse) Add(c *models.Course) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  err2 := txn.Insert(c)
  if err2 != nil {
    panic(err)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return nil
}

func (dalC *DALCourse) Get(id int) *models.Course {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  co, err := txn.Get(models.Course{}, id)
  if err != nil {
    panic(err)
  }
  if co == nil {
    return nil
  }

  d, err := txn.Get(models.Degree{}, co.(*models.Course).DegreeId)
  if err != nil {
    panic(err)
  }
  if d == nil {
    return nil
  }
  co.(*models.Course).Degree = d.(*models.Degree)

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return co.(*models.Course)
}

func (dalC *DALCourse) GetAll() []*models.Course {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  results, err := txn.Select(models.Course{},
    `select * from "Course"`)
  if err != nil {
    panic(err)
  }

  var courses []*models.Course
  for _, r := range results {
    co := r.(*models.Course)
    courses = append(courses, co)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return courses
}

func (dalC *DALCourse) Update(c *models.Course) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  success, err := txn.Update(c)
  if err != nil || success == 0 {
    panic(err)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return nil
}

func (dalC *DALCourse) Delete(id int) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  success, err := txn.Delete(&models.Course{ CourseId: id })
  if err != nil || success == 0 {
    panic(err)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return nil
}
