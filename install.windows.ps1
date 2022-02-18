# How to run this file:
# iwr -useb https://raw.githubusercontent.com/approvals/ApprovalTests.Go.StarterProject/main/install.windows.ps1 | iex

iwr -useb https://raw.githubusercontent.com/JayBazuzi/machine-setup/main/windows.ps1 | iex
iwr -useb https://raw.githubusercontent.com/JayBazuzi/machine-setup/main/golang-goland.ps1 | iex


& "C:\Program Files\Git\cmd\git.exe" clone https://github.com/approvals/ApprovalTests.Go.StarterProject.git C:\Code\ApprovalTests.Go.StarterProject
cd C:\Code\ApprovalTests.Go.StarterProject
