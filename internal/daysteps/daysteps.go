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

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return errors.New("длина слайса не равна 2")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("ошибка преобразования количества шагов")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным числом")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("ошибка преобразования продолжительности")
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка вычисления калорий: %w", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return result, nil

}
