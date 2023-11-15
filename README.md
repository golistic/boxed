boxed - Draw boxes around data
==============================

Copyright (c) 2021, 2023, Geert JM Vanderkelen

Draw boxes around your data and display tables/listings on your favorite TTY.

Originally part of a (private) project conceived in the summer of 2021.

## Features

* Define headers and rows displaying the typical tables as seen in popular RDBMS client tools.
* Define your own boxing styles if stock owns are not enough.

## Examples

### Basic-style

```go
package boxed_test

import (
	"github.com/golistic/boxed"
	"github.com/golistic/boxed/style"
)

func ExampleBasic() {
	box := boxed.New()

	_ = box.AddHeader("id", "user", "host", "locked")
	_ = box.Append(
		boxed.NewRow("1", "alice", "%", false),
		boxed.NewRow("2", "bob", "127.0.0.1", false),
		boxed.NewRow("3", "trudy", "%", true),
	)

	_ = box.Render()
	// Output:
	// ┌────┬───────┬───────────┬────────┐
	// │ id │ user  │ host      │ locked │
	// ├────┼───────┼───────────┼────────┤
	// │ 1  │ alice │ %         │ false  │
	// │ 2  │ bob   │ 127.0.0.1 │ false  │
	// │ 3  │ trudy │ %         │ true   │
	// └────┴───────┴───────────┴────────┘
}
```

### Setting-style

```go
package boxed_test

import (
	"fmt"
	"strings"
	
	"github.com/golistic/boxed"
	"github.com/golistic/boxed/style"
)

func ExampleSettings() {

	box := boxed.New(boxed.WithStyle(style.Record))

	_ = box.Append(
		boxed.NewRow("username", "alice"),
		boxed.NewRow("host", "%"),
		boxed.NewRow("locked", true),
		boxed.NewRow("max_questions", 2000),
	)

	// testable example trims spaces
	for _, l := range strings.Split(box.RenderAsString(), "\n") {
		fmt.Println(strings.TrimRight(l, " "))
	}

	// Output:
	// username      : alice
	// host          : %
	// locked        : true
	// max_questions : 2000
}
```

# Styles

## Basic

```
box := boxed.New(boxed.WithStyle(style.Basic))
```

```
┌────┬───────┬───────────┬────────┐
│ id │ user  │ host      │ locked │
├────┼───────┼───────────┼────────┤
│ 1  │ alice │ %         │ false  │
│ 2  │ bob   │ 127.0.0.1 │ false  │
│ 3  │ trudy │ %         │ true   │
└────┴───────┴───────────┴────────┘
```

### BasicNoOuter

```
box := boxed.New(boxed.WithStyle(style.BasicNoOuter))
```

```
 id │ user  │ host      │ locked 
════╪═══════╪═══════════╪════════
 1  │ alice │ %         │ false  
 2  │ bob   │ 127.0.0.1 │ false  
 3  │ trudy │ %         │ true
```

## ANSI

This mimics the MySQL CLI output.

```
box := boxed.New(boxed.WithStyle(style.ANSI))
```

```
+----+-------+-----------+--------+
| id | user  | host      | locked |
+----+-------+-----------+--------+
| 1  | alice | %         | false  |
| 2  | bob   | 127.0.0.1 | false  |
| 3  | trudy | %         | true   |
+----+-------+-----------+--------+
```

### ANSINoOuter

```
box := boxed.New(boxed.WithStyle(style.ANSINoOuter))
```

This mimics the PostgreSQL CLI output.

```
 id | user  | host      | locked
----+-------+-----------+--------
 1  | alice | %         | false
 2  | bob   | 127.0.0.1 | false
 3  | trudy | %         | true
```

## Record

```
box := boxed.New(boxed.WithStyle(style.Record))
```

No boxes here: we are showing a "record" which is actually a table with only 2 columns.

```
username      : alice
host          : %
locked        : true
max_questions : 2000
```

License
-------

Distributed under the MIT license. See `LICENSE.txt` for more information.
