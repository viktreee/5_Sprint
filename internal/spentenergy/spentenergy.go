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

// WalkingSpentCalories рассчитывает калории при ходьбе.
// При неположительных шагах, весе, росте или продолжительности возвращает ошибку.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("walking duration must be greater than zero")
	}

	speed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := weight * speed * durationMinutes / minInH
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories рассчитывает калории при беге.
// При неположительных шагах, весе, росте или продолжительности возвращает ошибку.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("running duration must be greater than zero")
	}

	speed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := weight * speed * durationMinutes / minInH
	return calories, nil
}

// MeanSpeed возвращает среднюю скорость в километрах в час.
// При отрицательных шагах или неположительной продолжительности возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	distanceKm := Distance(steps, height)
	return distanceKm / duration.Hours()
}

// Distance рассчитывает дистанцию в километрах по количеству шагов и росту в метрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}
