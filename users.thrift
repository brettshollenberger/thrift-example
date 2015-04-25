include "types.thrift"

struct User {
  1: required i64 id,
  2: required string first_name,
  3: required string last_name,
  4: required string email,
}

service UsersService {
  User find() throws (1:types.ServiceException e)
}
