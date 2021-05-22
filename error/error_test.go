package error

import (
	"database/sql"
  gerrors "github.com/pkg/errors"
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


// Wrap 方法用来包装底层错误，增加上下文文本信息并添加调用栈。一般用于包装对第三方代码(标准库或者第三方库)的调用
// innnerError 内部错误
func innerError()error{
   return gerrors.Wrap(sql.ErrNoRows,"inner")
}


// WithMessage 方法仅添加上下文文本信息，不附加调用栈。如果确定错误已被Wrap过或不关心调用栈，可以使用次方法。
// outError 外部错误
func outError()error{
   return gerrors.WithMessage(innerError(),"outError failed")
}