package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var BG_COLOR = tcell.ColorDefault

var app = tview.NewApplication()


type HttpRequest struct {
    Method string
    Endpoint string
}

var requests = []HttpRequest {
    {"GET", "/api/test/hello"},
    {"GET", "/api/test/health"},
    {"POST", "/api/test/user"},
    {"GET", "/api/test/users"},
    {"GET", "/api/test/some/really/long/address"},
}

func main() {


    statusBar := tview.NewTextView()
    statusBar.SetBorder(true)
    statusBar.SetBackgroundColor(BG_COLOR)
    statusBar.SetTitle("Status")
    statusBar.SetTitleAlign(tview.AlignLeft)

    setStatus := func(msg string) {
        statusBar.SetText(msg)
    }

    reqMethods := []string{"GET", "POST", "PUT", "DELETE"}

    methodDropdown := tview.NewDropDown().SetOptions(reqMethods, nil).SetCurrentOption(0).SetFieldBackgroundColor(BG_COLOR).SetFieldTextColor(BG_COLOR)
    methodDropdown.SetTitle("Method [C-e]")
    methodDropdown.SetTitleAlign(tview.AlignLeft)
    methodDropdown.SetBackgroundColor(BG_COLOR)
    methodDropdown.SetBorder(true)
    methodDropdown.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

    baseUrls := []string{"http://localhost:8000"}

    baseUrlDropdown := tview.NewDropDown().SetOptions(baseUrls, nil).SetCurrentOption(0).SetFieldBackgroundColor(BG_COLOR).SetFieldTextColor(BG_COLOR)
    baseUrlDropdown.SetTitle("Base [C-b]")
    baseUrlDropdown.SetTitleAlign(tview.AlignLeft)
    baseUrlDropdown.SetBackgroundColor(BG_COLOR)
    baseUrlDropdown.SetBorder(true)
    baseUrlDropdown.SetListStyles(tcell.StyleDefault.Background(tcell.ColorGray), tcell.StyleDefault.Dim(true))

    urlInput := tview.NewInputField().SetFieldBackgroundColor(BG_COLOR).SetFieldTextColor(BG_COLOR)
    urlInput.SetBackgroundColor(BG_COLOR)
    urlInput.SetBorder(true)
    urlInput.SetTitle("URL [C-u]")
    urlInput.SetTitleAlign(tview.AlignLeft)

    urlInputCapture := func(event *tcell.EventKey) *tcell.EventKey {
        return event
    }

    urlInput.SetInputCapture(urlInputCapture)

    sendBtn := tview.NewButton("Send [⏎]")
    sendBtn.SetBorder(true)
    sendBtn.SetStyle(tcell.StyleDefault.Background(BG_COLOR))
    sendBtn.SetLabelColorActivated(tcell.ColorDarkBlue)
    sendBtn.SetActivatedStyle(tcell.StyleDefault.Background(BG_COLOR))

    sendRequest := func() {
       setStatus("req sent")
    }

    sendBtn.SetSelectedFunc(sendRequest)

    reqList := tview.NewList()

    for _, r := range requests {
        reqList.AddItem(fmt.Sprintf("%-4s", r.Method) + "  " + r.Endpoint, "", 0, nil)
    }

    reqList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
        selected := requests[index]

        // TODO: need to find method index 
        methodDropdown.SetCurrentOption(0)
        urlInput.SetText(selected.Endpoint)
    })

    reqList.ShowSecondaryText(false)
    reqList.SetBorder(true)
    reqList.SetBackgroundColor(BG_COLOR)
    reqList.SetTitle("Requests [C-r]")
    reqList.SetTitleAlign(tview.AlignLeft)
    reqList.SetBorderPadding(1,1,1,1)

    reqListInputCapture := func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyRune {
            idx := reqList.GetCurrentItem()
            switch event.Rune() {
            case 'j':
                reqList.SetCurrentItem(idx+1)
            case 'k':
                reqList.SetCurrentItem(idx-1)
            }
            return event
        }
        return event
    }

    reqList.SetInputCapture(reqListInputCapture)

    layout := tview.NewFlex().
        AddItem(reqList, 50, 1, true).
        AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
            AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
                AddItem(methodDropdown, 15, 1, false).
                AddItem(baseUrlDropdown, 25, 1, false).
                AddItem(urlInput, 0, 1, false).
                AddItem(sendBtn, 12, 1, false),
            3, 1, false).
            AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)").SetBackgroundColor(BG_COLOR), 0, 5, false).
            AddItem(statusBar, 3, 1, false), 
            0, 2, false,
        )

    
    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyCtrlC:
            return nil
        case tcell.KeyEnter:
            sendRequest()
            return nil
        case tcell.KeyCtrlR:
            app.SetFocus(reqList) 
        case tcell.KeyCtrlB:
            app.SetFocus(baseUrlDropdown)
        case tcell.KeyCtrlU:
            app.SetFocus(urlInput)
        case tcell.KeyCtrlE:
            app.SetFocus(methodDropdown)
        case tcell.KeyRune:
            if event.Rune() == 'q' && !urlInput.HasFocus() {
                app.Stop()
            }
        }
        return event
    })

    app.EnableMouse(true)
    app.SetRoot(layout, true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
