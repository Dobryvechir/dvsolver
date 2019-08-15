package main

import (
     "fmt"
)

func alles() (csv []byte,svg []byte,err error) {
     conf, err1:= ReadConfig()
     if err1!=nil {
         return nil,nil,err1
     }          
     colors,err2:=ReadColors()
     if err2!=nil {
         return nil,nil,err2
     }        
     var manager *dvthermo.ManageInfo  
     manager, err = tabelleVorbereiten(conf,colors.Colors) 
     if err!=nil {
         return nil,nil,err
     }
     err = berechnen(manager)
     if err!=nil {
         return nil,nil,err
     }
     csv,err = csvProduzieren(manager)    
     if err!=nil {
         return nil,nil,err
     }
     svg,err = svgProduzieren(manager)    
     return   
}

func allesSpeichern() error {
     csv,svg,err:=alles()
     if err!=nil {
        return err
     }
     err = ioutil.WriteFile("table.csv",csv,0644)
     if err!=nil {
        return err
     }
     err = ioutil.WriteFile("chart.svg",svg,0644)
     return err            
}

func main() {
     err:=allesSpeichern()
     if err!=nil {
        fmt.Printf("Error: %s\n",err.Error())
     }
}