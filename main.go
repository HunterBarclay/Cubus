package main

import (
	"fmt"
	"os"

	"golang.org/x/term"

	"github.com/HunterBarclay/cubus/src"
)

func main() {
    fmt.Println("Hello, world.")
    src.TestFunc();

    if !term.IsTerminal(int(os.Stdout.Fd())) {
        fmt.Printf("Not in a terminal\n");
        return
    }

    w, h, err := term.GetSize(int(os.Stdout.Fd()))

    if err != nil {
        fmt.Printf("err.Error(): %v\n", err.Error())
        return
    }

    fmt.Printf("(%d, %d)\n", w, h);

    buffer := make([]rune, w * h)

    for y := 0; y < h; y++ {
        for x := 0; x < w; x++ {
            var cellPos byte = boolMask(x == 0, 0) | boolMask(x == w - 1, 1) | boolMask(y == 0, 2) | boolMask(y == h - 1, 3)

            ind := y * w + x
            buffer[ind] = ' '
            
            if cellPos == 0b0101 { // Top Left
                buffer[ind] = '┌'
            } else if cellPos == 0b0110 { // Top Right
                buffer[ind] = '┐'
            } else if cellPos == 0b1010 { // Bottom Right
                buffer[ind] = '┘'
            } else if cellPos == 0b1001 { // Bottom Left
                buffer[ind] = '└'
            } else if cellPos & 0b0011 > 0 {
                buffer[ind] = '│'
            } else if cellPos & 0b1100 > 0 {
                buffer[ind] = '─'
            }
        }
    }

    printBuff(&buffer, w, h)

    // for i := 0; i < 1000000; i++ {
    //     fmt.Printf("\r%5d")
    // }
}

func printBuff(buff *[]rune, width int, height int) {
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            fmt.Printf("%c", (*buff)[y * width + x])
        }

        if y < height - 1 {
            fmt.Print("\n");
        }
    }

    fmt.Print("\033[0;0H")
}

func boolMask(a bool, leftShift int) byte {
    if (a) {
        return 0b1 << leftShift
    } else {
        return 0
    }
}
