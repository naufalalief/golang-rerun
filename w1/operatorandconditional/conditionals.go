package main

import "fmt"

// Define constants for days of the week
const (
	dayMonday    = "Monday"
	dayTuesday   = "Tuesday"
	dayWednesday = "Wednesday"
	dayThursday  = "Thursday"
	dayFriday    = "Friday"
	daySaturday  = "Saturday"
	daySunday    = "Sunday"
)

// Define constants for repeated strings
const (
	msgGreaterThanFive = "The number is greater than 5."
	msgEqualToFive     = "The number is equal to 5."
	msgLessThanFive    = "The number is less than 5."
	msgTodayIs         = "Today is "
	msgInvalidDay      = "Invalid day."
)

// ifCondition checks if the number is greater than 5.
// If the number is greater than 5, it prints "The number is greater than 5."
func ifCondition(num int) {
	if num > 5 {
		fmt.Println(msgGreaterThanFive)
	}
}

// ifElseCondition checks if the number is greater than 5.
// If the number is greater than 5, it prints "The number is greater than 5."
// Otherwise, it prints "The number is less than or equal to 5."
func ifElseCondition(num int) {
	if num > 5 {
		fmt.Println(msgGreaterThanFive)
	} else {
		fmt.Println("The number is less than or equal to 5.")
	}
}

// ifElseIfCondition checks if the number is greater than 5.
// If the number is greater than 5, it prints "The number is greater than 5."
// If the number is equal to 5, it prints "The number is equal to 5."
// Otherwise, it prints "The number is less than 5."
func ifElseIfCondition(num int) {
	if num > 5 {
		fmt.Println(msgGreaterThanFive)
	} else if num == 5 {
		fmt.Println(msgEqualToFive)
	} else {
		fmt.Println(msgLessThanFive)
	}
}

// switchCase checks the value of the day variable.
// If the value is "Monday", it prints "Today is Monday."
// If the value is "Tuesday", it prints "Today is Tuesday."
// If the value is "Wednesday", it prints "Today is Wednesday."
// If the value is "Thursday", it prints "Today is Thursday."
// If the value is "Friday", it prints "Today is Friday."
// If the value is "Saturday", it prints "Today is Saturday."
// If the value is "Sunday", it prints "Today is Sunday."
// If the value is none of the above, it prints "Invalid day."
func switchCase(day string) {
	switch day {
	case dayMonday:
		fmt.Println(msgTodayIs + dayMonday + ".")
	case dayTuesday:
		fmt.Println(msgTodayIs + dayTuesday + ".")
	case dayWednesday:
		fmt.Println(msgTodayIs + dayWednesday + ".")
	case dayThursday:
		fmt.Println(msgTodayIs + dayThursday + ".")
	case dayFriday:
		fmt.Println(msgTodayIs + dayFriday + ".")
	case daySaturday:
		fmt.Println(msgTodayIs + daySaturday + ".")
	case daySunday:
		fmt.Println(msgTodayIs + daySunday + ".")
	default:
		fmt.Println(msgInvalidDay)
	}
}

// switchCaseFallthrough checks the value of the day variable.
// If the value is "Monday", it prints "Today is Monday."
// If the value is "Tuesday", it prints "Today is Tuesday."
// If the value is "Wednesday", it prints "Today is Wednesday."
// If the value is "Thursday", it prints "Today is Thursday."
// If the value is "Friday", it prints "Today is Friday."
// If the value is "Saturday", it prints "Today is Saturday."
// If the value is "Sunday", it prints "Today is Sunday."
// If the value is none of the above, it prints "Invalid day."
func switchCaseFallthrough(day string) {
	switch day {
	case dayMonday:
		fmt.Println(msgTodayIs + dayMonday + ".")
		fallthrough
	case dayTuesday:
		fmt.Println(msgTodayIs + dayTuesday + ".")
		fallthrough
	case dayWednesday:
		fmt.Println(msgTodayIs + dayWednesday + ".")
		fallthrough
	case dayThursday:
		fmt.Println(msgTodayIs + dayThursday + ".")
		fallthrough
	case dayFriday:
		fmt.Println(msgTodayIs + dayFriday + ".")
		fallthrough
	case daySaturday:
		fmt.Println(msgTodayIs + daySaturday + ".")
		fallthrough
	case daySunday:
		fmt.Println(msgTodayIs + daySunday + ".")
		fallthrough
	default:
		fmt.Println(msgInvalidDay)
	}
}

// nestedIfElseCondition checks if the number is greater than 5.
// If the number is greater than 5, it prints "The number is greater than 5."
// If the number is equal to 5, it prints "The number is equal to 5."
// If the number is less than 5, it prints "The number is less than 5."
func nestedIfElseCondition(num int) {
	if num > 5 {
		fmt.Println(msgGreaterThanFive)
	} else {
		if num == 5 {
			fmt.Println(msgEqualToFive)
		} else {
			fmt.Println(msgLessThanFive)
		}
	}
}

// nestedSwitchCase checks the value of the day variable.
// If the value is "Monday", it prints "Today is Monday."
// If the value is "Tuesday", it prints "Today is Tuesday."
// If the value is "Wednesday", it prints "Today is Wednesday."
// If the value is "Thursday", it prints "Today is Thursday."
// If the value is "Friday", it prints "Today is Friday."
// If the value is "Saturday", it prints "Today is Saturday."
// If the value is "Sunday", it prints "Today is Sunday."
// If the value is none of the above, it prints "Invalid day."
func nestedSwitchCase(day string) {
	switch day {
	case dayMonday:
		fmt.Println(msgTodayIs + dayMonday + ".")
	case dayTuesday:
		fmt.Println(msgTodayIs + dayTuesday + ".")
	case dayWednesday:
		fmt.Println(msgTodayIs + dayWednesday + ".")
	case dayThursday:
		fmt.Println(msgTodayIs + dayThursday + ".")
	case dayFriday:
		fmt.Println(msgTodayIs + dayFriday + ".")
	default:
		switch day {
		case daySaturday:
			fmt.Println(msgTodayIs + daySaturday + ".")
		case daySunday:
			fmt.Println(msgTodayIs + daySunday + ".")
		default:
			fmt.Println(msgInvalidDay)
		}
	}
}

// conditionals is the entry point for this snippet.
func conditionals() {
	fmt.Println("ifCondition num = 10 :")
	ifCondition(10)

	fmt.Println("ifElseCondition num = 3 :")
	ifElseCondition(3)

	fmt.Println("ifElseIfCondition num = 5 :")
	ifElseIfCondition(5)

	fmt.Println("switchCase day =", dayMonday, ":")
	switchCase(dayMonday)

	fmt.Println("switchCaseFallthrough day =", dayMonday, ":")
	switchCaseFallthrough(dayMonday)

	fmt.Println("nestedIfElseCondition num = 5 :")
	nestedIfElseCondition(5)

	fmt.Println("nestedSwitchCase day =", dayMonday, ":")
	nestedSwitchCase(dayMonday)
}
