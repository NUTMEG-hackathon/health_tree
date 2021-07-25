package model

import (
  "fmt"
)

type Tree struct {
	ID        int      `json:"id" gorm:"praimaly_key"`
	Name      string   `json:"name"`
	Type      int      `json:"type"`
	Point     float64  `json:"point"`
	UserID    int      `json:"user_id"`
	Logs     []Log
}

type Trees []Tree

func TreeIndex() Trees {
  db := Init()
  defer db.Close()

  var trees Trees
  db.Find(&trees)
  return trees
}

func CreateTree(tree *Tree) {
  db := Init()
  defer db.Close()

  db.Create(tree)
}

func FindTree(t *Tree) Tree {
  db := Init()
  defer db.Close()

  var tree Tree
  db.Where(t).Find(&tree)
  return tree
}

func UpdateTree(t *Tree) error {
  db := Init()
  defer db.Close()

	rows := db.Model(t).Update(map[string]interface{}{
		"name": t.Name,
		"type": t.Type,
		"point": t.Point,
		"user_id": t.UserID,
	}).RowsAffected
	if rows == 0 {
			return fmt.Errorf("Could not find Todo (%v) to update", t)
	}
	return nil
}

func TreeSeed(){
  db := Init()
  defer db.Close()

  trees := Trees{
    { ID:1, Name:"tree1", Type:1, Point:0, UserID:1 },
    { ID:2, Name:"tree2", Type:2, Point:0, UserID:2 },
    { ID:3, Name:"tree3", Type:3, Point:0, UserID:3 },
  }

  for i,tree := range trees{
    fmt.Println(i,tree)
    db.Create(tree)
  }

}
