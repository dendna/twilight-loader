package twilightloader

import (
	"fmt"
	"html/template"
	"os"
	"time"

	tw "github.com/dendna/twilight"
)

const insertTemplate = `
insert into lms.sun_schedules (id, area_id, "day", "month", calculated_sunrise_at, calculated_sunset_at, sunrise_at, sunset_at, on_at, off_at, tz, morning_twilight_started_at, evening_twilight_end_at)
values({{ .ID }}, 1, {{ .Day}}, {{ .Month }}, '{{ .Sunrise}}', '{{ .Sunset }}', '{{ .Sunrise }}', '{{ .Sunset }}', '{{ .Sunset }}', '{{ .Sunrise }}', '{{ .TimeZone }}', '{{ .MorningTwilight }}', '{{ .EveningTwilight }}');`

type SunSchedule struct {
	ID, Day, Month                   int
	Sunrise, Sunset                  string
	TimeZone                         string
	MorningTwilight, EveningTwilight string
}

type Config struct {
	Year                int         `json:"year"`
	TimezoneName        string      `json:"timezone_name"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	MorningTwilightType tw.DuskType `json:"morning_twilight_type"`
	SunriseType         tw.DuskType `json:"sunrise_type"`
	EveningTwilightType tw.DuskType `json:"evening_twilight_type"`
	SunsetType          tw.DuskType `json:"sunset_type"`
}

func Generate(config Config) error {
	//fmt.Println("Config: ", config)
	fmt.Printf("Config: %+v\n", config)

	loc, err := time.LoadLocation(config.TimezoneName)
	if err != nil {
		return err
	}
	fmt.Println("Location:", loc)

	days := time.Date(config.Year, time.December, 31, 0, 0, 0, 0, time.Local).YearDay()
	fmt.Println("Days:", days)

	insert, err := template.New("insert").Parse(insertTemplate)
	if err != nil {
		return err
	}

	curDate := time.Date(config.Year-1, time.December, 31, 0, 0, 0, 0, time.Local)
	for i := 1; i <= days; i++ {
		curDate = curDate.AddDate(0, 0, 1)
		morningTwilight, err := tw.CalcRise(config.Latitude, config.Longitude, config.MorningTwilightType, curDate.Year(), int(curDate.Month()), curDate.Day())
		if err != nil {
			return err
		}

		sunrise, err := tw.CalcRise(config.Latitude, config.Longitude, config.SunriseType, curDate.Year(), int(curDate.Month()), curDate.Day())
		if err != nil {
			return err
		}

		sunset, err := tw.CalcSet(config.Latitude, config.Longitude, config.SunsetType, curDate.Year(), int(curDate.Month()), curDate.Day())
		if err != nil {
			return err
		}

		eveningTwilight, err := tw.CalcSet(config.Latitude, config.Longitude, config.EveningTwilightType, curDate.Year(), int(curDate.Month()), curDate.Day())
		if err != nil {
			return err
		}

		sunSchedule := SunSchedule{
			ID:              i,
			Day:             curDate.Day(),
			Month:           int(curDate.Month()),
			Sunrise:         sunrise.In(loc).Format("15:04:05"),
			Sunset:          sunset.In(loc).Format("15:04:05"),
			TimeZone:        config.TimezoneName,
			MorningTwilight: morningTwilight.In(loc).Format("15:04:05"),
			EveningTwilight: eveningTwilight.In(loc).Format("15:04:05"),
		}

		err = insert.Execute(os.Stdout, sunSchedule)
		if err != nil {
			return err
		}
	}

	//fmt.Println(tw.Calc(config.Latitude, config.Longitude, config.MorningTwilightType, config.Year, 1, 1))
	//fmt.Println(tw.Calc(config.Latitude, config.Longitude, config.SunriseType, config.Year, 1, 1))
	//fmt.Println(tw.Calc(config.Latitude, config.Longitude, config.SunsetType, config.Year, 1, 1))
	//fmt.Println(tw.Calc(config.Latitude, config.Longitude, config.EveningTwilightType, config.Year, 1, 1))

	return nil
}
