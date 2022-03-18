package tl

import "testing"

func TestEncrptyAndDecrpty(t *testing.T) {
	str := "事实胜于雄辩helloworldhelloworldhelloworldhelloworldhelloworld马克思主义政党"
	str1 := Encrypt(str)
	str2 := Decrypt(str1)

	if str != str2 {
		t.Errorf("Encrpty Or Decrpty error,%q,%q,%q", str, str1, str2)
	} else {
		t.Logf("%q,%q,%q", str, str1, str2)
	}
}

func BenchmarkEncrptyAndDecrpty(b *testing.B) {
	str := "事实胜于雄辩helloworldhelloworldhelloworldhelloworldhelloworld马克思主义政党"
	for i := 0; i < b.N; i++ {
		str1 := Encrypt(str)
		str2 := Decrypt(str1)
		if str != str2 {
			b.Errorf("Encrpty Or Decrpty error")
		}
	}
}
