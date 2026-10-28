package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные прогулки и персональные данные пользователя.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку "шаги,продолжительность".
// Сохраняет данные только после успешной проверки шагов и продолжительности.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("expected two comma-separated values")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("parse steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("parse duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("walk duration must be greater than zero")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo возвращает отчет о прогулке.
// При некорректных данных возвращает пустую строку и ошибку.
func (ds DaySteps) ActionInfo() (string, error) {
	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("calculate walking calories: %w", err)
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps, distanceKm, calories,
	)
	return info, nil
}
