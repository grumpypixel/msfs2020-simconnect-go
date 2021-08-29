# msfs2020-simconnect-go

This is a Golang interface for Microsoft Flight Simulator 2020 (MSFS2020 ) using SimConnect.

## Installation

`$ go get github.com/grumpypixel/msfs2020-simconnect-go`

## Does it work?

Yes, it does. This package was created to power the [msfs2020-gopilot](https://github.com/grumpypixel/msfs2020-gopilot).

Please note that this interface is not complete, but a lot of the SimConnect functionality has been implemented.

## Where's the Documentation?

Check out the official [SimConnect API Reference](https://docs.flightsimulator.com/html/index.htm#t=Programming_Tools%2FSimConnect%2FSimConnect_API_Reference.htm).

Apart from that, there's no other documenation at the moment. Since this is still *work in progress* the code is your friend.

So go ahead and have a look at the file [defs.go](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/simconnect/defs.go) which is, more or less, the transfused code from SimConnect.h.

## Any Examples?

At the time of writing, there are two simple [examples](https://github.com/grumpypixel/msfs2020-simconnect-go/tree/main/examples) available.

[Example #1](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/examples/01_simconnect/main.go) shows the basic [SimConnect](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/simconnect/simconnect.go) interface. With this approach, you will need to manage all *simulation variables*, *requests* etc. yourself. This is most likely what you want.

[Example #2](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/examples/02_simmate/main.go) shows how to use the [SimMate](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/simconnect/simmate.go), a convenience class where the [management](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/simconnect/simvar_manager.go) of [SimVars](https://github.com/grumpypixel/msfs2020-simconnect-go/blob/main/simconnect/simvar.go) is handled for you. This encapsulation works for the [GoPilot](https://github.com/grumpypixel/msfs2020-gopilot) above mentioned, but it may not work for you. Just build your own - which is awesome because this package might get inspired by your creation and improvements.

## SimMate? Seriously?

Because I didn't want to call it *Something* *Something* *Manager*, that's why.

## Do you suffer from naming paralysis from time to time?

Yes. Way too often. But names do matter. And sometimes a stupid name is better than no name at all.

Nonetheless.

Cheers and Happy coding!



