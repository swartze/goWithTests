package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T) {
  got := Hello("Lee", "")
  want := "Hello, Lee."
  assertCorrectMessage(t, got, want)
})
t.Run("say 'Hello, World.' when an empty string is supplied", func (t *testing.T)  {
  got := Hello("", "")
  want := "Hello, World."
  assertCorrectMessage(t, got, want)
})
t.Run("In Spanish", func(t *testing.T) {
  got := Hello("Elodie", "Spanish")
  want := "Hola, Elodie."
  assertCorrectMessage(t, got, want)
})
t.Run("In French", func(t *testing.T) {
  got := Hello("Jean", "French")
  want := "Bonjour, Jean."
  assertCorrectMessage(t, got, want)
})
}

func TestSumAll(t *testing.T)  {
  got := SumAll([]int{1,2}, []int{0,9})
  want :[]int{3,9}

  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func assertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}
