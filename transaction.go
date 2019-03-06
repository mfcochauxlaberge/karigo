package karigo

// Tx ...
type Tx func(*Snapshot)

// TxNotImplemented ...
func TxNotImplemented(snap *Snapshot) {
	snap.Fail(ErrNotImplemented)
}

// TxUpdatePlayer ...
func TxUpdatePlayer(snap *Snapshot) {
	id := snap.Res.GetID()
	score := snap.Res.Get("score").(int)

	// snap.LockAll()
	// snap.Lock("players")
	// snap.Ready()

	snap.Add(Op{
		Key: Key{
			Set:   "players",
			ID:    id,
			Field: "score",
		},
		Op:    OpSet,
		Value: score,
	})

	snap.Flush()

	ranks := snap.Collection(QueryCol{
		Set:    "players",
		Fields: []string{"id"},
		Sort:   []string{"-score"},
	})

	for i := range ranks {
		snap.Add(Op{
			Key: Key{
				Set:   "players",
				ID:    ranks[i]["id"].(string),
				Field: "rank",
			},
			Op:    OpSet,
			Value: i + 1,
		})
	}

	// snap.Release("players")
	// snap.ReleaseAll()

	snap.Commit()
}
