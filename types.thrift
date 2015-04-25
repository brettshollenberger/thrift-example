enum ErrorCode {
    // Client errors
    CLIENT_ERROR,
    NOT_FOUND,

    // Server errors
    SERVER_ERROR,
    NOT_IMPLEMENTED_ERROR,
}

exception ServiceException {
    1: required ErrorCode code
    2: required string message
}
