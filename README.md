# EDitor

> An _application_ contains a list of _windows_ (pages) each with a _buffer_.

No `context`. Only global is `log`.

- bindings: all keybindings
- buffer: femto.Buffer with helpers
- config: configuration
- log: logger
- main: entry-point (init, app, ui...)
- open: open directory or file
- window: header, buffer, footer

## Dependencies

- [femto](github.com/pgavlin/femto) / [tview](github.com/rivo/tview) / [tcell](github.com/gdamore/tcell/v2)
- [xdg](github.com/adrg/xdg) + [kdl](github.com/sblinch/kdl-go)
- [semver](github.com/Masterminds/semver/v3)
- [pp](github.com/k0kubun/pp/v3)
