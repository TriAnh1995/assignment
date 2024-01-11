package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) CheckFriendship(ctx context.Context, userEmails []string) (bool, error) {
	exist, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail1.EQ(userEmails[0]),
		orm.RelationshipWhere.UserEmail2.EQ(userEmails[1]),
		orm.RelationshipWhere.Friendship.EQ("friend"),
	).Exists(ctx, i.pgConn)
	return exist, err
}
