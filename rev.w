\input header

@** Introduction.
This is quick-n-dirty implementation of \.{Rev} command for \.{Acme}. It behaves like \.{Look} but in reverse order.


@** Implementation.
@c

@i license

package main

import (
	@<Imports@>
)@#

@
@c
func main () {
	@<Obtaining of |id| of a window@>
	@<Opening the window by |id|@> 
	@<Initial search@>
	@<Processing of events@>
}


@
@<Imports@>=
"os"
"strconv"

@
@<Obtaining of |id| of a window@>=
id, err:=strconv.Atoi(os.Getenv("winid"))
if err!=nil {
	return
}

@
@<Imports@>=
"bitbucket.org/santucco/goacme"

@
@<Opening the window by |id|@>=
w, err:=goacme.Open(id)
if err!=nil {
	return
}

@
@<Imports@>=
"strings"

@ If a string to be looked for is specified in command line, we should use it.
@<Initial search@>=
if len(os.Args)>1 {
	s:=strings.Join(os.Args[1:], " ")
	@<Make a reverse search of |s|@>
} else 
	@<Look for selected string@> 
@<Set dot@>


@
@<Look for selected string@>=
{
	@<Set the addr address to dot@>
	@<Read the addr address into |b, e|@>
	@<Read selected string from |"xdata"| file to |s|@>
	@<Set the addr address to |b, e|@>
	@<Make a reverse search of |s|@>
}

@
@<Set the addr address to dot@>=
if w.WriteCtl("addr=dot")!=nil {
	return
}

@
@<Read the addr address into |b, e|@>=
b, e, err:=w.ReadAddr()
if err!=nil {
	return
}

@
@<Read selected string from |"xdata"| file to |s|@>=
s:=""
{
	d, err:=w.File("xdata")
	if err!=nil {
		return
	}

	buf:=make([]byte, e-b+1)

	for n, _:=d.Read(buf); n>0; n, _=d.Read(buf) {
		s+=string(buf[:n])
	}
}


@
@<Set the addr address to |b, e|@>=
if w.WriteAddr("#%d,#%d", b, e-len(s))!=nil {
	return
}

@ Reverse search is processed by writing |"?<regex>?"| to |"addr"| file, but before regex-specific symbols of |s| have to be escaped
@<Make a reverse search of |s|@>=
{
	es:=""
	for _, v:=range s {
		if strings.ContainsRune("\\/[].+?()*^$", v) {
			es+="\\"
		}
		es+=string(v)
	}
	if w.WriteAddr("?%s?", es)!=nil {
		return
	}
}

@
@<Set dot@>=
if	w.WriteCtl("dot=addr\nshow")!=nil {
	return
}

@ Events is being read from |w| and processed in next manner:

\yskip\item{$\bullet$} for text manupulating events the program is finished
\yskip\item{$\bullet$} for command executions |"Rev"| command is processed, |"Look"| command finishes the program, others are sent back to \.{Acme}.
\yskip\item{$\bullet$} looking is processed in reverse order   
 
|ev| can contain a needed string in |ev.Arg| for |"Rev"| command and in |ev.Text| and |ev.Arg| for looking event. In the latter case 
non-empty |ev.Arg| is joined to |s|
 
@<Processing of events@>=
{
	for ev, err:=w.ReadEvent(); err==nil; ev, err=w.ReadEvent() {
		if ev.Type==goacme.Insert || ev.Type==goacme.Delete {
			return
		}
		switch ev.Type&^goacme.Tag  {
			case goacme.Execute:
				switch ev.Text {
					case "Look":
						w.UnreadEvent(ev)
						return
					case "Rev":
						if len(ev.Arg)>0 {
							b:=ev.Begin
							e:=ev.End
							s:=ev.Arg
							@<Look for |s| at addr |b,e|@>
						} else
							@<Look for selected string@>
						@<Set dot@>
						continue
				}
				w.UnreadEvent(ev)
			case goacme.Look:
				b:=ev.Begin
				e:=ev.End
				s:=ev.Text
				if len(ev.Arg)>0 {
				 	s+=" "+ev.Arg
				}
				@<Look for |s| at addr |b,e|@>
				@<Set dot@>
				continue
		}
	}
}


@
@<Look for |s| at addr |b,e|@>=
@<Set the addr address to |b, e|@>
@<Make a reverse search of |s|@>

