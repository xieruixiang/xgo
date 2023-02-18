package server

type FilterBuild func(nex Filter) Filter
type Filter func(ctx Context)
