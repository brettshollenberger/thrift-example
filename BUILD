for file in $(ls | grep .thrift); do
  echo "Building $file"
  thrift --gen rb --gen go $file
done
