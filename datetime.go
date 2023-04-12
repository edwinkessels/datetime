package DateTime

import (
	"math"
	"strconv"
	"strings"
	"time"
)

/* Formaten:
===========================================
JAAR4		= 2018
JAAR2		= 18
JAAR 		= 2018
MAAND2		= 03
MAAND		= 3
MND			= March
DAG2		= 09
DAG			= 9
UUR2		= 09
UUR			= 9
MINUUT2		= 07
MINUUT		= 7
SECONDE2	= 03
SECONDE		= 3
MILLI	    = 933
NANO 		= 721184800
*/

const Debug bool = true

const Format_YYYYMMDD = "JAAR4MAAND2DAG2"
const Format_HHMMSS = "UUR2:MINUUT2:SECONDE2"
const Format_YYYYMMDDHHMM = "JAAR4MAAND2DAG2UUR2MINUUT2"
const Format_YYYYMMDDHHMMSS = "JAAR4MAAND2DAG2UUR2MINUUT2SECONDE2"
const Format_READABLE_TIMESTAMP = "DAG2-MAAND2-JAAR4 UUR2:MINUUT2:SECONDE2"

// Functie om het jaar te vervangen in een string
func replaceYear(inJaar int, inDate string) (outDate string) {

	// Berekenen van de verschillende JAAR-formaten
	var JAAR4 string = strconv.Itoa(inJaar)
	var JAAR2 string = strconv.Itoa(inJaar - 2000)

	if Debug {
		// fmt.Println("JAAR4", JAAR4)
		// fmt.Println("JAAR2", JAAR2)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "JAAR4", JAAR4, -1)
	outDate = strings.Replace(outDate, "JAAR2", JAAR2, -1)
	outDate = strings.Replace(outDate, "JAAR", JAAR4, -1)

	return
}

// Functie om de Maand te vervangen in een string
func replaceMonth(inMonth int, inDate string) (outDate string) {

	var MAAND = strconv.Itoa(inMonth)
	var MAAND2 = strconv.Itoa(inMonth)
	if inMonth < 10 {
		MAAND2 = "0" + MAAND2
	}

	var MND = ""
	switch inMonth {
	case 1:
		MND = "January"
	case 2:
		MND = "February"
	case 3:
		MND = "March"
	case 4:
		MND = "April"
	case 5:
		MND = "May"
	case 6:
		MND = "June"
	case 7:
		MND = "July"
	case 8:
		MND = "August"
	case 9:
		MND = "September"
	case 10:
		MND = "October"
	case 11:
		MND = "November"
	case 12:
		MND = "December"
	}

	if Debug {
		// fmt.Println("MAAND2", MAAND2)
		// fmt.Println("MAAND", MAAND)
		// fmt.Println("MND", MND)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "MAAND2", MAAND2, -1)
	outDate = strings.Replace(outDate, "MAAND", MAAND, -1)
	outDate = strings.Replace(outDate, "MND", MND, -1)

	return
}

// Functie om de Dag te vervangen in een string
func replaceDay(inDay int, inDate string) (outDate string) {

	var DAG = strconv.Itoa(inDay)
	var DAG2 = strconv.Itoa(inDay)
	if inDay < 10 {
		DAG2 = "0" + DAG2
	}

	if Debug {
		// fmt.Println("DAG2", DAG2)
		// fmt.Println("DAG", DAG)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "DAG2", DAG2, -1)
	outDate = strings.Replace(outDate, "DAG", DAG, -1)

	return
}

// Functie om de Uur te vervangen in een string
func replaceHour(inHour int, inDate string) (outDate string) {

	var UUR = strconv.Itoa(inHour)
	var UUR2 = strconv.Itoa(inHour)
	if inHour < 10 {
		UUR2 = "0" + UUR2
	}

	if Debug {
		// fmt.Println("UUR2", UUR2)
		// fmt.Println("UUR", UUR)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "UUR2", UUR2, -1)
	outDate = strings.Replace(outDate, "UUR", UUR, -1)

	return
}

// Functie om de Minuut te vervangen in een string
func replaceMinute(inMinute int, inDate string) (outDate string) {

	var MINUUT = strconv.Itoa(inMinute)
	var MINUUT2 = strconv.Itoa(inMinute)
	if inMinute < 10 {
		MINUUT2 = "0" + MINUUT2
	}

	if Debug {
		// fmt.Println("MINUUT2", MINUUT2)
		// fmt.Println("MINUUT", MINUUT)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "MINUUT2", MINUUT2, -1)
	outDate = strings.Replace(outDate, "MINUUT", MINUUT, -1)

	return
}

// Functie om de Seconde te vervangen in een string
func replaceSecond(inSecond int, inDate string) (outDate string) {

	var SECONDE = strconv.Itoa(inSecond)
	var SECONDE2 = strconv.Itoa(inSecond)
	if inSecond < 10 {
		SECONDE2 = "0" + SECONDE2
	}

	if Debug {
		// fmt.Println("SECONDE2", SECONDE2)
		// fmt.Println("SECONDE", SECONDE)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "SECONDE2", SECONDE2, -1)
	outDate = strings.Replace(outDate, "SECONDE", SECONDE, -1)

	return
}

// Functie om de MilliSeconde te vervangen in een string
func replaceMilliSecond(inMilliSecond int, inDate string) (outDate string) {

	var MILLISECONDE = strconv.Itoa(inMilliSecond)

	// Alleen de 3 meest linkse tekens nemen
	MILLISECONDE = MILLISECONDE[0:3]

	if Debug {
		// fmt.Println("MILLISECONDE", MILLISECONDE)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "MILLI", MILLISECONDE, -1)

	return
}

// Functie om de NanaSeconde te vervangen in een string
func replaceNanoSecond(inNanoSecond int, inDate string) (outDate string) {

	var NANOSECONDE = strconv.Itoa(inNanoSecond)

	if Debug {
		// fmt.Println("NANOSECONDE", NANOSECONDE)
	}

	outDate = inDate
	outDate = strings.Replace(outDate, "NANO", NANOSECONDE, -1)

	return
}

func GetCurrentDateTime() (outVal time.Time) {

	time.LoadLocation("Europe/Amsterdam")
	outVal = time.Now()
	return
}

func getMonthByName(inMonth time.Month) (outMaandNr int) {
	// Vertalen van de MaandNaam naar een nummer
	switch inMonth {
	case time.January:
		outMaandNr = 1
	case time.February:
		outMaandNr = 2
	case time.March:
		outMaandNr = 3
	case time.April:
		outMaandNr = 4
	case time.May:
		outMaandNr = 5
	case time.June:
		outMaandNr = 6
	case time.July:
		outMaandNr = 7
	case time.August:
		outMaandNr = 8
	case time.September:
		outMaandNr = 9
	case time.October:
		outMaandNr = 10
	case time.November:
		outMaandNr = 11
	case time.December:
		outMaandNr = 12
	}
	return
}

func getMonthByNumber(inMaandNr int) (outMonth time.Month) {
	// Vertalen van de MaandNaam naar een nummer
	switch inMaandNr {
	case 1:
		outMonth = time.January
	case 2:
		outMonth = time.February
	case 3:
		outMonth = time.March
	case 4:
		outMonth = time.April
	case 5:
		outMonth = time.May
	case 6:
		outMonth = time.June
	case 7:
		outMonth = time.July
	case 8:
		outMonth = time.August
	case 9:
		outMonth = time.September
	case 10:
		outMonth = time.October
	case 11:
		outMonth = time.November
	case 12:
		outMonth = time.December
	}
	return
}

// Functie om een DatumTijd te converteren
func FormatDateTime(inDateTime time.Time, inFormat string) (outDateTime string) {

	// De meegegeven datum in componenten opsplitsen
	var jaar int = inDateTime.Year()
	var maand int = getMonthByName(inDateTime.Month())
	var dag int = inDateTime.Day()
	var uur int = inDateTime.Hour()
	var minuut int = inDateTime.Minute()
	var seconde int = inDateTime.Second()
	var millisec int = inDateTime.Nanosecond()
	var nanosec int = inDateTime.Nanosecond()

	if Debug {
		// fmt.Println("Jaar:", jaar)
		// fmt.Println("Maand:", maand)
		// fmt.Println("Dag:", dag)
		// fmt.Println("Uur:", uur)
		// fmt.Println("Minuut:", minuut)
		// fmt.Println("Seconde:", seconde)
		// fmt.Println("Milisec:", millisec)
		// fmt.Println("Nanosec:", nanosec)
	}

	outDateTime = replaceYear(jaar, inFormat)
	outDateTime = replaceMonth(maand, outDateTime)
	outDateTime = replaceDay(dag, outDateTime)
	outDateTime = replaceHour(uur, outDateTime)
	outDateTime = replaceMinute(minuut, outDateTime)
	outDateTime = replaceSecond(seconde, outDateTime)
	outDateTime = replaceMilliSecond(millisec, outDateTime)
	outDateTime = replaceNanoSecond(nanosec, outDateTime)

	return
}

// Methode voor het opvragen van een kort timestamp (yyyymmddhhmmss)
func GetTimeStamp() (outTimestamp string) {
	t := GetCurrentDateTime()
	outTimestamp = FormatDateTime(t, "JAAR4MAAND2DAG2UUR2MINUUT2SECONDE2")
	return
}

// Methode voor het opvragen van een kort timestamp (yyyymmddhhmmss.nano)
func GetLongTimeStamp() (outTimestamp string) {
	t := GetCurrentDateTime()
	outTimestamp = FormatDateTime(t, "JAAR4MAAND2DAG2UUR2MINUUT2SECONDE2.NANO")
	return
}

// Methode om een aantal uur op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddHours(inDate time.Time, inHours int) (outDate time.Time) {
	outDate = inDate.Add(time.Duration(inHours) * time.Hour)
	return
}

// Methode om een aantal minuten op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddMinutes(inDate time.Time, inMins int) (outDate time.Time) {
	outDate = inDate.Add(time.Duration(inMins) * time.Minute)
	return
}

// Methode om een aantal seconden op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddSeconds(inDate time.Time, inSecs int) (outDate time.Time) {
	outDate = inDate.Add(time.Duration(inSecs) * time.Second)
	return
}

// Methode om een aantal jaar op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddYears(inDate time.Time, inYear int) (outDate time.Time) {
	outDate = inDate.AddDate(inYear, 0, 0)
	return
}

// Methode om een aantal maanden op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddMonths(inDate time.Time, inMonths int) (outDate time.Time) {
	outDate = inDate.AddDate(0, inMonths, 0)
	return
}

// Methode om een aantal dagen op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddDays(inDate time.Time, inDays int) (outDate time.Time) {
	outDate = inDate.AddDate(0, 0, inDays)
	return
}

// Methode om een aantal weken op te tellen bij een opgegeven datum
// Het resultaat is een datum
func DateAddWeeks(inDate time.Time, inWeeks int) (outDate time.Time) {
	outDate = inDate.AddDate(0, 0, inWeeks*7)
	return
}

// Methode om van de opgegeven datum het tijd element op 0 te zetten
// Hierbij wordt geen rekening gehouden met de milli en nano-seconden
func TruncateDate(inDate time.Time) (outDate time.Time) {

	var ok = false
	outDate = inDate

	// Corrigeren van de uren
	if outDate.Hour() == 0 {
		ok = true
	}
	for ok == false {
		outDate = DateAddHours(outDate, -1)
		if outDate.Hour() == 0 {
			ok = true
		}
	}

	// Corrigeren van de minuten
	if outDate.Minute() != 0 {
		ok = false
	}
	for ok == false {
		outDate = DateAddMinutes(outDate, -1)
		if outDate.Minute() == 0 {
			ok = true
		}
	}

	// Corrigeren van de seconden
	if outDate.Second() != 0 {
		ok = false
	}
	for ok == false {
		outDate = DateAddSeconds(outDate, -1)
		if outDate.Second() == 0 {
			ok = true
		}
	}

	return
}

// Methode om een datum samen te stellen op basis van dag, maand en jaar
// Het tijd-gedeelte wordt op 0 gezet
func SetDate(inYear int, inMonth int, inDay int) (outDate time.Time) {

	// Uitgangspositie is het huidige tijdstip
	tmpDate := GetCurrentDateTime()

	// Stap 1: het goedzetten van het jaar
	var jaarOK = false
	for jaarOK == false {
		if tmpDate.Year() > inYear {
			tmpDate = tmpDate.AddDate(-1, 0, 0)
		}
		if tmpDate.Year() < inYear {
			tmpDate = tmpDate.AddDate(1, 0, 0)
		}
		if tmpDate.Year() == inYear {
			jaarOK = true
		}

		// Stap 2: Het goedzetten van de Maand
		//var Maand = getMonthByNumber(inMonth)
		var maandOK = false
		for maandOK == false {
			tmp := getMonthByName(tmpDate.Month())
			if tmp > inMonth {
				tmpDate = tmpDate.AddDate(0, -1, 0)
			}
			if tmp < inMonth {
				tmpDate = tmpDate.AddDate(0, 1, 0)
			}
			if tmp == inMonth {
				maandOK = true
			}

			// Stap 3: Het goedzetten van de dag
			var dagOk = false
			for dagOk == false {
				if tmpDate.Day() > inDay {
					tmpDate = tmpDate.AddDate(0, 0, -1)
				}
				if tmpDate.Day() < inDay {
					tmpDate = tmpDate.AddDate(0, 0, 1)
				}
				if tmpDate.Day() == inDay {
					dagOk = true
				}
			}

		}

	}

	// Truncaten van het tijdcomponent
	outDate = TruncateDate(tmpDate)
	return
}

// Functie voor het bepalen van het verschil in (seconden, minuten, uren) tussen twee Timestamps
func GetElapsedTimeInSeconds(inTime1 time.Time, inTime2 time.Time) (outElapsedTime float64) {

	elaspedTime := inTime1.Sub(inTime2)
	outElapsedTime = math.Abs(elaspedTime.Seconds())
	return
}

// Functie om een readable TimeStamp (yyyy-mm-dd hh24:mm:ss) te bepalen
func GetReadableTimestamp() (outTimestamp string) {
	now := GetCurrentDateTime()
	outTimestamp = FormatDateTime(now, Format_READABLE_TIMESTAMP)
	return
}
