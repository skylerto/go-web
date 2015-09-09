package views

type standLocator struct {
	Title  string
	Active string
}

func GetStandLocator() standLocator {
	result := standLocator{
		Title:  "Lemonade Stand Supply - Stand Locator",
		Active: "stand_locator",
	}
	return result
}
