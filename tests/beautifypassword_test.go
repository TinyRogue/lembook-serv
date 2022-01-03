package tests

import (
	hash2 "github.com/TinyRogue/lembook-serv/src/pkg/hash"
	"math"
	"testing"
)

func TestBeautifyPasswordSpeed(t *testing.T) {
	bench := testing.Benchmark(BenchmarkBeautifyPerformance)
	halfOfSec := int64(math.Pow10(9) / 4)
	if bench.NsPerOp() > halfOfSec {
		t.Errorf("Took more than 250ms.")
	}
}

func BenchmarkBeautifyPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := hash2.BeautifyPassword("pa$$word", nil)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func TestBeautifyPasswordCorrectness(t *testing.T) {
	probes := 20
	hashes := make([]string, probes)
	for i := 0; i < probes; i++ {
		hashedPassword, err := hash2.BeautifyPassword("pa$$word", nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		hashes[i] = hashedPassword
	}

	for i := 0; i < probes; i++ {
		for j := 0; j < probes; j++ {
			if i == j {
				continue
			}
			if hashes[i] == hashes[j] {
				t.Errorf("Got %v and %v. Hashes with salt should never be the same.", hashes[i], hashes[j])
			}
		}
	}
}

func TestCompare(t *testing.T) {
	password := "The3Greate$tPasswords3v3r"
	hashedPassword, err := hash2.BeautifyPassword(password, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	match, err := hash2.Compare(password, hashedPassword)
	if err != nil {
		t.Errorf(err.Error())
	} else if !match {
		t.Errorf("Password shoud be a great Tinder's match")
	}
}

func TestDecodeHash(t *testing.T) {
	toDecode := []struct {
		hash string
		ans  bool
		desc string
	}{
		{hash: "$argon2id$v=19$m=65536,t=2,p=4$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG", ans: true, desc: "correct"},
		{hash: "$argon2i$v=19$m=65536,t=2,p=4$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG", ans: false, desc: "wrong alg"},
		{hash: "$argon2id$v=18$m=65536,t=2,p=4$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG", ans: false, desc: "wrong version"},
	}
	for _, tt := range toDecode {
		t.Run(tt.hash, func(t *testing.T) {
			_, _, _, err := hash2.DecodeHash(tt.hash)
			if !tt.ans && err == nil {
				t.Errorf(tt.desc)
			}
		})
	}
}
