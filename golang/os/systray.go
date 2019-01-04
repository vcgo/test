package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"

	"github.com/go-vgo/robotgo"
)

var Stop bool

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	systray.SetIcon(icon.Data)
	systray.SetTooltip("Pretty awesome超级棒")

	mQuitOrig := systray.AddMenuItem("Quit", "Quit APP !")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

	startCh := make(chan int, 1)
	for {
		f11 := robotgo.AddEvent("k")
		if f11 == 0 {
			Stop = false
			fmt.Println("... key press F11")
			systray.SetTooltip("Going...F12 Stop")

			go func() {
				f12 := robotgo.AddEvent("s")
				if f12 == 0 {
					fmt.Println("... key press F12")
					systray.SetTooltip("F11 Start")
					startCh <- 2
				}
			}()

			go func() {
				for {
					if Stop {
						return
					}
					fmt.Println("... do", time.Now())
					time.Sleep(999 * time.Millisecond)
				}
			}()
		}
		ch := <-startCh
		Stop = true
		fmt.Println("... master", ch, Stop)
	}

}

func onExit() {
	fmt.Println("Quit APP !")
}
