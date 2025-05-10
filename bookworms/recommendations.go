package main

type bookRecommendations map[Book]bookCollection

type bookCollection map[Book]struct{}

func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	// initialize recommendations
	recommendations := make(bookRecommendations)

	// fill recommendations
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			// get the other books of the bookworm
			otherBooks := listOtherBooksFromBookworm(i, bookworm)
			// store it in recommendations
			registerBookRecommendations(recommendations, book, otherBooks)
		}
	}

	// recommend books to each bookworm
	bookwormsWithRecommendations := getBookwormsWithRecommendations(bookworms, recommendations)

	// return bookworms with a list of recommended books
	return bookwormsWithRecommendations
}

func getBookwormsWithRecommendations(bookworms []Bookworm, recommendations bookRecommendations) []Bookworm {
	bookwormsWithRecommendations := make([]Bookworm, 0, len(bookworms))

	for _, bookworm := range bookworms {
		bookwormShelf := make(map[Book]bool)
		for _, book := range bookworm.Books {
			bookwormShelf[book] = true
		}

		bc := make(bookCollection)
		for _, book := range bookworm.Books {
			collectionOfRecommendedBooks := recommendations[book]
			for recommendedBook := range collectionOfRecommendedBooks {
				if bookwormShelf[recommendedBook] {
					continue
				}

				bc[recommendedBook] = struct{}{}
			}
		}

		bcslice := make([]Book, 0, len(bc))
		for book := range bc {
			bcslice = append(bcslice, book)
		}

		bcslice = sortBooks(bcslice)
		bookwormsWithRecommendations = append(bookwormsWithRecommendations, Bookworm{
			Name:  bookworm.Name,
			Books: bcslice,
		})

	}

	return bookwormsWithRecommendations
}

func registerBookRecommendations(recommendations bookRecommendations, book Book, otherBooks []Book) {
	for _, otherBook := range otherBooks {
		collection, ok := recommendations[book]
		if !ok {
			collection = newCollection()
			recommendations[book] = collection
		}

		collection[otherBook] = struct{}{}
	}
}

func listOtherBooksFromBookworm(index int, bookworm Bookworm) []Book {
	var otherBooks []Book

	for i, book := range bookworm.Books {
		if i != index {
			otherBooks = append(otherBooks, book)
		}
	}

	return otherBooks
}
