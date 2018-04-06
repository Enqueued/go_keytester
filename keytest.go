/* Created by George Wilborn
   Mainly to just see if I can and in
   the hopes of creating a better alternative
   to the xbindkeys to just check.
*/
package main
import ("fmt")//; "bufio"; "os")

/*
  Issue number 1: it seems like go uses UTF-8 as the go-to encoding model
  gotta see about if a default library does exist.
  doesnt seem to be one for reading in and holding the keyboard hostage...
*/
func main(){
  asciiS := "ABC"
  ascNum := []byte(asciiS)
  fmt.Printf("OK: String = %d", ascNum[1])
}
