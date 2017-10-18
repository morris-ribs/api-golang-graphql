package testutil

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var (
	OkComputer      Disc
	TheQueenIsDead  Disc
	BeHereNow       Disc
	WhatsTheSt      Disc
	AppetiteForDest Disc
	UseYourIllusion Disc
	BackToBlack     Disc
	HotelCal        Disc
	Bad             Disc
	Thriller        Disc

	Radiohead    Artist
	TheSmiths    Artist
	GunsNR       Artist
	AmyWinehouse Artist
	Eagles       Artist
	Oasis        Artist
	MichaelJ     Artist

	DiscData   map[int]Disc
	ArtistData map[int]Artist

	discType   *graphql.Object
	artistType *graphql.Object

	MusicSchema graphql.Schema
)

type Disc struct {
	Id    string
	Title string
	Year  int
}

type Artist struct {
	Id      string
	Name    string
	Country string
	Style   string
	Discs   []Disc
}

func init() {
	OkComputer = Disc{
		Title: "OK Computer",
		Year:  1997,
		Id:    "1",
	}
	TheQueenIsDead = Disc{
		Title: "The Queen is dead",
		Year:  1986,
		Id:    "2",
	}
	BeHereNow = Disc{
		Title: "Be Here Now",
		Year:  1997,
		Id:    "3",
	}
	AppetiteForDest = Disc{
		Title: "Appetite for Destruction",
		Year:  1987,
		Id:    "4",
	}
	BackToBlack = Disc{
		Title: "Back To Black",
		Year:  2006,
		Id:    "5",
	}
	HotelCal = Disc{
		Title: "Hotel California",
		Year:  1976,
		Id:    "6",
	}
	BeHereNow = Disc{
		Title: "What's the story Morning Glory",
		Year:  1997,
		Id:    "7",
	}
	UseYourIllusion = Disc{
		Title: "Use Your Illusion",
		Year:  1997,
		Id:    "8",
	}
	Bad = Disc{
		Title: "Bad",
		Year:  1987,
		Id:    "9",
	}
	Thriller = Disc{
		Title: "Thriller",
		Year:  1983,
		Id:    "10",
	}
	Radiohead = Artist{
		Id:      "1000",
		Name:    "Radiohead",
		Country: "UK",
		Style:   "Rock",
		Discs:   []Disc{OkComputer},
	}
	TheSmiths = Artist{
		Id:      "1001",
		Name:    "The Smiths",
		Country: "UK",
		Style:   "Rock",
		Discs:   []Disc{TheQueenIsDead},
	}
	GunsNR = Artist{
		Id:      "1002",
		Name:    "Guns N' Roses",
		Country: "US",
		Style:   "Rock",
		Discs:   []Disc{AppetiteForDest, UseYourIllusion},
	}
	AmyWinehouse = Artist{
		Id:      "1003",
		Name:    "Amy Winehouse",
		Country: "UK",
		Style:   "Blues",
		Discs:   []Disc{BackToBlack},
	}
	Eagles = Artist{
		Id:      "1004",
		Name:    "Eagles",
		Country: "US",
		Style:   "Rock",
		Discs:   []Disc{HotelCal},
	}
	Oasis = Artist{
		Id:      "1005",
		Name:    "Oasis",
		Country: "UK",
		Style:   "Rock",
		Discs:   []Disc{WhatsTheSt, BeHereNow},
	}
	MichaelJ = Artist{
		Id:      "1006",
		Name:    "Michael Jackson",
		Country: "US",
		Style:   "Pop",
		Discs:   []Disc{Thriller, Bad},
	}

	DiscData = map[int]Disc{
		1:  OkComputer,
		2:  TheQueenIsDead,
		3:  BeHereNow,
		4:  AppetiteForDest,
		5:  BackToBlack,
		6:  HotelCal,
		7:  WhatsTheSt,
		8:  UseYourIllusion,
		9:  Bad,
		10: Thriller,
	}

	ArtistData = map[int]Artist{
		1000: Radiohead,
		1001: TheSmiths,
		1002: GunsNR,
		1003: AmyWinehouse,
		1004: Eagles,
		1005: Oasis,
		1006: MichaelJ,
	}

	discType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Disc",
		Description: "A set of songs from one or many artists.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The Identifier of the disc.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Id, nil
					}
					return nil, nil
				},
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "The Title of the album.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Title, nil
					}
					return nil, nil
				},
			},
			"year": &graphql.Field{
				Type:        graphql.Int,
				Description: "The album release year.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Year, nil
					}
					return nil, nil
				},
			},
		},
	})

	artistType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Artist",
		Description: "A representation of an artist and their info.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The Identifier of the artist.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if artist, ok := p.Source.(Artist); ok {
						return artist.Id, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the artist.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if artist, ok := p.Source.(Artist); ok {
						return artist.Name, nil
					}
					return nil, nil
				},
			},
			"country": &graphql.Field{
				Type:        graphql.String,
				Description: "The artist country.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if artist, ok := p.Source.(Artist); ok {
						return artist.Country, nil
					}
					return nil, nil
				},
			},
			"discs": &graphql.Field{
				Type:        graphql.NewList(discType),
				Description: "The artist list of albums.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if artist, ok := p.Source.(Artist); ok {
						discs := []map[string]interface{}{}
						for _, disc := range artist.Discs {
							discs = append(discs, map[string]interface{}{
								"title": disc.Title,
								"id":    disc.Id,
							})
						}
						return artist.Discs, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"disc": &graphql.Field{
				Type: discType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the disc",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return GetDisc(id), nil
				},
			},
			"artist": &graphql.Field{
				Type: artistType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the artist",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return GetArtist(id), nil
				},
			},
		},
	})
	MusicSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

func GetDisc(id int) Disc {
	if disc, ok := DiscData[id]; ok {
		return disc
	}
	return Disc{}
}

func GetArtist(id int) Artist {
	if artist, ok := ArtistData[id]; ok {
		return artist
	}
	return Artist{}
}
