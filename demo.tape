Output demo.gif
Set FontSize 18
Set Width 1200
Set Height 600
Set Padding 20
Set Theme "Catppuccin Mocha"


Type "ls -l" Enter
Sleep 1s

Type "go build -o auto-license autolicense.go" Enter
Sleep 1.5s

Type "cp ./auto-license ~/linux" Enter
Sleep 0.5s

Type "cp ./license-EXAMPLE.txt ~/linux/license.txt" Enter
Sleep 0.5s

Type "cp ./auto_license-config-EXAMPLE.json ~/linux/auto_license-config.json" Enter
Sleep 0.5s

Type "cd ~/linux" Enter
Sleep 0.5s

Type "time ./auto-license" Enter
Sleep 15s

Type "time ./auto-license" Enter
Sleep 10s

Sleep 2s
