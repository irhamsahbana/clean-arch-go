package slice

import (
	"ca-boilerplate/domain"
)

var books = []domain.Book{
	{
		UUID:        "1",
		Title:       "The Lord of the Rings",
		Description: "The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien. The story began as a sequel to Tolkien's 1937 fantasy novel The Hobbit, but eventually developed into a much larger work. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling novels ever written, with over 150 million copies sold.",
		Author:      "J. R. R. Tolkien",
		Year:        1954,
	},
	{
		UUID:        "2",
		Title:       "Harry Potter and the Philosopher's Stone",
		Description: "Harry Potter and the Philosopher's Stone is a fantasy novel written by British author J. K. Rowling. The first novel in the Harry Potter series and Rowling's debut novel, it follows Harry Potter, a young wizard who discovers his magical heritage on his eleventh birthday, when he receives a letter of acceptance to Hogwarts School of Witchcraft and Wizardry. Harry makes close friends and a few enemies during his first year at the school, and with the help of his friends, Harry faces an attempted comeback by the dark wizard Lord Voldemort, who killed Harry's parents, but failed to kill Harry when he was just 15 months old.",
		Author:      "J. K. Rowling",
		Year:        1997,
	},
}

func NewBookSliceRepository() domain.BookRepositoryContract {
	return &bookSliceRepository{}
}

type bookSliceRepository struct {
}
