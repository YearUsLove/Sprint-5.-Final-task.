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
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес не должен быть меньше или равен 0")
	}
	if height <= 0 {
		return 0, errors.New("рост не должен быть меньше или равен 0")
	}
	if duration <= 0 {
		return 0, errors.New("длительность не должна быть меньша или равна 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	walkingSpentCalories := calories * walkingCaloriesCoefficient

	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес не должен быть меньше или равен 0")
	}
	if height <= 0 {
		return 0, errors.New("рост не должен быть меньше или равен 0")
	}
	if duration <= 0 {
		return 0, errors.New("длительность не должна быть меньша или равна 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	hours := duration.Hours()

	meanSpeed := distance / hours

	return meanSpeed

}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient

	distance := float64(steps) * stepLength

	distanceKm := distance / mInKm

	return distanceKm
}
