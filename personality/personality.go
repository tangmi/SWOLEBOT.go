package personality


import "math/rand" 

var r = rand.New(rand.NewSource(0))

func TELL_ME_THE_HOUR(hour int) string {
	time := "";
	if hour % 12 != 0 {
		if hour >= 12 {
			hour -= 12;
		}
		switch hour {
			case 1:
				time = "ONE";
				break;
			case 2:
				time = "TWO";
				break;
			case 3:
				time = "THREE";
				break;
			case 4:
				time = "FOUR";
				break;
			case 5:
				time = "FIVE";
				break;
			case 6:
				time = "SIX";
				break;
			case 7:
				time = "SEVEN";
				break;
			case 8:
				time = "EIGHT";
				break;
			case 9:
				time = "NINE";
				break;
			case 10:
				time = "TEN";
				break;
			case 11:
				time = "ELEVEN";
				break;
		}
	} else {
		if hour == 0 {
			time = "MIDNIGHT";
		} else if hour == 12 {
			time = "NOON";
		}
	}

	if r.Float32() < 0.2 {
		if r.Float32() >= 0.5 {
			time = "PUSHUP";
		} else {
			time = "SWOLE"
		}
	}

	time = time + " O'CLOCK";

	return time;
}

var quotes = [...]string{
	"IT'S TIME TO GET SWOLE!",
	"LETS GET SWOLE!",
	"TIME TO GET SWOLE.",
	"SWOLE UP, FOLKS!",
	"IT'S SWOLE TIME!",
	"ARE WE SWOLEING OR WHAT?!",
	"WHAT ARE YOU DOING? YOU'RE NOT GETTING SWOLE!",
	"WHAT? IS GETTING SWOLE NOT GOOD ENOUGH FOR YOU?!",
}

func INSPIRE_ME_PLEASE() string {
	return quotes[r.Intn(len(quotes))]
}