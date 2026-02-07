package app

import (
	"cdrops/gcmz"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// DropFunc matches the signature expected by gcmz drop helpers.
type DropFunc func(layer int, msAdv int, paths []string, data *gcmz.GcmzDropsData) error

// Run parses command-line parameters and dispatches drop requests via the supplied function.
func Run(params []string, drop DropFunc) error {
	if drop == nil {
		return fmt.Errorf("drop function is nil")
	}

	data, err := gcmz.ReadGCMZDropsData()
	if err != nil {
		return err
	}

	for _, raw := range params {
		if raw == "" {
			continue
		}

		segments := strings.Split(raw, "*")
		if len(segments) < 2 {
			return fmt.Errorf("invalid parameter: %s", raw)
		}

		layer, err := strconv.Atoi(segments[0])
		if err != nil {
			return fmt.Errorf("invalid layer %s: %w", segments[0], err)
		}

		msAdv, err := strconv.Atoi(segments[1])
		if err != nil {
			return fmt.Errorf("invalid msAdv %s: %w", segments[1], err)
		}

		paths, err := normalizePaths(segments[2:])
		if err != nil {
			return err
		}

		if err := drop(layer, msAdv, paths, data); err != nil {
			return err
		}
	}

	return nil
}

func normalizePaths(items []string) ([]string, error) {
	result := make([]string, 0, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}

		cleaned := path.Clean(item)
		abs, err := filepath.Abs(cleaned)
		if err != nil {
			return nil, err
		}

		result = append(result, abs)
	}

	return result, nil
}
