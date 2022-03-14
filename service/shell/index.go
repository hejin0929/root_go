package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"os/exec"
	"sync"
)

func main() {
	command := "./index.sh"
	cmd := exec.Command("/bin/zsh", command)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
}

func command(ctx context.Context, cmd string) error {
	// c := exec.CommandContext(ctx, "cmd", "/C", cmd)
	c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			// 其实这段去掉程序也会正常运行，只是我们就不知道到底什么时候Command被停止了，而且如果我们需要实时给web端展示输出的话，这里可以作为依据 取消展示
			select {
			// 检测到ctx.Done()之后停止读取
			case <-ctx.Done():
				if ctx.Err() != nil {
					fmt.Printf("程序出现错误: %q", ctx.Err())
				} else {
					fmt.Println("程序被终止")
				}
				return
			default:
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					return
				}
				fmt.Print(readString)
			}
		}
	}(&wg)
	err = c.Start()
	wg.Wait()
	return err
}
