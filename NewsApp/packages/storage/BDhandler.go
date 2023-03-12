package storage

import "newsApp/packages/proxyStructs"

type BD interface{
	Push(ch <-chan proxyStructs.Article)
	Get(number int)[]proxyStructs.Article
}

