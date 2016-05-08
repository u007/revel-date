package revel_date

import (
  "github.com/revel/revel"
  "time"
  "fmt"
  "strings"
  "strconv"
)

//parse date-year, date-month, date-day
func ParseDate(c *revel.Controller, form_prefix string) (time.Time, error) {
  dob_year, err := parseFieldInt(c, form_prefix, "year")
  dob_month, err2 := parseFieldInt(c, form_prefix, "month")
  dob_day, err3 := parseFieldInt(c, form_prefix, "day")

  all_err := []string{}
  if (err != nil) {
    all_err = append(all_err, err.Error())
  }
  if (err2 != nil) {
    all_err = append(all_err, err2.Error())
  }
  if (err3 != nil) {
    all_err = append(all_err, err3.Error())
  }

  if len(all_err) > 0 {
    return time.Now(), fmt.Errorf(strings.Join(all_err, ", "))
  }

  layout := "2006-01-02"
  dob_string := fmt.Sprintf("%04d-%02d-%02d", dob_year, dob_month, dob_day)
  // Debug("dob string %s", dob_string)
  dob, parse_err := time.Parse(layout, dob_string)
  if parse_err != nil {
    return time.Time{}, parse_err
  }

  Debug("dob %s = %q", dob_string, dob)
  return dob, nil
}

func parseFieldInt(c *revel.Controller, form_prefix string, name string) (int, error) {
  field_name := fmt.Sprintf("%s%s", form_prefix, name)
  i, err := strconv.Atoi(c.Params.Get(field_name))
  if err != nil {
    return 0, fmt.Errorf("Invalid field %s = %v", field_name, c.Params.Get(field_name))
  }
  // res, err := controller.Params(field_name)
  // if (err != nil) {
  //   Error("field %s=%v", field_name, controller.GetString(field_name))
  //   err = fmt.Errorf("Invalid field %s = %v", name, controller.GetString(field_name))
  // } else {
  //   // Debug("field %s=%d", name, res)
  // }

  return i, nil
}

const PREFIX = "[ REVEL_DATE ] "

func Debug(format string, v... interface{}) {
	revel.INFO.Printf(PREFIX + format, v...)
}
func Warning(format string, v... interface{}) {
  revel.WARN.Printf(PREFIX + format, v...)
}
func Error(format string, v... interface{}) {
	revel.ERROR.Printf(PREFIX + format, v...)
}
