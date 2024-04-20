package cities

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name string
	Lat  float64
	Long float64
}

func LoadCitiesFromFile(logger *slog.Logger) ([]City, map[string]City, error) {
	// Load cities file
	file, err := os.Open("./cities15000.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	return ReadCities(scanner)
}

func ReadCities(scanner *bufio.Scanner) ([]City, map[string]City, error) {
	var line string
	cities := make([]City, 0, 25_000)
	coordinates := make(map[string]City, 25_000)

	for scanner.Scan() {
		line = scanner.Text()
		split := strings.Split(line, "\t")
		if len(split) < 6 {
			continue
		}

		lat, err := strconv.ParseFloat(split[4], 64)
		if err != nil {
			continue
		}
		long, err := strconv.ParseFloat(split[5], 64)
		if err != nil {
			continue
		}
		city := City{
			Name: strings.ToLower(split[1]),
			Lat:  lat,
			Long: long,
		}
		cities = append(cities, city)
		coordinates[strings.ToLower(split[1])] = city
	}

	return cities, coordinates, nil
}
