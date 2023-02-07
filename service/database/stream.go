package database

import (
	"WASAPhoto/service/objects"
	"fmt"
)

func (db *appdbimpl) GetStream(idReq uint64) ([]objects.PhotoMetadata, error) {
	// TODO  Enforce a limit and add from date parameter to endpoint in order to retrieve older posts
	// TODO remember to order by data in reverse
	// TODO remember to return empty list instead of nil
	query := fmt.Sprintf(`SELECT id FROM %s WHERE owner_id IN 
	( SELECT follow_id FROM %s WHERE id=%d) or owner_id=%d ORDER BY time_stamp DESC`,
		MEDIATABLE, FOLLOWERSTABLE, idReq, idReq)
	raws, err := db.c.Query(query)
	if err != nil {
		return nil, err
	}
	var metadataList []objects.PhotoMetadata

	for raws.Next() {
		if err = raws.Err(); err != nil {
			return nil, err
		}
		var id uint64
		err = raws.Scan(&id)
		if err != nil {
			return nil, err
		}
		obj, err := db.GetMediaMetadata(idReq, id)
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, *obj)

	}
	return metadataList, nil
}
