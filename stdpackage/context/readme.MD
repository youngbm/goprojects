
WaitGroup：需要将单一个工作分解成多个子任务，等到全部完成后，才能进行下一步，这时候用 WaitGroup 最适合了
Channel + Select：Channel 只能用在比较单纯的 Goroutine 情况下，如果要管理多个 Goroutine，建议还是 走 context 会比较适合
Context：  如果您想一次控制全部的 Goroutine，相信用 context 会是最适合不过的
