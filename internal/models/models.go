package models

type CreateActor struct {
	Name      string `json:"name" db:"name"`
	BirthDate string `json:"birth_date" db:"birth_date"`
	Gender    string `json:"gender" db:"gender"`
}
type Actor struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	BirthDate string `json:"birth_date" db:"birth_date"`
	Gender    string `json:"gender" db:"gender"`
}
type UpdateActor struct {
	ID        int     `json:"id" db:"id"`
	Name      *string `json:"name,omitempty" db:"name"`
	BirthDate *string `json:"birth_date,omitempty" db:"birth_date"`
	Gender    *string `json:"gender,omitempty" db:"gender"`
}
type CreateMovie struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ReleaseDate int    `json:"release_year" db:"release_Date"`
	ActorID     int    `json:"actor_id" db:"actor_id"`
}
type Movie struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ReleaseDate int    `json:"release_date" db:"release_date"`
	ActorID     int    `json:"actor_id" db:"actor_id"`
}
type UpdateMovie struct {
	ID          int     `json:"id" db:"id"`
	Title       *string `json:"title,omitempty" db:"title"`
	Description *string `json:"description,omitempty" db:"description"`
	ReleaseDate *int    `json:"release_date,omitempty" db:"release_date"`
	ActorID     *int    `json:"actor_id,omitempty" db:"actor_id"`
}
type Register struct {
	ID        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"-" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Role      string `json:"role" db:"role"`
}
type User struct {
	ID        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"-" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Role      string `json:"role" db:"role"`
}
