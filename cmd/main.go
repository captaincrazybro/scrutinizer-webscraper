package main

import sw "github.com/captaincrazybro/scrutinizer-webscraper"

func main() {
	repos, avg := sw.FetchScrutinizerRepos()
	sw.SendEmail(repos, avg) //Doesn't work right now, this is just a template
}
