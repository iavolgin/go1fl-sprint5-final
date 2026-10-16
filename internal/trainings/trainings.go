package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	trainingData := strings.Split(datastring, ",")
	if len(trainingData) != 3 {
		err = fmt.Errorf("invalid input data")
		return
	}
	steps, err := strconv.Atoi(trainingData[0])
	if err != nil {
		return
	}
	if steps <= 0 {
		err = fmt.Errorf("steps value is invalid")
		return
	}

	duration, err := time.ParseDuration(trainingData[2])
	if err != nil {
		return
	}
	if duration <= 0 {
		err = fmt.Errorf("duration value is invalid")
	}

	t.Steps = steps
	t.TrainingType = trainingData[1]
	t.Duration = duration

	return
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		spentCalories float64
		err           error
	)
	switch t.TrainingType {
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("invalid training type")
	}

	ans := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories)
	return ans, nil
}
