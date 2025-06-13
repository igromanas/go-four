package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrParsingData     = errors.New("parsing error")
	ErrConvToInt       = errors.New("convert to int error")
	ErrNotPositiveVal  = errors.New("not positive value error")
	ErrUnknownActivity = errors.New("неизвестный тип тренировки")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	args := strings.Split(data, ",")
	if len(args) != 3 {
		return 0, "", 0, fmt.Errorf("%w at parseTraining: 3 args required", ErrParsingData)
	}

	steps, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("%w at parseTraining: steps undefined", ErrConvToInt)
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("%w at parseTraining: not positive steps", ErrNotPositiveVal)
	}

	activity := args[1]

	duration, err := time.ParseDuration(args[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("%w at parseTraining: duration undefined", ErrParsingData)
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("%w at parseTraining: not positive duration", ErrNotPositiveVal)
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("%w at RunningSpentCalories: not positive steps", ErrNotPositiveVal)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("%w at RunningSpentCalories: not positive weight", ErrNotPositiveVal)
	}

	if height <= 0 {
		return 0, fmt.Errorf("%w at RunningSpentCalories: not positive height", ErrNotPositiveVal)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("%w at RunningSpentCalories: not positive duration", ErrNotPositiveVal)
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if steps <= 0 {
		return "", err
	}

	var cal float64
	switch activity {
	case "Бег":
		cal, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			//fmt.Println(err)
			return "", err
		}
	case "Ходьба":
		cal, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			//fmt.Println(err)
			return "", err
		}
	default:
		return "", ErrUnknownActivity
	}

	dist := distance(steps, height)
	ms := meanSpeed(steps, height, duration)

	layout := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
	return fmt.Sprintf(layout, activity, duration.Hours(), dist, ms, cal), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("%w at WalkingSpentCalories: not positive steps", ErrNotPositiveVal)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("%w at WalkingSpentCalories: not positive weight", ErrNotPositiveVal)
	}

	if height <= 0 {
		return 0, fmt.Errorf("%w at WalkingSpentCalories: not positive height", ErrNotPositiveVal)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("%w at WalkingSpentCalories: not positive duration", ErrNotPositiveVal)
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes() * walkingCaloriesCoefficient) / minInH, nil
	// rsc, err := RunningSpentCalories(steps, weight, height, duration)
	// return rsc * walkingCaloriesCoefficient, err
}
