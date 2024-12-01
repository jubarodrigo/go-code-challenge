package cmd

func StartApp() {
	star := NewApp()
	buildDirector := NewAppBuilder(star)

	buildDirector.BuildStarship()
}
