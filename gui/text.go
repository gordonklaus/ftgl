package gui

import (
	"code.google.com/p/gordon-go/ftgl"
	gl "github.com/chsc/gogl/gl21"
)

var font ftgl.Font

func initFont() {
	font = ftgl.NewTextureFont("/Library/Fonts/Times New Roman.ttf")
	font.SetFaceSize(18, 1)
}

type Text struct {
	*ViewBase
	text                string
	textColor           Color
	frameSize           float64
	frameColor          Color
	backgroundColor     Color
	Validate            func(*string) bool
	Accept, TextChanged func(string)
	Reject              func()
}

func NewText(text string) *Text {
	t := &Text{}
	t.ViewBase = NewView(t)
	t.textColor = Color{1, 1, 1, 1}
	t.backgroundColor = Color{0, 0, 0, 1}
	t.SetText(text)
	return t
}

func (t Text) Text() string { return t.text }
func (t *Text) SetText(text string) {
	t.text = text
	Resize(t, Pt(2*t.frameSize+font.Advance(t.text), 2*t.frameSize-font.Descender()+font.Ascender()))
	if t.TextChanged != nil {
		t.TextChanged(text)
	}
}

func (t *Text) SetTextColor(c Color) {
	t.textColor = c
	Repaint(t)
}

func (t *Text) SetBackgroundColor(c Color) {
	t.backgroundColor = c
	Repaint(t)
}

func (t *Text) SetFrameColor(c Color) {
	t.frameColor = c
	Repaint(t)
}

func (t *Text) SetFrameSize(size float64) {
	t.frameSize = size
	Resize(t, Pt(2*t.frameSize+font.Advance(t.text), 2*t.frameSize-font.Descender()+font.Ascender()))
}

func (t *Text) KeyPress(event KeyEvent) {
	if len(event.Text) > 0 {
		text := t.text + event.Text
		if t.Validate == nil || t.Validate(&text) {
			t.SetText(text)
		}
	}
	switch event.Key {
	case KeyBackspace:
		if len(t.text) > 0 {
			text := t.text[:len(t.text)-1]
			if t.Validate == nil || t.Validate(&text) {
				t.SetText(text)
			}
		}
	case KeyEnter:
		if t.Accept != nil {
			t.Accept(t.text)
		}
	case KeyEscape:
		if t.Reject != nil {
			t.Reject()
		}
	}
}

func (t *Text) Paint() {
	SetColor(t.backgroundColor)
	FillRect(Rect(t).Inset(t.frameSize))
	if t.frameSize > 0 {
		SetColor(t.frameColor)
		SetLineWidth(t.frameSize)
		DrawRect(Rect(t))
	}

	SetColor(t.textColor)
	gl.Translated(gl.Double(t.frameSize), gl.Double(t.frameSize-font.Descender()), 0)
	font.Render(t.text)
}
