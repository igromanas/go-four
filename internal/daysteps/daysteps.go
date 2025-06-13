package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

var (
	ErrParsingData    = errors.New("parsing error")
	ErrConvToInt      = errors.New("convert to int error")
	ErrNotPositiveVal = errors.New("not positive value error")
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	args := strings.Split(data, ",")
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("%w at parsePackage: 2 args required", ErrParsingData)
	}

	steps, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w at parsePackage: steps undefined", ErrConvToInt)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("%w at parsePackage: not positive steps", ErrNotPositiveVal)
	}

	duration, err := time.ParseDuration(args[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w at parsePackage: duration undefined", ErrParsingData)
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("%w at parsePackage: not positive duration", ErrNotPositiveVal)
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		// fmt.Println(err)
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		log.Println(err)
		return ""
	}

	dist := float64(steps) * stepLength / mInKm
	cal, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		// fmt.Println(err)
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, cal)
}
