

/*2:*/


//line rev.w:8


//line license:1
// This file is part of rev
// Author Alexander Sychev
//
// Copyright (c) 2014, 2020 Alexander Sychev. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * The name of author may not be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//line rev.w:11

package main

import(


/*4:*/


//line rev.w:29

"os"
"strconv"



/*:4*/



/*6:*/


//line rev.w:41

"github.com/santucco/goacme"



/*:6*/



/*8:*/


//line rev.w:52

"strings"



/*:8*/


//line rev.w:15

)



/*:2*/



/*3:*/


//line rev.w:19

func main(){


/*5:*/


//line rev.w:34

id,err:=strconv.Atoi(os.Getenv("winid"))
if err!=nil{
return
}



/*:5*/


//line rev.w:21



/*7:*/


//line rev.w:45

w,err:=goacme.Open(id)
if err!=nil{
return
}



/*:7*/


//line rev.w:22



/*9:*/


//line rev.w:56

if len(os.Args)> 1{
s:=strings.Join(os.Args[1:]," ")


/*15:*/


//line rev.w:112

{
es:=""
for _,v:=range s{
if strings.ContainsRune("\\/[].+?()*^$",v){
es+= "\\"
}
es+= string(v)
}
if w.WriteAddr("?%s?",es)!=nil{
return
}
}



/*:15*/


//line rev.w:59

}else


/*10:*/


//line rev.w:66

{


/*11:*/


//line rev.w:76

if w.WriteCtl("addr=dot")!=nil{
return
}



/*:11*/


//line rev.w:68



/*12:*/


//line rev.w:82

b,e,err:=w.ReadAddr()
if err!=nil{
return
}



/*:12*/


//line rev.w:69



/*13:*/


//line rev.w:89

s:=""
{
d,err:=w.File("xdata")
if err!=nil{
return
}

buf:=make([]byte,e-b+1)

for n,_:=d.Read(buf);n> 0;n,_= d.Read(buf){
s+= string(buf[:n])
}
}




/*:13*/


//line rev.w:70



/*14:*/


//line rev.w:106

if w.WriteAddr("#%d,#%d",b,e-len(s))!=nil{
return
}



/*:14*/


//line rev.w:71



/*15:*/


//line rev.w:112

{
es:=""
for _,v:=range s{
if strings.ContainsRune("\\/[].+?()*^$",v){
es+= "\\"
}
es+= string(v)
}
if w.WriteAddr("?%s?",es)!=nil{
return
}
}



/*:15*/


//line rev.w:72

}



/*:10*/


//line rev.w:61



/*16:*/


//line rev.w:127

if w.WriteCtl("dot=addr\nshow")!=nil{
return
}



/*:16*/


//line rev.w:62





/*:9*/


//line rev.w:23



/*17:*/


//line rev.w:141

{
for ev,err:=w.ReadEvent();err==nil;ev,err= w.ReadEvent(){
if ev.Type==goacme.Insert||ev.Type==goacme.Delete{
return
}
switch ev.Type&^goacme.Tag{
case goacme.Execute:
switch ev.Text{
case"Look":
w.UnreadEvent(ev)
return
case"Rev":
if len(ev.Arg)> 0{
b:=ev.Begin
e:=ev.End
s:=ev.Arg


/*18:*/


//line rev.w:181



/*14:*/


//line rev.w:106

if w.WriteAddr("#%d,#%d",b,e-len(s))!=nil{
return
}



/*:14*/


//line rev.w:182



/*15:*/


//line rev.w:112

{
es:=""
for _,v:=range s{
if strings.ContainsRune("\\/[].+?()*^$",v){
es+= "\\"
}
es+= string(v)
}
if w.WriteAddr("?%s?",es)!=nil{
return
}
}



/*:15*/


//line rev.w:183



/*:18*/


//line rev.w:158

}else


/*10:*/


//line rev.w:66

{


/*11:*/


//line rev.w:76

if w.WriteCtl("addr=dot")!=nil{
return
}



/*:11*/


//line rev.w:68



/*12:*/


//line rev.w:82

b,e,err:=w.ReadAddr()
if err!=nil{
return
}



/*:12*/


//line rev.w:69



/*13:*/


//line rev.w:89

s:=""
{
d,err:=w.File("xdata")
if err!=nil{
return
}

buf:=make([]byte,e-b+1)

for n,_:=d.Read(buf);n> 0;n,_= d.Read(buf){
s+= string(buf[:n])
}
}




/*:13*/


//line rev.w:70



/*14:*/


//line rev.w:106

if w.WriteAddr("#%d,#%d",b,e-len(s))!=nil{
return
}



/*:14*/


//line rev.w:71



/*15:*/


//line rev.w:112

{
es:=""
for _,v:=range s{
if strings.ContainsRune("\\/[].+?()*^$",v){
es+= "\\"
}
es+= string(v)
}
if w.WriteAddr("?%s?",es)!=nil{
return
}
}



/*:15*/


//line rev.w:72

}



/*:10*/


//line rev.w:160



/*16:*/


//line rev.w:127

if w.WriteCtl("dot=addr\nshow")!=nil{
return
}



/*:16*/


//line rev.w:161

continue
}
w.UnreadEvent(ev)
case goacme.Look:
b:=ev.Begin
e:=ev.End
s:=ev.Text
if len(ev.Arg)> 0{
s+= " "+ev.Arg
}


/*18:*/


//line rev.w:181



/*14:*/


//line rev.w:106

if w.WriteAddr("#%d,#%d",b,e-len(s))!=nil{
return
}



/*:14*/


//line rev.w:182



/*15:*/


//line rev.w:112

{
es:=""
for _,v:=range s{
if strings.ContainsRune("\\/[].+?()*^$",v){
es+= "\\"
}
es+= string(v)
}
if w.WriteAddr("?%s?",es)!=nil{
return
}
}



/*:15*/


//line rev.w:183



/*:18*/


//line rev.w:172



/*16:*/


//line rev.w:127

if w.WriteCtl("dot=addr\nshow")!=nil{
return
}



/*:16*/


//line rev.w:173

continue
}
}
}




/*:17*/


//line rev.w:24

}




/*:3*/


