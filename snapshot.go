package ankadb

import (
	"strconv"

	"github.com/zhs007/ankadb/database"
	pb "github.com/zhs007/ankadb/proto"
)

const keySnapshotMgr = "ankadb:snapshotmgr"

// SnapshotMgr - key list snapshot
type SnapshotMgr struct {
	db          database.Database
	mgrSnapshot pb.SnapshotMgr
	mapSnapshot map[int64]*pb.Snapshot
}

func newSnapshotMgr(db database.Database) SnapshotMgr {
	return SnapshotMgr{
		db:          db,
		mapSnapshot: make(map[int64]*pb.Snapshot),
	}
}

func (mgr *SnapshotMgr) init() error {
	err := GetMsgFromDB(mgr.db, []byte(keySnapshotMgr), &mgr.mgrSnapshot)
	if err != nil {
		mgr.mgrSnapshot.MaxSnapshotID = 1

		err = PutMsg2DB(mgr.db, []byte(keySnapshotMgr), &mgr.mgrSnapshot)

		return err
	}

	return nil
}

func (mgr *SnapshotMgr) makeKey(snapshotid int64) []byte {
	return []byte("ankadb:snapshot:" + strconv.FormatInt(snapshotid, 10))
}

// Add - add snapshot
func (mgr *SnapshotMgr) Add(pSnapshot *pb.Snapshot) (int64, error) {
	pSnapshot.SnapshotID = mgr.mgrSnapshot.MaxSnapshotID + 1

	err := PutMsg2DB(mgr.db, mgr.makeKey(pSnapshot.SnapshotID), pSnapshot)
	if err != nil {
		return -1, err
	}

	err = PutMsg2DB(mgr.db, []byte(keySnapshotMgr), &mgr.mgrSnapshot)
	if err != nil {
		return -1, err
	}

	mgr.mapSnapshot[pSnapshot.SnapshotID] = pSnapshot

	return pSnapshot.SnapshotID, nil
}

// Get - get snapshot
func (mgr *SnapshotMgr) Get(snapshotid int64) *pb.Snapshot {
	pSnapshot, ok := mgr.mapSnapshot[snapshotid]
	if ok {
		return pSnapshot
	}

	pSnapshot = &pb.Snapshot{}
	err := GetMsgFromDB(mgr.db, mgr.makeKey(snapshotid), pSnapshot)
	if err != nil {
		return nil
	}

	return pSnapshot
}

// NewSnapshot - new snapshot
func (mgr *SnapshotMgr) NewSnapshot(prefix []byte) (*pb.Snapshot, error) {
	pSnapshot := &pb.Snapshot{
		SnapshotID: 0,
	}

	curit := mgr.db.NewIteratorWithPrefix(prefix)
	for curit.Next() {
		key := curit.Key()

		pSnapshot.Keys = append(pSnapshot.Keys, string(key))
	}

	curit.Release()
	err := curit.Error()
	if err != nil {
		return nil, err
	}

	_, err = mgr.Add(pSnapshot)
	if err != nil {
		return nil, err
	}

	return pSnapshot, nil
}
