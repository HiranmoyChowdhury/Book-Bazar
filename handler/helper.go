package handler

import "learnProject/First-Project-With-Go/model"

func IsValid(value *model.Book) bool {
	temp := value.Name
	if temp == "" {
		return false
	}
	if len(value.AuthorList) == 0 {
		return false
	}
	for _, name := range value.AuthorList {
		if name == "" {
			return false
		}
	}
	temp = value.PublishDate
	if temp == "" {
		return false
	}
	temp = value.ISBN
	if temp == "" {
		return false
	}
	return true

}
