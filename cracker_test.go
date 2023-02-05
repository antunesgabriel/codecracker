package codecracker

import "testing"

func TestEncrypt(t *testing.T) {
	t.Run("it should return '!' when receiving 'a' as a parameter", func(t *testing.T) {
		expected := `!`
		received := Encrypt("a")

		if received != expected {
			t.Errorf("[FAIL] - expected %s but received %s", expected, received)
		}
	})

	t.Run(`it should return '!)"(' when receiving 'abcd' as a parameter`, func(t *testing.T) {
		expected := `!)"(`
		received := Encrypt("abcd")

		if received != expected {
			t.Errorf("[FAIL] - expected %s but received %s", expected, received)
		}
	})

	t.Run(`it should return 'k!h"d' when receiving 'vasco' as a parameter`, func(t *testing.T) {
		expected := `k!h"d`
		received := Encrypt("vasco")

		if received != expected {
			t.Errorf("[FAIL] - expected %s but received %s", expected, received)
		}
	})
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Encrypt("eusereioqueserei")
    }
}

func TestDecrypt(t *testing.T) {
	t.Run(`it should return 'vasco' when receiving 'k!h"d' as a parameter`, func(t *testing.T) {
		expected := `vasco`
		received := Decrypt(`k!h"d`)

		if received != expected {
			t.Errorf("[FAIL] - expected %s but received %s", expected, received)
		}
	})

	t.Run(`it should return 'brazil' when receiving ')g!o>a' as a parameter`, func(t *testing.T) {
		expected := `brazil`
		received := Decrypt(`)g!o>a`)

		if received != expected {
			t.Errorf("[FAIL] - expected %s but received %s", expected, received)
		}
	})
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Decrypt(")g!o>a)g!o>a)g!o>a")
    }
}