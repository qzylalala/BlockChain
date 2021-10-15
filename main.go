/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/14 15:43
 */

package main

func main() {
	bc := newBlockChain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
