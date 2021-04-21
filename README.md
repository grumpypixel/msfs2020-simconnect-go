# msfs2020-simconnect-go

## tl;dr

This library a Golang interface for Microsoft Flight Simulator 2020 (MSFS2020 ) using SimConnect.

## Does it work?

Yes, it does. This package was created to power the [msfs2020-gopilot](https://github.com/grumpypixel/msfs2020-gopilot). Please be aware that the interface is far from complete.

## Where's the Documentation?

Uhm, no documenation here at the moment. Since this is still *work in progress* the code is your friend.

## Any Examples?

At the time of writing, there are two simple [examples](https://github.com/grumpypixel/msfs2020-simconnect-go/tree/master/examples) available.

[Example #1](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/master/examples/01_simconnect/main.go) shows the basic [SimConnect](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/master/simconnect/simconnect.go) interface. With this approach, you will need to manage all *simulation variables*, *requests* etc. yourself. This is most likely what you want.

[Example #2](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/master/examples/02_simmate/main.go) shows how to use the [SimMate](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/master/simconnect/simmate.go), a convenience class where the [management](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/master/simconnect/simvar_manager.go) of [SimVars](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/master/simconnect/simvar.go) is handled for you. This encapsulation works for the above mentioned [GoPilot](https://github.com/grumpypixel/msfs2020-gopilot), but it may not work for you. Just build your own - which is awesome because this package might get inspired by your creation and improvements.

## SimMate? Seriously?

Because I didn't want to call it *Something* *Something* *Manager*, that's why.

## Do you suffer from naming paralysis from time to time?

Yes. Way too often. But names do matter. And sometimes a stupid name is better than no name at all.

Nonetheless.

Cheers and Happy coding!



