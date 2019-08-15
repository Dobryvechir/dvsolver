package main

type TaskConfig struct {
   SvgPaths	[]string  `json:"paths"`
   Name		string     `json:"name"`
   Params	[]float	   `json:"params"`
   SvgImage     bool       `json:"image"`    
}
func TaskExecutor(config *TaskConfig) {

}