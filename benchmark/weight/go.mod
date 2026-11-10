module webtyp.com/gocheck/benchmark/weight

go 1.26.8

require (
	webtyp.com/fmt v1.0.0
	webtyp.com/json v0.0.0
	webtyp.com/model v0.2.19
)

require webtyp.com/escape v0.1.0 // indirect

// Sibling checkouts: the measurement is of the local webtyp libraries.
replace (
	webtyp.com/fmt => ../../../fmt
	webtyp.com/json => ../../../json
)
