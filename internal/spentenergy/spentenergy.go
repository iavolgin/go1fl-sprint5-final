package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	runningCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return runningCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	validateErr := validateParameters(steps, weight, height, duration)
	if validateErr != nil {
		return 0, validateErr
	}
	meanSpeed := MeanSpeed(steps, height, duration)

	return (weight * meanSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)

	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

func validateParameters(steps int, weight, height float64, duration time.Duration) error {
	errSuffix := "value must be positive"
	if steps <= 0 {
		return errors.New("steps " + errSuffix)
	}
	if weight <= 0 {
		return errors.New("weight " + errSuffix)
	}
	if height <= 0 {
		return errors.New("height " + errSuffix)
	}
	if duration <= 0 {
		return errors.New("duration " + errSuffix)
	}

	return nil
}
