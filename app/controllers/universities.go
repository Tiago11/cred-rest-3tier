package controllers

import (
  "github.com/revel/revel"
  "credits_be/models"
  "encoding/json"
  "net/http"
  jwt "github.com/dgrijalva/jwt-go"
  "fmt"
  "errors"
  "credits_be/businesslogiclayer/bllinterfaces"
  "credits_be/businesslogiclayer/bllcontrollers"
)

type Universities struct {
  Application
}

func (c Universities) Index() revel.Result {
  var ibllU bllinterfaces.IBLLUniversity
  ibllU = &bllcontrollers.BLLUniversity{}
  ibllU.InitBLLUniversity()
  universities := ibllU.GetAll()

  return c.RenderJson(universities)
}

func (c Universities) loadUniversityById(id int) *models.University {
  var ibllU bllinterfaces.IBLLUniversity
  ibllU = &bllcontrollers.BLLUniversity{}
  ibllU.InitBLLUniversity()
  u := ibllU.Get(id)
  if u == nil {
    return nil
  }

  return u
}

func (c Universities) loadUniversityWithDegrees(id int) *models.University {
  var ibllU bllinterfaces.IBLLUniversity
  ibllU = &bllcontrollers.BLLUniversity{}
  ibllU.InitBLLUniversity()
  u := ibllU.GetWithDegrees(id)
  if u == nil {
    return nil
  }

  return u
}

func (c Universities) Show(id int) revel.Result {
  university := c.loadUniversityWithDegrees(id)
  if university == nil {
    return c.NotFound("University %d does not exist", id)
  }
  return c.RenderJson(university)
}

func (c Universities) parseUniversity() (models.University, error) {
  university := models.University{}
  err := json.NewDecoder(c.Request.Body).Decode(&university)
  return university, err
}

func (c Universities) Create() revel.Result {

  // Check if the request is authorized by a JSON Web Token.
  tokenString := c.Request.Header.Get("Authorization")

  if err := c.checkJWTUniversity(tokenString); err != nil {
    c.Response.Status = http.StatusUnauthorized
    c.Response.ContentType = "application/json"
    // Change it to send a JSON with the message of the error, use the error 'err'.
    return c.RenderText("ERROR: You dont have authorization to perform this action.")
  }

  if university, err := c.parseUniversity(); err != nil {
    return c.RenderText("Unable to parse a University from JSON.")
  } else {
    // Validate the model
    university.Validate(c.Validation)
    if c.Validation.HasErrors() {
      // Do something better here!
      return c.RenderText("You have an error in your University.")
    } else {
      var ibllU bllinterfaces.IBLLUniversity
      ibllU = &bllcontrollers.BLLUniversity{}
      ibllU.InitBLLUniversity()
      if err := ibllU.Add(&university); err != nil {
        return c.RenderText("Error inserting record.")
      } else {
        c.Response.Status = http.StatusCreated
        c.Response.ContentType = "application/json"
        return c.RenderJson(university)
      }
    }
  }
}

func (c Universities) Update(id int) revel.Result {
  // Check if the request is authorized by a JSON Web Token.
  tokenString := c.Request.Header.Get("Authorization")

  if err := c.checkJWTUniversity(tokenString); err != nil {
    c.Response.Status = http.StatusUnauthorized
    c.Response.ContentType = "application/json"
    // Change it to send a JSON with the message of the error, use the error 'err'.
    return c.RenderText("ERROR: You dont have authorization to perform this action.")
  }

  university, err := c.parseUniversity()
  if err != nil {
    return c.RenderText("Unable to parse a University from the JSON.")
  }
  // Ensure the Id is set.
  university.UniversityId = id

  var ibllU bllinterfaces.IBLLUniversity
  ibllU = &bllcontrollers.BLLUniversity{}
  ibllU.InitBLLUniversity()
  err2 := ibllU.Update(&university)
  if err2 != nil {
    return c.RenderText("Error in Update.")
  }

  return c.RenderJson(university)
}

func (c Universities) Delete(id int) revel.Result {
  // Check if the request is authorized by a JSON Web Token.
  tokenString := c.Request.Header.Get("Authorization")

  if err := c.checkJWTUniversity(tokenString); err != nil {
    c.Response.Status = http.StatusUnauthorized
    c.Response.ContentType = "application/json"
    // Change it to send a JSON with the message of the error, use the error 'err'.
    return c.RenderText("ERROR: You dont have authorization to perform this action.")
  }

  //dalU := &dalcontrollers.DALUniversity{}
  //err2 := dalU.Delete(id)
  var ibllU bllinterfaces.IBLLUniversity
  ibllU = &bllcontrollers.BLLUniversity{}
  ibllU.InitBLLUniversity()
  err2 := ibllU.Delete(id)
  if err2 != nil {
    return c.RenderText("Failed to remove the University.")
  }

  msg := "{ message: 'Eliminado.'}"
  return c.RenderJson(msg)
}

func (c Universities) checkJWTUniversity(tokenString string) error {
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
