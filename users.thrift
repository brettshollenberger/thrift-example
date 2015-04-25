include "types.thrift"

struct User {
  1: optional i64 id,
  2: required string first_name,
  3: required string last_name,
  4: required string email,
}

service UsersService {
  User find(1:i64 id) throws (1:types.ServiceException e)
  User create(1:User user) throws (1:types.ServiceException e)
}
