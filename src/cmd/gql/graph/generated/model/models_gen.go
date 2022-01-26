// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	UID         string    `json:"uid"`
	Authors     []*string `json:"authors"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Cover       *int      `json:"cover"`
	Genres      []*string `json:"genres"`
	InList      int       `json:"inList"`
}

type CategorizedBooks struct {
	Genre string  `json:"genre"`
	Books []*Book `json:"books"`
}

type Depiction struct {
	Res *string `json:"res"`
}

type Genre struct {
	Name  string `json:"name"`
	Liked bool   `json:"liked"`
}

type Genres struct {
	Genres []*Genre `json:"genres"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Registration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	UID           string    `json:"UID"`
	Username      string    `json:"Username"`
	Password      string    `json:"Password"`
	Token         []*string `json:"Token"`
	LikedBooks    []*string `json:"LikedBooks"`
	WillingToRead []*string `json:"WillingToRead"`
	DislikedBooks []*string `json:"DislikedBooks"`
	LikedGenres   []*string `json:"LikedGenres"`
}

type UserID struct {
	ID string `json:"id"`
}

type UserMeta struct {
	UID      string `json:"UID"`
	Username string `json:"Username"`
}

type UsersBooks struct {
	Slices []*CategorizedBooks `json:"slices"`
}
