package feature

type PartOfNameCharacters interface {
	Length() int
	Slice() []rune
	IsLastName() bool
}
