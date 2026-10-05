package uber

type PricingStrategy interface {
	getPrice() int
}

type NIGHT_CHARGES struct{}

func (*NIGHT_CHARGES) getPrice() int {
	return 150
}

type LONG_HOLIDAYS struct{}

func (*LONG_HOLIDAYS) getPrice() int {
	return 50
}

// a rule that decides what rules will apply to the ride
// but we wont make complex and return a simple night charge hardcoded

// type PricingFactory struct {
// 	holidays HolidayCalendar
// 	surge    SurgeService
// }

// func (f *PricingFactory) RulesFor(trip TripContext) []PricingStrategy {
// 	var rules []PricingStrategy
// 	if trip.IsNight() {
// 		rules = append(rules, NightCharges{})
// 	}
// 	if f.holidays.IsHoliday(trip.Time) {
// 		rules = append(rules, LongHoliday{})
// 	}
// 	if m := f.surge.Multiplier(trip.Src); m > 1 {
// 		rules = append(rules, SurgePricing{Multiplier: m})
// 	}
// 	return rules
// }

// hardcoded rule
func rule() PricingStrategy{
	return &NIGHT_CHARGES{}
}
