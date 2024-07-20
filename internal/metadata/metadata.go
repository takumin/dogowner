package metadata

var (
	appName    string = "dogowner"
	appDesc    string = "CLI wrapper tool with presets for commonly used Reviewdog settings"
	authorName string = "Takumi Takahashi"
)

func AppName() string {
	return appName
}

func AppDesc() string {
	return appDesc
}

func AuthorName() string {
	return authorName
}
