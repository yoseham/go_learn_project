package maps

import (
	"testing"
)

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s' given '%s'", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("expected to get an error.")
	}
	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func TestDictionary_Search(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknow word", func(t *testing.T) {
		_, err := dic.Search("hello")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrorNotFound)
	})
}

func TestDictionary_Add(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		err := dic.Add("test", "this is just a test")
		assertError(t, err, nil)
		assertDefinition(t, dic, "test", "this is just a test")
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		err := dic.Add(word, "new test")
		assertError(t, err, ErrorWordExist)
		assertDefinition(t, dic, word, definition)
	})

}

func TestDictionary_Update(t *testing.T) {
	t.Run("exist word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dic.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dic, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{}

		err := dic.Update(word, definition)
		assertError(t, err, ErrorWordNotExist)
	})

}

func TestDictionary_Delete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dic := Dictionary{word: definition}
	dic.Delete(word)
	_, err := dic.Search(word)
	if err == nil {
		t.Fatal("expected to get an error.")
	}
	assertError(t, err, ErrorNotFound)

}
