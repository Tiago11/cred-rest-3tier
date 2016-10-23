package dalcontrollers

import (
  "github.com/go-gorp/gorp"
  r "github.com/revel/revel"
  _ "github.com/lib/pq"
  "github.com/revel/modules/db/app"
  "credits_be/models"
)

var (
  Dbm *gorp.DbMap
)

func InitDB() {
  db.Init()
  Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

  setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
    for col, size := range colSizes {
      t.ColMap(col).MaxSize = size
    }
  }

  t := Dbm.AddTable(models.University{}).SetKeys(true, "UniversityId")
  t.ColMap("Degrees").Transient = true
  setColumnSizes(t, map[string]int{
    "Name":     50,
    "Country":  50,
  })

  t = Dbm.AddTable(models.Degree{}).SetKeys(true, "DegreeId")
  t.ColMap("University").Transient = true
  t.ColMap("Courses").Transient = true
  setColumnSizes(t, map[string]int{
    "Name": 50,
    "TotalCredits": 20,
  })

  t = Dbm.AddTable(models.Course{}).SetKeys(true, "CourseId")
  t.ColMap("Degree").Transient = true
  setColumnSizes(t, map[string]int{
    "Name":     50,
    "Credits":  20,
  })

  Dbm.TraceOn("[gorp]", r.INFO)
  Dbm.CreateTables()

  university := &models.University{0, "UdelaR", "Uruguay", nil}
  if err := Dbm.Insert(university); err != nil {
    panic(err)
  }

  university2 := &models.University{0, "UBA", "Argentina", nil}
  if err := Dbm.Insert(university2); err != nil {
    panic(err)
  }

  degree := &models.Degree{0, 1, "Ingenieria en Computacion", 450, university, nil}
  if err := Dbm.Insert(degree); err != nil {
    panic(err)
  }

  course := &models.Course{0, 1, "MAA", 10, degree}
  if err := Dbm.Insert(course); err != nil {
    panic(err)
  }

}
