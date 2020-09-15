# Building

```shell
innose2019-rdf-server/src$ make rdf-server
go build -x rdf-server.go logic.go
WORK=/tmp/go-build928043173
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/home/aslak/.cache/go-build/e0/e0ff266ee5a17e0428550f5e2047bd692c7937662dcd66262b8c3bb52f17e8b6-d
packagefile encoding/json=/usr/lib/go-1.13/pkg/linux_amd64/encoding/json.a
packagefile fmt=/usr/lib/go-1.13/pkg/linux_amd64/fmt.a
packagefile github.com/sbinet/go-python=/home/aslak/vcs/go/pkg/linux_amd64/github.com/sbinet/go-python.a
packagefile io/ioutil=/usr/lib/go-1.13/pkg/linux_amd64/io/ioutil.a
packagefile net/http=/usr/lib/go-1.13/pkg/linux_amd64/net/http.a
packagefile os=/usr/lib/go-1.13/pkg/linux_amd64/os.a
packagefile runtime=/usr/lib/go-1.13/pkg/linux_amd64/runtime.a
packagefile bytes=/usr/lib/go-1.13/pkg/linux_amd64/bytes.a
packagefile encoding=/usr/lib/go-1.13/pkg/linux_amd64/encoding.a
packagefile encoding/base64=/usr/lib/go-1.13/pkg/linux_amd64/encoding/base64.a
packagefile errors=/usr/lib/go-1.13/pkg/linux_amd64/errors.a
packagefile io=/usr/lib/go-1.13/pkg/linux_amd64/io.a
packagefile math=/usr/lib/go-1.13/pkg/linux_amd64/math.a
packagefile reflect=/usr/lib/go-1.13/pkg/linux_amd64/reflect.a
packagefile sort=/usr/lib/go-1.13/pkg/linux_amd64/sort.a
packagefile strconv=/usr/lib/go-1.13/pkg/linux_amd64/strconv.a
packagefile strings=/usr/lib/go-1.13/pkg/linux_amd64/strings.a
packagefile sync=/usr/lib/go-1.13/pkg/linux_amd64/sync.a
packagefile unicode=/usr/lib/go-1.13/pkg/linux_amd64/unicode.a
packagefile unicode/utf16=/usr/lib/go-1.13/pkg/linux_amd64/unicode/utf16.a
packagefile unicode/utf8=/usr/lib/go-1.13/pkg/linux_amd64/unicode/utf8.a
packagefile internal/fmtsort=/usr/lib/go-1.13/pkg/linux_amd64/internal/fmtsort.a
packagefile runtime/cgo=/usr/lib/go-1.13/pkg/linux_amd64/runtime/cgo.a
packagefile syscall=/usr/lib/go-1.13/pkg/linux_amd64/syscall.a
packagefile path/filepath=/usr/lib/go-1.13/pkg/linux_amd64/path/filepath.a
packagefile time=/usr/lib/go-1.13/pkg/linux_amd64/time.a
packagefile bufio=/usr/lib/go-1.13/pkg/linux_amd64/bufio.a
packagefile compress/gzip=/usr/lib/go-1.13/pkg/linux_amd64/compress/gzip.a
packagefile container/list=/usr/lib/go-1.13/pkg/linux_amd64/container/list.a
packagefile context=/usr/lib/go-1.13/pkg/linux_amd64/context.a
packagefile crypto/rand=/usr/lib/go-1.13/pkg/linux_amd64/crypto/rand.a
packagefile crypto/tls=/usr/lib/go-1.13/pkg/linux_amd64/crypto/tls.a
packagefile encoding/binary=/usr/lib/go-1.13/pkg/linux_amd64/encoding/binary.a
packagefile vendor/golang.org/x/net/http/httpguts=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/net/http/httpguts.a
packagefile vendor/golang.org/x/net/http/httpproxy=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/net/http/httpproxy.a
packagefile vendor/golang.org/x/net/http2/hpack=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/net/http2/hpack.a
packagefile vendor/golang.org/x/net/idna=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/net/idna.a
packagefile log=/usr/lib/go-1.13/pkg/linux_amd64/log.a
packagefile math/rand=/usr/lib/go-1.13/pkg/linux_amd64/math/rand.a
packagefile mime=/usr/lib/go-1.13/pkg/linux_amd64/mime.a
packagefile mime/multipart=/usr/lib/go-1.13/pkg/linux_amd64/mime/multipart.a
packagefile net=/usr/lib/go-1.13/pkg/linux_amd64/net.a
packagefile net/http/httptrace=/usr/lib/go-1.13/pkg/linux_amd64/net/http/httptrace.a
packagefile net/http/internal=/usr/lib/go-1.13/pkg/linux_amd64/net/http/internal.a
packagefile net/textproto=/usr/lib/go-1.13/pkg/linux_amd64/net/textproto.a
packagefile net/url=/usr/lib/go-1.13/pkg/linux_amd64/net/url.a
packagefile path=/usr/lib/go-1.13/pkg/linux_amd64/path.a
packagefile sync/atomic=/usr/lib/go-1.13/pkg/linux_amd64/sync/atomic.a
packagefile internal/oserror=/usr/lib/go-1.13/pkg/linux_amd64/internal/oserror.a
packagefile internal/poll=/usr/lib/go-1.13/pkg/linux_amd64/internal/poll.a
packagefile internal/syscall/unix=/usr/lib/go-1.13/pkg/linux_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/lib/go-1.13/pkg/linux_amd64/internal/testlog.a
packagefile internal/bytealg=/usr/lib/go-1.13/pkg/linux_amd64/internal/bytealg.a
packagefile internal/cpu=/usr/lib/go-1.13/pkg/linux_amd64/internal/cpu.a
packagefile runtime/internal/atomic=/usr/lib/go-1.13/pkg/linux_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/lib/go-1.13/pkg/linux_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/lib/go-1.13/pkg/linux_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/usr/lib/go-1.13/pkg/linux_amd64/internal/reflectlite.a
packagefile math/bits=/usr/lib/go-1.13/pkg/linux_amd64/math/bits.a
packagefile internal/race=/usr/lib/go-1.13/pkg/linux_amd64/internal/race.a
packagefile compress/flate=/usr/lib/go-1.13/pkg/linux_amd64/compress/flate.a
packagefile hash/crc32=/usr/lib/go-1.13/pkg/linux_amd64/hash/crc32.a
packagefile crypto/aes=/usr/lib/go-1.13/pkg/linux_amd64/crypto/aes.a
packagefile crypto/cipher=/usr/lib/go-1.13/pkg/linux_amd64/crypto/cipher.a
packagefile math/big=/usr/lib/go-1.13/pkg/linux_amd64/math/big.a
packagefile crypto=/usr/lib/go-1.13/pkg/linux_amd64/crypto.a
packagefile crypto/des=/usr/lib/go-1.13/pkg/linux_amd64/crypto/des.a
packagefile crypto/ecdsa=/usr/lib/go-1.13/pkg/linux_amd64/crypto/ecdsa.a
packagefile crypto/ed25519=/usr/lib/go-1.13/pkg/linux_amd64/crypto/ed25519.a
packagefile crypto/elliptic=/usr/lib/go-1.13/pkg/linux_amd64/crypto/elliptic.a
packagefile crypto/hmac=/usr/lib/go-1.13/pkg/linux_amd64/crypto/hmac.a
packagefile crypto/md5=/usr/lib/go-1.13/pkg/linux_amd64/crypto/md5.a
packagefile crypto/rc4=/usr/lib/go-1.13/pkg/linux_amd64/crypto/rc4.a
packagefile crypto/rsa=/usr/lib/go-1.13/pkg/linux_amd64/crypto/rsa.a
packagefile crypto/sha1=/usr/lib/go-1.13/pkg/linux_amd64/crypto/sha1.a
packagefile crypto/sha256=/usr/lib/go-1.13/pkg/linux_amd64/crypto/sha256.a
packagefile crypto/sha512=/usr/lib/go-1.13/pkg/linux_amd64/crypto/sha512.a
packagefile crypto/subtle=/usr/lib/go-1.13/pkg/linux_amd64/crypto/subtle.a
packagefile crypto/x509=/usr/lib/go-1.13/pkg/linux_amd64/crypto/x509.a
packagefile encoding/asn1=/usr/lib/go-1.13/pkg/linux_amd64/encoding/asn1.a
packagefile encoding/pem=/usr/lib/go-1.13/pkg/linux_amd64/encoding/pem.a
packagefile vendor/golang.org/x/crypto/chacha20poly1305=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/chacha20poly1305.a
packagefile vendor/golang.org/x/crypto/cryptobyte=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/cryptobyte.a
packagefile vendor/golang.org/x/crypto/curve25519=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/curve25519.a
packagefile vendor/golang.org/x/crypto/hkdf=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/hkdf.a
packagefile hash=/usr/lib/go-1.13/pkg/linux_amd64/hash.a
packagefile vendor/golang.org/x/text/secure/bidirule=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/text/secure/bidirule.a
packagefile vendor/golang.org/x/text/unicode/bidi=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/text/unicode/bidi.a
packagefile vendor/golang.org/x/text/unicode/norm=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/text/unicode/norm.a
packagefile mime/quotedprintable=/usr/lib/go-1.13/pkg/linux_amd64/mime/quotedprintable.a
packagefile vendor/golang.org/x/net/dns/dnsmessage=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/net/dns/dnsmessage.a
packagefile internal/nettrace=/usr/lib/go-1.13/pkg/linux_amd64/internal/nettrace.a
packagefile internal/singleflight=/usr/lib/go-1.13/pkg/linux_amd64/internal/singleflight.a
packagefile crypto/internal/subtle=/usr/lib/go-1.13/pkg/linux_amd64/crypto/internal/subtle.a
packagefile crypto/internal/randutil=/usr/lib/go-1.13/pkg/linux_amd64/crypto/internal/randutil.a
packagefile crypto/ed25519/internal/edwards25519=/usr/lib/go-1.13/pkg/linux_amd64/crypto/ed25519/internal/edwards25519.a
packagefile crypto/dsa=/usr/lib/go-1.13/pkg/linux_amd64/crypto/dsa.a
packagefile crypto/x509/pkix=/usr/lib/go-1.13/pkg/linux_amd64/crypto/x509/pkix.a
packagefile encoding/hex=/usr/lib/go-1.13/pkg/linux_amd64/encoding/hex.a
packagefile vendor/golang.org/x/crypto/cryptobyte/asn1=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/cryptobyte/asn1.a
packagefile vendor/golang.org/x/crypto/internal/chacha20=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/internal/chacha20.a
packagefile vendor/golang.org/x/crypto/internal/subtle=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/internal/subtle.a
packagefile vendor/golang.org/x/crypto/poly1305=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/crypto/poly1305.a
packagefile vendor/golang.org/x/sys/cpu=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/sys/cpu.a
packagefile vendor/golang.org/x/text/transform=/usr/lib/go-1.13/pkg/linux_amd64/vendor/golang.org/x/text/transform.a
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/lib/go-1.13/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=cjiKJC-APfVlpZRUct6y/QpIjom1v8i4kMjclgxEw/KFga3QQ6p9xVpe_z5YkI/cjiKJC-APfVlpZRUct6y -extld=gcc /home/aslak/.cache/go-build/e0/e0ff266ee5a17e0428550f5e2047bd692c7937662dcd66262b8c3bb52f17e8b6-d
/usr/lib/go-1.13/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out rdf-server
rm -r $WORK/b001/
```
