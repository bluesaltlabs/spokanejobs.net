package scrapers

import "slices"

// CompanySlugs contains all available scraper company slugs
// This is the single source of truth for company names
var CompanySlugs = []string{
	"avista",
	"city_of_spokane",
	"codespeed",
	"cvsd",
	"egnyte",
	"enhanced_software_products",
	"f5",
	"gestalt_diagnostics",
	"gravity_jack",
	"huntwood",
	"intellitect",
	"itron",
	"kaiser_aluminum",
	"limelyte",
	"numerica",
	"openeye",
	"paytrace",
	"providence",
	"risklens",
	"scld",
	"sel",
	"spokane_computer",
	"spokane_library",
	"sps",
	"stcu",
	"synergisticit",
	"treasury4",
	"two_barrels",
	"urm",
	"wagstaff",
	"winco",
}

// IsValidCompanySlug checks if a given slug is valid
func IsValidCompanySlug(slug string) bool {
	return slices.Contains(CompanySlugs, slug)
}

// GetCompanySlugs returns a copy of the company slugs slice
func GetCompanySlugs() []string {
	slugs := make([]string, len(CompanySlugs))
	copy(slugs, CompanySlugs)
	return slugs
}
