module go-musicfox

go 1.16

require (
	github.com/anhoder/bubbles v0.7.8 // indirect
	github.com/anhoder/bubbletea v0.12.8
	github.com/anhoder/netease-music v1.0.0 // indirect
	github.com/fogleman/ease v0.0.0-20170301025033-8da417bf1776 // indirect
	github.com/gookit/color v1.3.0 // indirect
	github.com/gookit/gcli/v2 v2.3.4
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/muesli/termenv v0.7.4
)

replace (
	github.com/anhoder/bubbletea => ../bubbletea
	github.com/anhoder/netease-music => ../netease-music
)
