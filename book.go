package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "A sad adio",
		Author:        "Miki M",
		YearPublished: 1111,
	},
	Book{
		ID:            2,
		Title:         "To je to",
		Author:        "Raj Donovan",
		YearPublished: 2222,
	},
	Book{
		ID:            3,
		Title:         "Idi pa vidi",
		Author:        "Tkalec",
		YearPublished: 5555,
	},
	Book{
		ID:            4,
		Title:         "Kada",
		Author:        "Pa sada",
		YearPublished: 1989,
	},
	Book{
		ID:            5,
		Title:         "Yast",
		Author:        "Pa tak",
		YearPublished: 1996,
	},
	Book{
		ID:            6,
		Title:         "I",
		Author:        "Pa Te",
		YearPublished: 2000,
	},
	Book{
		ID:            7,
		Title:         "Dokle",
		Author:        "Do Keca",
		YearPublished: 1999,
	},
	Book{
		ID:            8,
		Title:         "Mislis",
		Author:        "Dakle Postojis",
		YearPublished: 2006,
	},
	Book{
		ID:            9,
		Title:         "Sve bi seke",
		Author:        "Ljubile Mornare",
		YearPublished: 1987,
	},
	Book{
		ID:            10,
		Title:         "Nije li",
		Author:        "Pa je",
		YearPublished: 3000,
	},
}
