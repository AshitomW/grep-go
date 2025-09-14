package worklist

// keep track of the file that needs to be processed


type Entry struct{
	Path string
}

type Worklist struct{
	jobs chan Entry 	
}



func (w *Worklist) Add(work Entry){
	w.jobs <- work
}



func (w *Worklist) Next() Entry{
	j:= <- w.jobs
	return j
}



func New(buffSize int) Worklist{
	return Worklist{make(chan Entry,buffSize)}
}


func NewJob(path string) Entry{
	return Entry{path}
}



// Generate Empty Jobs , signals quit
// no of workers and add blank entry, when a worker gets a job with blank entry , it will terminate
func (w *Worklist) Finalize(numWorkers int){
	for range numWorkers{
		w.Add(Entry{""})
	}
}



