package error

import (
	"database/sql"
	"errors"
	"testing"
)

func TestNewCustomError(t *testing.T){ 
	err:=NewCustomError("ggg")
	t.Log(err.Error())
}

func TestNewCustomError1(t *testing.T){
	err:=NewCustomError1()
	// fmt.Errorf("3 layout: %w",sql.ErrConnDone)
	t.Log(errors.Is(err,sql.ErrNoRows))
	var se error
	ok:=errors.As(err,&se)
	if ok{
		t.Log(se.Error())
	}

}

func TestMyError(t *testing.T){
	e:=MyError{err: errors.New("前"),msg:"后"}
	t.Log(e.Error())
}