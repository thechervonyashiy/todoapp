package models

type TodoList struct {
	Id          int
	Title       string
	Description string
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
