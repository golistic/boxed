/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package boxed_test

import (
	"fmt"
	"strings"

	"github.com/golistic/boxed"
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

func ExampleANSI_mysql() {

	box := boxed.New(boxed.WithStyle(boxed.ANSI))

	_ = box.AddHeader("id", "user", "host", "locked")
	_ = box.Append(
		boxed.NewRow("1", "alice", "%", false),
		boxed.NewRow("2", "bob", "127.0.0.1", false),
		boxed.NewRow("3", "trudy", "%", true),
	)

	_ = box.Render()

	// Output:
	// +----+-------+-----------+--------+
	// | id | user  | host      | locked |
	// +----+-------+-----------+--------+
	// | 1  | alice | %         | false  |
	// | 2  | bob   | 127.0.0.1 | false  |
	// | 3  | trudy | %         | true   |
	// +----+-------+-----------+--------+
}

func ExampleANSI_psql() {

	box := boxed.New(boxed.WithStyle(boxed.ANSINoOuter))

	_ = box.AddHeader("id", "user", "host", "locked")
	_ = box.Append(
		boxed.NewRow("1", "alice", "%", false),
		boxed.NewRow("2", "bob", "127.0.0.1", false),
		boxed.NewRow("3", "trudy", "%", true),
	)

	// testable example trims spaces; we prepend a dot as workaround
	for _, l := range strings.Split(box.RenderAsString(), "\n") {
		fmt.Println(".   ", strings.TrimRight(l, " "))
	}

	// Output:
	// .     id | user  | host      | locked
	// .    ----+-------+-----------+--------
	// .     1  | alice | %         | false
	// .     2  | bob   | 127.0.0.1 | false
	// .     3  | trudy | %         | true
	// .
}

func ExampleRecord() {

	box := boxed.New(boxed.WithStyle(boxed.Record))

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
