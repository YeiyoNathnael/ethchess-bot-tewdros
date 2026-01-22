package lichess

type User struct {
	Id           string
	Username     string
	Perfs        Perfs `json:"perfs"`
	CreatedAt    int64
	SeenAt       int64
	PlayTime     PlayTime
	Url          string
	Count        Count
	Followable   bool
	Following    bool
	Blocking     bool
	Flair        string
	PatronColor  int64
	Playing      string
	Profile      Profile
	Streamer     Streamer
	Streaming    bool
	Title        bool
	TosViolation bool
	Verified     bool
}

type StreamChannel struct {
	Channel string
}

type Streamer struct {
	Twitch  StreamChannel
	Youtube StreamChannel
}
type Profile struct {
	Bio        string
	CfcRating  int64
	DsbRating  int64
	EcfRating  int64
	FideRating int64
	Flag       string
	Links      string
	Location   string
	RcfRating  int64
	RealName   string
	UscfRating int64
}

type Perfs struct {
	Bullet         PerfObj `json:"bullet"`
	Blitz          PerfObj
	Rapid          PerfObj
	Classical      PerfObj
	Correspondence PerfObj
	Puzzle         PerfObj
	Antichess      PerfObj
	Atomic         PerfObj
	Chess960       PerfObj
	Crazyhouse     PerfObj
	Horder         PerfObj
	KingOfTheHill  PerfObj
	Racer          PuzzleMode
	RacingKings    PerfObj
	Storm          PuzzleMode
	Streak         PuzzleMode
	ThreeCheck     PerfObj
	UltraBullet    PerfObj
}

type PuzzleMode struct {
	Runs  int64
	Score int64
}

type PerfObj struct {
	Games  int64
	Rating int64 `json:"rating"`
	Rd     int64
	Prog   int64
	Prov   bool
	Rank   int64
}
type PlayTime struct {
	Total int64
	Tv    int64
	Human int64
}

type Count struct {
	All      int64
	Rated    int64
	Draw     int64
	Loss     int64
	Win      int64
	Bookmark int64
	Playing  int64
	Import   int64
	Me       int64
}
