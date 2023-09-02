package iteration

import "testing"

func TestRepeat(t *testing.T) {
  t.Run("Test Repeating 5 times", func(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"
    assertCorrectResponse(t, expected, repeated)
  })
  t.Run("Test Repeating 6 times", func (t *testing.T)  {
  repeated := Repeat("b", 6)
  expected := "bbbbbb"
  assertCorrectResponse(t, expected, repeated)
  })
}

func assertCorrectResponse (t testing.TB, expected, repeated string) {
  if repeated != expected {
    t.Errorf("expected %q, but got %q", expected, repeated)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 5)
  }
}
