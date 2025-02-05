package relations

type Relation struct {
	UserID 		string `bson:"user_id" json:"user_id"`
	UserIDRel string `bson:"user_id_rel" json:"user_id_rel"`
}