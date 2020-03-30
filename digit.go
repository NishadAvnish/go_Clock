package main

type placeholder [5]string
//Zero can be exported
var Zero [5]string = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}
//One can be exported
var One [5]string = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}
//Two can be exported
var Two [5]string = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}
//Three can be exported
var Three [5]string = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}
//Four can be exported
var Four [5]string = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}
//Five can be exported
var Five [5]string = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}
//Six can be exported
var Six [5]string = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}
//Seven can be exported
var Seven [5]string = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}
//Eight can be exported
var Eight [5]string = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}
//Nine can be exported
var Nine [5]string = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

//DoubleDot  can be exported and is used to seperate the time variables(hours,minutes,seconds)
var DoubleDot [5]string=placeholder{
	"   ",
    " █ ",
	"   ",
	" █ ",
	"   ",
}

//EmptyDot 
var EmptyDot [5]string=placeholder{
	"   ",
    "   ",
	"   ",
	"   ",
	"   ",
}
