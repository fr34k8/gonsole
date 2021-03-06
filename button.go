package gonsole

import "github.com/nsf/termbox-go"

type Button struct {
	BasicControl

	// custom
	Text string
}

func NewButton(id string) *Button {
	button := &Button{}
	button.Init(id)
	button.SetBorder(LineSingleCorners)
	button.SetFocussable(true)
	return button
}

func (c *Button) Repaint() {
	if !c.Dirty() {
		return
	}
	c.BasicControl.Repaint()

	// content area
	DrawTextSimple(c.Text, c.ContentBox(), c.Foreground, c.Background)
}

func (btn *Button) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			btn.SubmitEvent(&Event{"clicked", btn, nil})
			return true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
