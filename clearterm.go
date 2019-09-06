package mylib

import ( 
	"os"
	"os/exec"
)


//Clearscreen ...
func Clearscreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("cmd", "/c", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}