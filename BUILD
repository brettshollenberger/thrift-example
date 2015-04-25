for file in $(ls | grep .thrift); do
  echo "Building $file"
  thrift -r --gen rb --gen go $file
done
