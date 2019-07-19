package sorting

import "time"

type Track struct {
	Title 	string
	Artist 	string
	Album 	string
	Year	int
	Length	time.Duration
}

var tracks = []*Track{
	{Title: "Go", Artist: "Delilah", Album: "From the Roots Up", Year: 2012, Length: length(  "3m38s")},
	{Title: "Go", Artist: "Moby", Album: "Moby", Year: 1992, Length: length(  "3m37s")},
	{Title: "Go Ahead", Artist: "Alicia Keys", Album: "As I Am", Year: 2017, Length: length( "4m36s")},
	{Title: "Ready 2 Go", Artist: "Martin Solveig", Album: "Smash", Year: 2011, Length: length( "4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return  d
}