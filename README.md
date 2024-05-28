# Ascii-art-fs

* This is a program that receives a string as an argument and outputting the string in a graphic representation using ASCII characters.

* The program reads from ASCII art files to generate the modified art.

* The program ensures the correct representation of each character, adhering to ASCII standards.

* The usage respect the format go run . [STRING] [BANNER], any other formats returns the following usage message:

## Usage

```bash
go run . [STRING] [BANNER]
```

``` bash 
go run . something standard 
```


* Three types of ASCII art are available:

1. **Standard**: Default ASCII art
```bash
 _    _  $
| |  | | $
| |__| | $
|  __  | $
| |  | | $
|_|  |_| $
         $
         $
```
2. **Shadow**: ASCII art with shadow effect
```bash
         $
_|    _| $
_|    _| $
_|_|_|_| $
_|    _| $
_|    _| $
         $
         $
```


3. **Thinkertoy**: ASCII art resembling thinkertoy
```bash
     $
o  o $
|  | $
O--O $
|  | $
o  o $
     $
     $
```

# Features

* Read from the commandline
* Converts string into graphic representation
* Program handles input with numbers, letters, spaces, special characters and \n.
* Based on the flag, program gives the user an opportunity to select which ASCII art file to use
*  Program uses flags;  "shadow", "thinkertoy" and "standard" for "shawdow.txt", "thinkertoy.txt" and "standard.txt" files respectively

# Installation

1. Clone the repository

```bash
git clone https://learn.zone01kisumu.ke/git/beopiyo/ascii-art-fs
```

2. Navigate to the project repository

``` bash
cd ascii-art-fs
```

3. Run the project

```bash
go run . "hello" standard | cat -e
```

# Project Usage

## Instructions

To generate different types of ascii-art, run the following command:

1. **Standard.txt**

* Input

``` bash
go run . "Hello\nThere" standard | cat -e
```

* Output

```bash
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                            $
|__   __| | |                           $
   | |    | |__     ___  _ __    ___    $
   | |    |  _ \  / _ \ | '__|  / _  \  $
   | |    | | | | |  __/ | |    |  __/  $
   |_|    |_| |_|  \___| |_|     \___|  $
                                        $
                                        $

```
2. **Shadow.txt**

* Input

```bash
go run . "Hello\nThere" shadow | cat -e
```
* Output

```bash
                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
                                               $
_|_|_|_|_| _|                                  $
    _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
    _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
    _|     _|    _| _|       _|       _|       $
    _|     _|    _|   _|_|_| _|         _|_|_| $
                                               $
                                               $
```

3. **Thinkertoy.txt**

* Input

```bash
go run . "Hello\nThere" thinkertoy | cat -e
```
* Output

```bash
                 $
o  o     o o     $
|  |     | |     $
O--O o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
                       $
o-O-o o                $
  |   |                $
  |   O--o o-o o-o o-o $
  |   |  | |-' |   |-' $
  o   o  o o-o o   o-o $
                       $
                       $
```


# Contributing

* Contributions are welcome! Feel free to open issues or submit pull requests. 
* To contribute, follow these steps:

1. Fork the repository.
2. Create a new branch
3. Make your changes
4. Test your changes
5. Submit a pull request


# Collaborators


* This project was made possible from the contributions made by the following team members: viz:- [Benard Opiyo](https://learn.zone01kisumu.ke/git/beopiyo), [Elijah Gathanga](https://learn.zone01kisumu.ke/git/egathang) and [Stephen Oginga](https://learn.zone01kisumu.ke/git/steodhiambo)
