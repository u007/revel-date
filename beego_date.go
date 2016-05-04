package beego_date

import (
  "github.com/astaxie/beego"
  "time"
  "fmt"
  "strings"
)

//parse dob-year, dob-month, dob-day 
func ParseDOB(controller *beego.Controller, form_prefix string) (time.Time, error) {
  dob_year, err := parseFieldInt(controller, form_prefix, "year")
  dob_month, err2 := parseFieldInt(controller, form_prefix, "month")
  dob_day, err3 := parseFieldInt(controller, form_prefix, "day")
  
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
  dob_string := fmt.Sprintf("%s-%s-%s", dob_year, dob_month,
    dob_day)
  dob, parse_err := time.Parse(layout, dob_string)
  if parse_err != nil {
    return time.Now(), parse_err
  }
  
  Debug("dob %s = %q", dob_string, dob)
  return dob, nil
}

func parseFieldInt(controller *beego.Controller, form_prefix string, name string) (int, error) {
  field_name := fmt.Sprintf("%s-%s", form_prefix, name)
  res, err := controller.GetInt(field_name)
  if (err != nil) {
    err = fmt.Errorf("Invalid field %s = %v", name, controller.GetString(field_name))
  }
  return res, err
}

const PREFIX = "[ BEEGO_DATE ] "

func Debug(format string, v... interface{}) {
	beego.Debug(fmt.Sprintf(PREFIX + format, v...))
}
func Warning(format string, v... interface{}) {
	beego.Warning(fmt.Sprintf(PREFIX + format, v...))
}
func Error(format string, v... interface{}) {
	beego.Error(fmt.Sprintf(PREFIX + format, v...))
}
