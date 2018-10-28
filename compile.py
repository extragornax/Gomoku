import os

import subprocess, sys

p = subprocess.Popen(["powershell.exe",
               '-ExecutionPolicy',
                'Unrestricted',".\\install-gomoku.ps1"],
              stdout=sys.stdout)
p.communicate()
# os.system("powershell install-go.ps1 -version 1.11.1")
# os.system("cp gomoku ")
