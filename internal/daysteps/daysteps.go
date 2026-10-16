package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	stepsData := strings.Split(datastring, ",")
	if len(stepsData) != 2 {
		err = fmt.Errorf("invalid input data")
		return
	}
	steps, err := strconv.Atoi(stepsData[0])
	if err != nil {
		return
	}
	if steps <= 0 {
		err = fmt.Errorf("steps value is invalid")
		return
	}

	duration, err := time.ParseDuration(stepsData[1])
	if err != nil {
		return
	}
	if duration <= 0 {
		err = fmt.Errorf("duration value is invalid")
	}

	ds.Steps = steps
	ds.Duration = duration

	return
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	ans := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories)
	return ans, nil
}
