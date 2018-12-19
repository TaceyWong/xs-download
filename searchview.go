package main

import "github.com/andlabs/ui"

func searchPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hsearchbox := ui.NewHorizontalBox()
	hsearchbox.SetPadded(true)
	searchEntry := ui.NewSearchEntry()
	hsearchbox.Append(searchEntry,true)
	searchBtn := ui.NewButton(Tr("开始搜索"))
	hsearchbox.Append(searchBtn,false)

	vbox.Append(hsearchbox, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	hbox.Append(ui.NewLabel(Tr("勾选搜索源:")), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)
	vbox.Append(hbox, false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	
	

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	// group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", ui.NewMultilineEntry(), true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	return vbox
}
