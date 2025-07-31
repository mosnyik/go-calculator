# MY SAMPLE PROJECTS

## Calculator App
## Avarage Number Calculator

THis is a simple program that simply takes an abitrary number of imputs - in the terminal or from a file and compute the avarage.

### Run the app
On mac or linux
```bash
go run .
# then you enter the numbers you want to avarage out. Once you reach the last number, on Mac or Linux - Ctrl + D on windows - Ctrl + Z + enter

```
To compute from a .txt file
on windows 
```bash
Get-Content examples/avg.txt | go run .

```

on mac
```bash
go run . < example/avg.txt
```