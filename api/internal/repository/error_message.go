package repository

import "github.com/friendsofgo/errors"

var (
	FriendshipAlreadyExist          = errors.New("orm: unable to insert into friendships: pq: duplicate key value violates unique constraint \"unique_friends\"")
	UserAlreadyExist                = errors.New("orm: unable to insert into user_accounts: pq: duplicate key value violates unique constraint \"user_accounts_email_key\"")
	AlreadyBlockedToTarget          = errors.New("orm: unable to insert into subscription: pq: duplicate key value violates unique constraint \"subscribed\"")
	InternalErrorAddUser            = errors.New("orm: unable to insert into user_accounts: all expectations were already fulfilled, call to Query 'INSERT INTO \"user_accounts\" (\"name\",\"email\") VALUES ($1,$2) RETURNING \"user_id\"' with args [{Name: Ordinal:1 Value:InternalErrorUser} {Name: Ordinal:2 Value:internal-error-email@example.com}] was not expected")
	InternalErrorCheckEmail         = errors.New("orm: failed to check if user_accounts exists: all expectations were already fulfilled, call to Query 'SELECT COUNT(*) FROM \"user_accounts\" WHERE (\"user_accounts\".\"email\" = $1) LIMIT 1;' with args [{Name: Ordinal:1 Value:internal-error-email@example.com}] was not expected")
	InternalErrorAddFriendship      = errors.New("orm: unable to insert into friendships: all expectations were already fulfilled, call to Query 'INSERT INTO \"friendships\" (\"user_email_1\",\"user_email_2\") VALUES ($1,$2) RETURNING \"friendship_id\"' with args [{Name: Ordinal:1 Value:internal-error-email_1@example.com} {Name: Ordinal:2 Value:internal-error-email_2@example.com}] was not expected")
	InternalErrorCheckFriendship    = errors.New("orm: failed to check if friendships exists: all expectations were already fulfilled, call to Query 'SELECT COUNT(*) FROM \"friendships\" WHERE ((user_email_1 = $1 AND user_email_2 = $2) OR (user_email_1 = $3 AND user_email_2 = $4)) LIMIT 1;' with args [{Name: Ordinal:1 Value:internal-error-email_1@example.com} {Name: Ordinal:2 Value:internal-error-email_2@example.com} {Name: Ordinal:3 Value:internal-error-email_2@example.com} {Name: Ordinal:4 Value:internal-error-email_1@example.com}] was not expected")
	InternalErrorBlockNonSubscribed = errors.New("orm: unable to insert into subscription: all expectations were already fulfilled, call to Query 'INSERT INTO \"subscription\" (\"requester\",\"target\",\"status\") VALUES ($1,$2,$3) RETURNING \"subscription_id\"' with args [{Name: Ordinal:1 Value:error_requester@example.com} {Name: Ordinal:2 Value:error_target@example.com} {Name: Ordinal:3 Value:blocked}] was not expected")
	InternalErrorBlockSubscribed    = errors.New("orm: unable to update all for subscription: all expectations were already fulfilled, call to ExecQuery 'UPDATE \"subscription\" SET \"status\" = $1 WHERE (\"subscription\".\"requester\" = $2) AND (\"subscription\".\"target\" = $3);' with args [{Name: Ordinal:1 Value:blocked} {Name: Ordinal:2 Value:error_requester@example.com} {Name: Ordinal:3 Value:error_target@example.com}] was not expected")
	InternalErrorCheckIfFollowed       = errors.New("orm: failed to check if subscription exists: all expectations were already fulfilled, call to Query 'SELECT COUNT(*) FROM \"subscription\" WHERE (\"subscription\".\"requester\" = $1) AND (\"subscription\".\"target\" = $2) LIMIT 1;' with args [{Name: Ordinal:1 Value:error_requester@example.com} {Name: Ordinal:2 Value:error_target@example.com}] was not expected")
	InternalErrorCheckIfBlocked        = errors.New("orm: failed to check if subscription exists: all expectations were already fulfilled, call to Query 'SELECT COUNT(*) FROM \"subscription\" WHERE (\"subscription\".\"requester\" = $1) AND (\"subscription\".\"target\" = $2) AND (\"subscription\".\"status\" = $3) LIMIT 1;' with args [{Name: Ordinal:1 Value:error_requester@example.com} {Name: Ordinal:2 Value:error_target@example.com} {Name: Ordinal:3 Value:blocked}] was not expected")
	InternalErrorSubscribeToBlocked    = errors.New("orm: unable to update all for subscription: all expectations were already fulfilled, call to ExecQuery 'UPDATE \"subscription\" SET \"status\" = $1 WHERE (\"subscription\".\"requester\" = $2) AND (\"subscription\".\"target\" = $3) AND (\"subscription\".\"status\" = $4);' with args [{Name: Ordinal:1 Value:followed} {Name: Ordinal:2 Value:error_requester@example.com} {Name: Ordinal:3 Value:error_target@example.com} {Name: Ordinal:4 Value:blocked}] was not expected")
	InternalErrorGetFriendList         = errors.New("orm: failed to assign all query results to Friendship slice: bind failed to execute query: all expectations were already fulfilled, call to Query 'SELECT \"friendships\".* FROM \"friendships\" WHERE (user_email_1 = $1 OR user_email_2 = $2);' with args [{Name: Ordinal:1 Value:already_exist_email_1@example.com} {Name: Ordinal:2 Value:already_exist_email_1@example.com}] was not expected")
	InternalErrorSubscribeToNonBlocked = errors.New("orm: unable to insert into subscription: all expectations were already fulfilled, call to Query 'INSERT INTO \"subscription\" (\"requester\",\"target\",\"status\") VALUES ($1,$2,$3) RETURNING \"subscription_id\"' with args [{Name: Ordinal:1 Value:requester@example.com} {Name: Ordinal:2 Value:target@example.com} {Name: Ordinal:3 Value:followed}] was not expected")
)