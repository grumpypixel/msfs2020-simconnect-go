#!/bin/bash

go run main.go --in "/c/MSFS SDK/Samples/SimvarWatcher/bin/x64/Release/SimConnect.dll" --out "../../simconnect/dllpack.go" --template template.gopher --package simconnect --function PackedSimConnectDLL
