package main

type LyricsListEntry struct {
	Song  string
	Start float32
	End   float32
	Text  string
	Next  *LyricsListEntry
}

type LyricsList struct {
	Head *LyricsListEntry
	Tail *LyricsListEntry
}

func (self *LyricsList) AddTail(new *LyricsListEntry) *LyricsList {
	if self.Tail != nil {
		self.Tail.Next = new
	} else {
		self.Head = new
	}
	self.Tail = new
	return self
}
func (self *LyricsList) AddHead(new *LyricsListEntry) *LyricsList {
	if self.Tail != nil {
		new.Next = self.Head
	} else {
		self.Tail = new
	}
	self.Head = new
	return self
}

type SpliceEntry struct {
	Song  string
	Start float32
	End   float32
	Next  *SpliceEntry
}

type SplicesList struct {
	Head *SpliceEntry
	Tail *SpliceEntry
}

func (self *SplicesList) AddTail(new *SpliceEntry) *SplicesList {
	if self.Tail != nil {
		self.Tail.Next = new
	} else {
		self.Head = new
	}
	self.Tail = new
	return self
}
func (self *SplicesList) AddHead(new *SpliceEntry) *SplicesList {
	if self.Tail != nil {
		new.Next = self.Head
	} else {
		self.Tail = new
	}
	self.Head = new
	return self
}