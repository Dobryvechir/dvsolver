package main

type KondKonfig struct {
     BaseFolder string `json:"baseFolder"`
     Resolution int    `json:"resolution"`
     OnIdle     int    `json:"onIdle"`
}

var ConfigParameters *KondKonfig
const (
    ConfigFileName = "dvkond.conf"
)

func tryReadConfigIn(folder string) (*KondKonfig, bool) {
}

func provideConfigDefaultValues() {
}

func readConfigParameters() {
     kondKonfig,ok:=tryReadConfigIn(".")
     if !ok {
          s:=os.Args[0]
          n:=strings.LastIndex(s,"/")
          p:=strings.LastIndex(s,"\\")
          if p>n {
             n=p
          }
          if p>1 {
                s = s[:p]
                kondKonfig,_=tryReadConfigIn(s)
          } 
     }
     if kondKonfig==nil {
         kondKonfig = &KondKonfig{}
     }
     ConfigParameters = kondKonfig
     provideConfigDefaultValues()
}
