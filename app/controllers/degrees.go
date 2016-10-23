package controllers

import (
  "github.com/revel/revel"
  "credits_be/models"
  "encoding/json"
  "net/http"
  "errors"
  jwt "github.com/dgrijalva/jwt-go"
  "fmt"
  "credits_be/businesslogiclayer/bllinterfaces"
  "credits_be/businesslogiclayer/bllcontrollers"
)

type Degrees struct {
  Application
}

func (c Degrees) Index() revel.Result {
  var ibllD bllinterfaces.IBLLDegree
  ibllD = &bllcontrollers.BLLDegree{}
  ibllD.InitBLLDegree()
  degrees := ibllD.GetAll()

  return c.RenderJson(degrees)
}

func (c Degrees) loadDegreeById(id int) *models.Degree {
  var ibllD bllinterfaces.IBLLDegree
  ibllD = &bllcontrollers.BLLDegree{}
  ibllD.InitBLLDegree()
  d := ibllD.Get(id)

  if d == nil {
    return nil
  }

  return d
}

func (c Degrees) loadDegreeWithCourses(id int) *models.Degree {
  //dalD := &dalcontrollers.DALDegree{}
  //d := dalD.GetWithCourses(id)
  var ibllD bllinterfaces.IBLLDegree
  ibllD = &bllcontrollers.BLLDegree{}
  ibllD.InitBLLDegree()
  d := ibllD.GetWithCourses(id)

  if d == nil {
    return nil
  }

  return d
}

func (c Degrees) Show(id int) revel.Result {
  degree := c.loadDegreeWithCourses(id)
  if degree == nil {
    return c.NotFound("Degree %d does not exist", id)
  }
  return c.RenderJson(degree)
}


func (c Degrees) parseDegree() (models.Degree, error) {
  degree := models.Degree{}
  err := json.NewDecoder(c.Request.Body).Decode(&degree)
  return degree, err
}

func (c Degrees) Create() revel.Result {
  // Check if the request is authorized by a JSON Web Token.
  tokenString := c.Request.Header.Get("Authorization")

  if err := c.checkJWTDegrees(tokenString); err != nil {
    c.Response.Status = http.StatusUnauthorized
    c.Response.ContentType = "application/json"
    // Change it to send a JSON with the message of the error, use the error 'err'.
    return c.RenderText("ERROR: You dont have authorization to perform this action.")
  }

  if degree, err := c.parseDegree(); err != nil {
    return c.RenderText("Unable to parse a degree from the JSON.")
  } else {
    // Validate the model.
    degree.Validate(c.Validation)
    if c.Validation.HasErrors() {
      // Do something better here!
      return c.RenderText("You have an error in your degree.")
    } else {
      var ibllD bllinterfaces.IBLLDegree
      ibllD = &bllcontrollers.BLLDegree{}
      ibllD.InitBLLDegree()
      if err := ibllD.Add(&degree); err != nil {
        return c.RenderText("Error inserting record for the degree.")
      } else {
        c.Response.Status = http.StatusCreated
        c.Response.ContentType = "application/json"
        return c.RenderJson(degree)
      }
    }
  }
}

func (c Degrees) Update(id int) revel.Result {
  // Check if the request is authorized by a JSON Web Token.
  tokenString := c.Request.Header.Get("Authorization")

  if err := c.checkJWTDegrees(tokenString); err != nil {
    c.Response.Status = http.StatusUnauthorized
    c.Response.ContentType = "application/json"
    // Change it to send a JSON with the message of the error, use the error 'err'.
    return c.RenderText("ERROR: You dont have authorization to perform this action.")
  }

  degree, err := c.parseDegree()
  if err != nil {
    return c.RenderText("Unable to parse a Degree from the JSON.")
  }
  // Ensure the Id is set.
  degree.DegreeId = id

  var ibllD bllinterfaces.IBLLDegree
  ibllD = &bllcontrollers.BLLDegree{}
  ibllD.InitBLLDegree()
  err2 := ibllD.Update(&degree)
  if err2 != nil {
    return c.RenderText("Error on update.")
  }

  return c.RenderJson(degree)
}

func (c Degrees) Delete(id int) revel.Result {
  // Check if the request is authorized by a JSON Web Token.
  tokenString := c.Request.Header.Get("Authorization")

  if err := c.checkJWTDegrees(tokenString); err != nil {
    c.Response.Status = http.StatusUnauthorized
    c.Response.ContentType = "application/json"
    // Change it to send a JSON with the message of the error, use the error 'err'.
    return c.RenderText("ERROR: You dont have authorization to perform this action.")
  }

  var ibllD bllinterfaces.IBLLDegree
  ibllD = &bllcontrollers.BLLDegree{}
  ibllD.InitBLLDegree()
  err2 := ibllD.Delete(id)
  if err2 != nil {
    return c.RenderText("Error on Delete.")
  }

  msg := "{ message: 'Eliminado.'}"
  return c.RenderJson(msg)
}

func (c Degrees) checkJWTDegrees(tokenString string) error {
  if tokenString == "" {
    return errors.New("Authorization header not found.")
  }

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return []byte(revel.Config.StringDefault("jwt.secret", "default")), nil
  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    // Esto tiene que ser diferente.
    if claims["admin"] == true {
      return nil
    } else {
      return errors.New("Wrong claim inside JWT.")
    }
  } else {
    return err
  }
}
