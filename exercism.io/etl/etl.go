// Package etl implements Extract-Transform-Load (ETL) functionalities
package etl

import "strings"

// Transform crufty, legacy data over to expected new data format
func Transform(data map[int][]string) map[string]int{
	out := make(map[string]int)
	
	for k,v := range data{
		for _, s := range v{
			out[strings.ToLower(s)] += k
		}
	}
	
	return out
}