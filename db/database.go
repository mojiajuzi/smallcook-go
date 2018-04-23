package mojiajuzi

import "time"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
}

type Menu struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Time  int    `json:"time"`
}

type Inqredients struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

type Step struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	MenuID  int    `json:"menu_id"`
}

type MenuIngredients struct {
	ID            int    `json:"id"`
	MenuID        int    `json:"menu_id"`
	IngredientsID int    `json:"ingreditent_id"`
	Heigth        string `json:"height"`
}

type updatedAt struct {
	UpdatedAt time.Time `json:"updated_at"`
}

type createAt struct {
	CreatetedAt time.Time `json:"created_at"`
}
