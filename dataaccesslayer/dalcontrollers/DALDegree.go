package dalcontrollers

import (
  "database/sql"
  "credits_be/models"
)

type DALDegree struct {

}

func (dalD *DALDegree) Add(d *models.Degree) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  if err := txn.Insert(d); err != nil {
    panic(err)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return nil
}

func (dalD *DALDegree) Get(id int) *models.Degree {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  d, err := txn.Get(models.Degree{}, id)
  if err != nil {
    panic(err)
  }
  if d == nil {
    return nil
  }

  u, err := txn.Get(models.University{}, d.(*models.Degree).UniversityId)
  if err != nil {
    panic(err)
  }
  if u == nil {
    return nil
  }
  d.(*models.Degree).University = u.(*models.University)

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return d.(*models.Degree)
}

func (dalD *DALDegree) GetWithCourses(id int) *models.Degree {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  d, err := txn.Get(models.Degree{}, id)
  if err != nil {
    panic(err)
  }
  if d == nil {
    return nil
  }

  u, err := txn.Get(models.University{}, d.(*models.Degree).UniversityId)
  if err != nil {
    panic(err)
  }
  if u == nil {
    return nil
  }
  d.(*models.Degree).University = u.(*models.University)

  // Get all the degree's courses.
  var courses []*models.Course
  _, err2 := txn.Select(&courses,
    `select * from "Course" where "DegreeId"=$1;`, id)
  if err2 != nil {
    panic(err)
  }
  d.(*models.Degree).Courses = append(d.(*models.Degree).Courses, courses...)

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return d.(*models.Degree)
}

func (dalD *DALDegree) GetAll() []*models.Degree {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  results, err := txn.Select(models.Degree{},
    `select * from "Degree"`)
  if err != nil {
    panic(err)
  }

  var degrees []*models.Degree
  for _, r := range results {
    d := r.(*models.Degree)
    degrees = append(degrees, d)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return degrees
}

func (dalD *DALDegree) Update(d *models.Degree) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  success, err := txn.Update(d)
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

func (dalD *DALDegree) Delete(id int) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  success, err := txn.Delete(&models.Degree{ DegreeId: id })
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
