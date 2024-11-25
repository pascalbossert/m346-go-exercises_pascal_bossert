func outputWithZodiacSign(p Person) {
	var zodiacSign rune

	if (p.Month == 3 && p.Day >= 21) || (p.Month == 4 && p.Day <= 20) {
		zodiacSign = Aries
	} else if (p.Month == 4 && p.Day >= 21) || (p.Month == 5 && p.Day <= 21) {
		zodiacSign = Taurus
	} else if (p.Month == 5 && p.Day >= 22) || (p.Month == 6 && p.Day <= 21) {
		zodiacSign = Gemini
	} else if (p.Month == 6 && p.Day >= 22) || (p.Month == 7 && p.Day <= 22) {
		zodiacSign = Cancer
	} else if (p.Month == 7 && p.Day >= 23) || (p.Month == 8 && p.Day <= 23) {
		zodiacSign = Leo
	} else if (p.Month == 8 && p.Day >= 24) || (p.Month == 9 && p.Day <= 23) {
		zodiacSign = Virgo
	} else if (p.Month == 9 && p.Day >= 24) || (p.Month == 10 && p.Day <= 23) {
		zodiacSign = Libra
	} else if (p.Month == 10 && p.Day >= 24) || (p.Month == 11 && p.Day <= 22) {
		zodiacSign = Scorpius
	} else if (p.Month == 11 && p.Day >= 23) || (p.Month == 12 && p.Day <= 21) {
		zodiacSign = Sagittarius
	} else if (p.Month == 12 && p.Day >= 22) || (p.Month == 1 && p.Day <= 20) {
		zodiacSign = Capricornus
	} else if (p.Month == 1 && p.Day >= 21) || (p.Month == 2 && p.Day <= 19) {
		zodiacSign = Aquarius
	} else {
		zodiacSign = Pisces
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

// unknown error, revisit later... 