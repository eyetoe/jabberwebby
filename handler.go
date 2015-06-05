// Package of functions for the jabberwebby page
package jabberwebby

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	//suffix := r.URL.Path[len("/jabberwebby/"):]
	//fmt.Fprintf(w, "I see %s", suffix)

	fmt.Println("Inside Handler - jabberwebby")
	fmt.Fprintln(w, "<html>")

	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<link rel='stylesheet' type='text/css' href='jabberwebby.css' />")
	fmt.Fprintln(w, "<link href='http://fonts.googleapis.com/css?family=Metamorphous' rel='stylesheet' type='text/css'>")
	fmt.Fprintln(w, "<link href='http://fonts.googleapis.com/css?family=Gentium+Basic:400,700,700italic' rel='stylesheet' type='text/css'>")

	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")

	fmt.Fprintln(w, "<center>")
	fmt.Fprintln(w, "<table><tr><th>")

	fmt.Fprintln(w, "<div>JABBER")
	fmt.Fprintln(w, "WEBBY</div></th><th>")
	fmt.Fprintf(w, "<h2>The Day "+Firstname()+" "+Lastname()+" Defeated the "+Monster()+".</h2></th></tr>")

	fmt.Fprintf(w, "<tr><td colspan='2'>")
	fmt.Fprintf(w, "<hr class='style-one>'")
	fmt.Fprintf(w, "<p><i>It was a cold forboding morning and <u>"+Firstname()+"</u> was inspecting his <u>"+Melee()+"</u>, hoping for some action.  Just as papa <u>"+Lastname()+"</u> had taught.  You can't be too prepared.<br><br>Suddenly <u>"+Firstname()+"</u> spied a <u>"+Monster()+"</u> coming over the horizon, but still out of range of his <u>"+Long()+"</u>.  It wasn't surprised and emitted a loud aggressive sound.<br><br>As the <u>"+Monster()+"</u> charged, <u>"+Firstname()+"</u> lobbed a <u>"+Grenade()+"</u> gernade toward the foe.  The creature was momentarily stunned while <u>"+Firstname()+"</u> pulled out his <u>"+Ranged()+"</u> and aimed at the beast.  Three shots were not enought to deter it, as the <u>"+Monster()+"</u> came ever closer.<br><br>Finally the beast arrived at close range.  With a quick stroke of the <u>"+Melee()+"</u> the <u>"+Monster()+"</u> took it right in the <u>"+Part()+"</u> and died in a lump.  This would be a good tale to tell over at the tavern tonight.  Three cheers for <u>"+Firstname()+"</u> <u>"+Lastname()+"</u></p>")

	fmt.Fprintf(w, "<hr class='style-one>'")

	fmt.Fprintln(w, "</td></tr></table>")
	fmt.Fprintln(w, "</center>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}
