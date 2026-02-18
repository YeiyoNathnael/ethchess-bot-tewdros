package lichess

type CreateArenaTournament struct {
	Name       string
	Password   string
	Position   string
	Rated      bool
	StartDate  int64
	Streakable bool
	Variant    string

	Description    string
	HasChat        bool
	ClockIncrement int64
	ClockTime      int64
	Minutes        int64
	Berserkable    bool
	Conditions     conditions
}

// ///////////////////////////////////////////
type conditions struct {
	allowList   string
	bots        bool
	maxRating   maxRating
	minRating   minRating
	nbRatedGame nb
	teamMember  teamId
	accountAge  int64
}

// ///////////////////////////////////////////
type maxRating int64

func (max maxRating) toInt() int64 {
	return int64(max)
}

const (
	_2200 maxRating = 2200 - (iota * 100)
	_2100
	_2000
	_1900
	_1800
	_1700
	_1600
	_1500
	_1400
	_1300
	_1200
	_1100
	_1000
	_900
	_800
)

// ///////////////////////////////////////////
type minRating int64

func (min minRating) toInt() int64 {
	return int64(min)
}

const (
	m1000 minRating = iota*100 + 1000
	m1100
	m1200
	m1300
	m1400
	m1500
	m1600
	m1700
	m1800
	m1900
	m2000
	m2100
	m2200
	m2300
	m2400
	m2500
	m2600
)

// ///////////////////////////////////////////
type nb int64

func (nb nb) toInt() int64 {
	return int64(nb)
}

const (
	_0   nb = 0
	_5   nb = 5
	_10  nb = 10
	_15  nb = 15
	_20  nb = 20
	_30  nb = 30
	_40  nb = 40
	_50  nb = 50
	_75  nb = 75
	_100 nb = 100
	_150 nb = 150
	_200 nb = 200
)

// ///////////////////////////////////////////
type teamId struct {
	teamId string
}

/////////////////////////////////////////////

type ArenaTournamentResponse struct {
	Id string `json:"id"`
}
