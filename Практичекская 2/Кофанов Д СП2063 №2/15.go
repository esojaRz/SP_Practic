package main

import "fmt"

type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genres []string
}

func getBest(movies []Movie) Movie {
	best := movies[0]
	for _, m := range movies {
		if m.Rating > best.Rating {
			best = m
		}
	}
	return best
}

func findByGenre(movies []Movie, genre string) []Movie {
	var found []Movie
	for _, m := range movies {
		for _, g := range m.Genres {
			if g == genre {
				found = append(found, m)
			}
		}
	}
	return found
}

func main() {
	movies := []Movie{
		{"Интерстеллар", 2014, 8.6, []string{"Фантастика", "Драма"}},
		{"Начало", 2010, 8.8, []string{"Фантастика", "Боевик"}},
		{"Побег из Шоушенка", 1994, 9.3, []string{"Драма"}},
		{"Зеленая миля", 1999, 8.6, []string{"Драма", "Криминал"}},
		{"Матрица", 1999, 8.7, []string{"Фантастика", "Боевик"}},
	}

	fmt.Println("Хороший фильм:", getBest(movies).Title)
	fmt.Println("\nДрамы:")
	for _, m := range findByGenre(movies, "Драма") {
		fmt.Println("-", m.Title)
	}
}
