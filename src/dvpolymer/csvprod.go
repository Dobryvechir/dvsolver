package main

import (
       "fmt"
)

func printTime(buf []byte,v float64) []byte {
     s:=fmt.Sprintf("%f0.2", v)
     return append(buf, []byte(s)...)
}

func printTemperature(buf []byte,v float64) []byte {
     s:=fmt.Sprintf("%f0.1", v)
     return append(buf, []byte(s)...)
}

func printNextTab(buf []byte)[]byte {
     return append(buf, ',')
}

func printEOL(buf []byte)[]byte {
    return append(buf, 13, 10)
}

func csvProduzieren(manager *dvthermo.ManageInfo) (csv []byte, err error) {
   extra,_ = manager.Extra.(*RechnenInfo)
   csv:=make([]byte,0,extra.SvgAllTime * 128)
   n:=extra.SvgAllTime
   m:=extra.CsvLen
   for i:=0;i<n;i++ {
        csv = printTime(csv, float64(i) * extra.Config.TimeShow)
        pos:=i * m
        for j:=0;j<m;j++ {
            csv = printNextTab(csv)
            csv = printTemperature(csv, extra.CsvTable[pos+j])
        }
        csv = printEOL(csv)  
   }
   return
}