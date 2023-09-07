package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "token", "senhaToken")
	bookHotel(ctx, "hotel")
}

// Por convenção o contexto é sempre a primeira variável a ser passada em uma função
func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)

}
