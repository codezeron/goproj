package types

import "time"

type UserStore interface{
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface{
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProducts()	([]*Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type OrderStore interface{
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type Order struct{
	ID 						int 				`json:"id"`
	UserID 				int 				`json:"userID"`
	Total 				float64 		`json:"total"`
	Status 				string 			`json:"status"`
	Address     	string      `json:"address"`
	CreatedAt 		time.Time 	`json:"created_at"`
}

type OrderItem struct{
	ID 						int 				`json:"id"`
	OrderID 			int 				`json:"orderID"`
	ProductID 		int 				`json:"productID"`
	Quantity 			int 				`json:"quantity"`
	Price     		float64     `json:"price"`
	CreatedAt 		time.Time 	`json:"created_at"`
}
type Product struct{
	ID 						int 				`json:"id"`
	Name 					string 			`json:"name"`
	Description 	string 			`json:"description"`
	Image 				string 			`json:"image"`
	Price     		float64     `json:"price"`
	// note that this isn't the best way to handle quantity
	// because it's not atomic (in ACID), but it's good enough for this example
	Quantity  		int    			`json:"quantity"` 
	CreatedAt 		time.Time 	`json:"created_at"`
}
type User struct{
	ID 				int 			`json:"id"`
	FirstName string 		`json:"firstname"`
	LastName 	string 		`json:"lastname"`
	Email 		string 		`json:"email"`
	Password  string    `json:"-"` //não queremos que a senha seja serializada
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUser struct{
	Email 		string 	`json:"email" validate:"required,email"`	
	Password 	string 	`json:"password" validate:"required,min=6,max=100"`
	FirstName string 	`json:"firstname" validate:"required"`
	LastName 	string 	`json:"lastname" validate:"required"`
}

type LoginUser struct{
	Email 		string 	`json:"email" validate:"required,email"`	
	Password 	string 	`json:"password" validate:"required"`
}

type CartItem struct{
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`

}
type CartCheckoutPayload struct{
	Items []CartItem `json:"items" validate:"required"`
}