package collector

type CookInfo struct {
	Id int64

	Name string  `db:"first_name"`
	Img string   `db:"img"`
	Material string `db:"material"`
	Brief string `db:"brief"`
    Content string `db:"content"`
	Key string `db:"key"`
}
type UpFile struct {
	Name string
	Base64Str string
}
