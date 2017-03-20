package hdfs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTruncate(t *testing.T) {
	t.Skip("Truncate not support in Hadoop-2.6")
	client := getClient(t)

	baleet(t, "/_test/truncate/1.txt")
	mkdirp(t, "/_test/truncate")
	writer, _ := client.Create("/_test/truncate/1.txt")
	_, _ = writer.Write([]byte("foobar\nfoobar\n"))
 	_ = writer.Close()

	err := client.Truncate("/_test/truncate/1.txt", 4)
	require.NoError(t, err)

	stat, err := client.Stat("/_test/truncate/1.txt")
	require.NoError(t, err)
	assert.EqualValues(t, 4, stat.Size())

	err = client.Truncate("/_test/truncate/1.txt", 10)
	assert.NotNil(t, err)
}

func TestTruncateNoExistent(t *testing.T) {
	if os.Getenv("HADOOP_DISTRO") == "cdh" {
		t.Skip("Truncate not support in Hadoop-2.6")
	}
	client := getClient(t)

	err := client.Truncate("/_test/nonexistent", 100)
	assertPathError(t, err, "truncate", "/_test/nonexistent", os.ErrNotExist)
}
