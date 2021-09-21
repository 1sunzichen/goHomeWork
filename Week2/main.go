package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync/errgroup"
)

func StartHttpServer(srv *http.Server) error{
	http.HandleFunc("/hello",HelloServer2)
	fmt.Println("http server start")
	err:=srv.ListenAndServe()
	return err
}

func HelloServer2(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"hello,world!@\n")
}

func main(){
	ctx:=context.Background()
	ctx,cancel:=context.WithCancel(ctx)
	group,errCtx:=errorgroup.WithCancel(ctx)
	srv:=&http.Server{Addr: ":9090"}
	group.Go(func() error{
		return StartHttpServer(srv)
	})
	group.GO(func()error{
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errctx)
	})
	chanel:=make(chan os.Signal,1)
	signal.Notify(chanel)
	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): // 因为 cancel、timeout、deadline 都可能导致 Done 被 close
				return errCtx.Err()
			case <-chanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}


