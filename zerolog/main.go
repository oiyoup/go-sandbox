package main

import (
    "fmt"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "os"
    "time"
)

type Hello uint8
const (
    Hello1 Hello = iota
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).Level(zerolog.DebugLevel)

    logger.Print("hello world")
    
    log.Debug().
        Str("Scale", "833 cents").
        Float64("Interval", 833.09).
        Msg("Fibonacci is everywhere")
    
    log.Debug().
        Str("Name", "Tom").
        Send()

    // Panic after exit && print backtrace
    //log.Panic().
    //    Msg("This is Panic Level")

    // Fatal after exit
    //log.Fatal().
    //    Msg("This is Fatal Level")
    
    log.Error().
        Msg("This is Error Level")

    log.Print("hello world")
    
    fmt.Println(int(Hello1))
}
