package null

import "testing"

func TestSanity(t *testing.T) {
	nullInt := New(0)

	if !nullInt.Valid {
		t.Error("Expected nullInt to be valid")
	}

	if nullInt.Get() != 0 {
		t.Error("Expected nullInt to be 0")
	}

	invalidNullInt := NewInvalid[int]()

	if invalidNullInt.Valid {
		t.Error("Expected invalidNullInt to be invalid")
	}

	if invalidNullInt.Get() != 0 {
		t.Error("Expected invalidNullInt to be 0")
	}

	var invalidNullVarDeclare Null[int]

	if invalidNullVarDeclare.Valid {
		t.Error("Expected invalidNullVarDeclare to be invalid")
	}
}
