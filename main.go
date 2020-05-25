package main

import (
    "fmt"
    "time"
    "os/exec"
    "errors"
)

func RunWithDeadline(command string, args ...string) (string, error) {
    result := make(chan string)
    cmd := exec.Command(command, args...)
    go func() {
        out, err := cmd.Output()
        if err != nil {
            return
        }
        // Отправляем значение в канал который никто УЖЕ не читает!
        result <- string(out)
    }()
    select {
    case <-time.After(2 * time.Second):
        cmd.Process.Kill()
        return "", errors.New("deadline")
    case res := <-result:
        return res, nil
    }
}


func main() {
    res, err := RunWithDeadline("sleep", "10")
    fmt.Println(res, err)
}
