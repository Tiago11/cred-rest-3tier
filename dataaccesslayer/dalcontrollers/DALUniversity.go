package dalcontrollers

import (
  "database/sql"
  "credits_be/models"
  //"credits_be/dataaccesslayer"
)

type DALUniversity struct {

}

func (dalU *DALUniversity) Add(u *models.University) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  if err := txn.Insert(u); err != nil {
    panic(err)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return nil
}

func (dalU *DALUniversity) Get(id int) *models.University {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  u, err := txn.Get(models.University{}, id)
  if err != nil {
    panic(err)
  }
  if u == nil {
    return nil
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return u.(*models.University)
}

func (dalU *DALUniversity) GetWithDegrees(id int) *models.University {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  u, err := txn.Get(models.University{}, id)
  if err != nil {
    panic(err)
  }
  if u == nil {
    return nil
  }

  // Get all the university's degrees.
  var degrees []*models.Degree
  _, err2 := txn.Select(&degrees,
    `select * from "Degree" where "UniversityId"=$1;`, id)
  if err2 != nil {
    panic(err)
  }
  u.(*models.University).Degrees = append(u.(*models.University).Degrees, degrees...)

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return u.(*models.University)
}

func (dalU *DALUniversity) GetAll() []*models.University {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  results, err := txn.Select(models.University{},
    `select * from "University"`)
  if err != nil {
    panic(err)
    panic(results)
  }

  var universities []*models.University
  for _, r := range results {
    u := r.(*models.University)
    universities = append(universities, u)
  }

  if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
    if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
      panic(err)
    }
  }

  return universities
}

func (dalU *DALUniversity) Update(u *models.University) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  success, err := txn.Update(u)
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

func (dalU *DALUniversity) Delete(id int) error {
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }

  success, err := txn.Delete(&models.University{ UniversityId: id })
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
