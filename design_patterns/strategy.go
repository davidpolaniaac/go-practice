package main

import "fmt"

type PasswordProtector struct {
	user string
	pass string
	hash Hash
}

type Hash interface {
	execute(p *PasswordProtector)
}

func NewPasswordProtector(user string, password string, hash Hash) *PasswordProtector {
	return &PasswordProtector{
		user: user,
		pass: password,
		hash: hash,
	}
}

func (p *PasswordProtector) Hash() {
	p.hash.execute(p)
}

func (p *PasswordProtector) SetHashAlgorithm(hash Hash) {
	p.hash = hash
}

type SHA struct{}

func (SHA) execute(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.pass)
}

type MD5 struct{}

func (MD5) execute(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.pass)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("test", "test", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
