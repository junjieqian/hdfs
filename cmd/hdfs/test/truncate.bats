#!/usr/bin/env bats

load helper

setup() {
  $HDFS mkdir -p /_test_cmd/truncate
  $HDFS touch /_test_cmd/truncate/a
}

@test "truncate larger" {
  skip "Not support until Hadoop-2.7.2"
  run $HDFS truncate -s 10 /_test_cmd/a
  assert_failure
}

@test "truncate nonexistent" {
  skip "Not support until Hadoop-2.7.2"
  run $HDSF truncate -s 10 /_test_cmd/nonexistent
  assert_failure
}

@test "truncate" {
  skip "Not support until Hadoop-2.7.2"
  run $HDFS put $ROOT_TEST_DIR/test/foo.txt /_test_cmd/truncate/1
  run $HDFS truncate -s 2 /_test_cmd/truncate/1
  assert_success
}

teardown() {
  $HDFS rm -r /_test_cmd/truncate
}
