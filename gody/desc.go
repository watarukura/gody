package gody

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DescOption struct {
	TableName string `validate:"required"`
	Format    string
	Header    bool
	Field     string
}

func Desc(option *DescOption, cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to get table")
		cmd.Println(err)
		os.Exit(1)
	}

	design, err := table.Design()
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to get table design")
		cmd.Println(err)
		os.Exit(1)
	}

	name := design.GetName()
	pkey := design.GetHashKeyName()
	skey := design.GetRangeKeyName()
	if skey == "" {
		skey = "_"
	}
	count := design.GetItemCount()
	var gsiNames []string
	var gsiPkeys []string
	var gsiSkeys []string
	var hasSkey []bool
	if len(design.GSI) > 0 {
		for i, idx := range design.GSI {
			gsiNames = append(gsiNames, *idx.IndexName)
			keySchema := idx.KeySchema
			hasSkey = append(hasSkey, false)
			for _, ks := range keySchema {
				if *ks.KeyType == "HASH" {
					gsiPkeys = append(gsiPkeys, *ks.AttributeName)
				}
				if *ks.KeyType == "RANGE" {
					gsiSkeys = append(gsiSkeys, *ks.AttributeName)
					hasSkey[i] = true
				}
			}
			if hasSkey[i] == true {
				gsiSkeys = append(gsiSkeys, "_")
			}
		}
	} else {
		gsiNames = append(gsiNames, "_")
	}

	result := map[string]interface{}{
		"name":  name,
		"pkey":  pkey,
		"skey":  skey,
		"count": count,
	}

	var gsiPkey string
	var gsiSkey string
	for i, name := range gsiNames {
		gsiPkey = "GSI_pkey/" + name
		gsiSkey = "GSI_skey/" + name
		result[gsiPkey] = gsiPkeys[i]
		result[gsiSkey] = gsiSkeys[i]
	}

	var resultSlice []map[string]interface{}
	resultSlice = append(resultSlice, result)

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	var formatTarget = FormatTarget{
		ddbresult: resultSlice,
		format:    option.Format,
		header:    option.Header,
		fields:    fields,
		cmd:       cmd,
	}
	Format(formatTarget)
}
