# RDF Server for a 2019 InnoSE Project

This server is a [Go](https://golang.org) program that wraps some [Python](https://www.python.org) code for access to [rdflib](https://rdflib.readthedocs.io).

## Requirements

- The `python2.7-dev` package (as it is a dependency for go-python)
- The following python 2.7 modules:
  - `rdflib` (both 4.2.2 and 5.0.0 have been tested)
  - `requests` (for some reason rdflib makes use of it without depending on it)

For testin:
- `mosquitto-clients` (manual MQTT publication)

## Building

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

## Running

```shell
innose2019-rdf-server/src$ ./rdf-server
Syntax: ./rdf-server CONFIG_FILE
        ./rdf-server ../etc/default_config.json
innose2019-rdf-server/src$ ./rdf-server ../etc/default_config.json
RDFLib Version: 5.0.0
STATUS: Loading model "../var/model/2.ttl".
Listening to 0.0.0.0:8001
^C
```

**Note:** The current version of Brick is quite large and thus takes a considerable amount of time to load/store.

## Interface

### Fetch time from server

```shell
$ curl -X PUT -d '42' http://localhost:8001/time
{
    "success": true,
    "time": 1589391233.535275
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

### Fetch store model to disk

```shell
$ curl -X PUT -d '42' http://localhost:8001/store
{
    "filename": "../var/model/3.ttl",
    "success": true
}
```

### Fetch the current namespace mapping

```shell
$ curl -X PUT -d '42' http://localhost:8001/namespaces
{
    "namespaces": {
        "bldg": "https://brickschema.org/schema/1.1.0/ExampleBuilding#",
        "brick": "https://brickschema.org/schema/1.1.0/Brick#",
        "dcterms": "http://purl.org/dc/terms#",
        "n": "http://ses.sdk.dk/junk/test#",
        "owl": "http://www.w3.org/2002/07/owl#",
        "rdf": "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
        "rdfs": "http://www.w3.org/2000/01/rdf-schema#",
        "sdo": "http://schema.org#",
        "skos": "http://www.w3.org/2004/02/skos/core#",
        "tag": "https://brickschema.org/schema/1.1.0/BrickTag#",
        "xml": "http://www.w3.org/XML/1998/namespace",
        "xsd": "http://www.w3.org/2001/XMLSchema#"
    },
    "success": true
}
```

### Resolving a query

```shell
$ curl -X PUT -d '"SELECT ?pred ?obj WHERE {brick:Sensor ?pred ?obj .}"' http://localhost:8001/query
{
    "resultset": [
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2111"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "https://brickschema.org/schema/1.1.0/Brick#Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2646"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb906"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb769"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2394"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1822"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1487"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2730"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb508"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1920"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2398"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb313"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1188"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1795"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1395"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb513"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1608"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2666"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "https://brickschema.org/schema/1.1.0/Brick#Point"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1238"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb721"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb453"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb491"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2541"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb303"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb648"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb4"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2063"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1898"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1988"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb755"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb771"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1776"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1229"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb965"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb484"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb709"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2405"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1227"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb714"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1251"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb210"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2304"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2628"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1834"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2805"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1143"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2522"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2708"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb232"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1273"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2371"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb242"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb149"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb960"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb146"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2768"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2036"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2813"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1364"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2602"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1787"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb903"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2451"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2210"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb416"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2348"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1163"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2264"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb941"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb60"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1739"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2456"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb737"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1629"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2562"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb138"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb101"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb371"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1235"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb752"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2566"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2690"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1656"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1811"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1473"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb579"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1755"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1430"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2656"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1259"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1764"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb398"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb954"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb236"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2625"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2103"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb599"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb37"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2192"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2196"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb922"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2309"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb687"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2424"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1401"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2680"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb990"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2270"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2039"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2512"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2693"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2239"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#label",
            "Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb538"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2787"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb596"
        ],
        [
            "http://www.w3.org/2002/07/owl#sameAs",
            "https://brickschema.org/schema/1.1.0/Brick#Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2373"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1051"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2777"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1029"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "http://www.w3.org/2002/07/owl#Thing"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1858"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2756"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1367"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1529"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1150"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1095"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1625"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1865"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb701"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2751"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1910"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1994"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb520"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2669"
        ],
        [
            "http://www.w3.org/1999/02/22-rdf-syntax-ns#type",
            "http://www.w3.org/2002/07/owl#Class"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2631"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2275"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1622"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2356"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb611"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1412"
        ],
        [
            "http://www.w3.org/2002/07/owl#equivalentClass",
            "f7384c45206bb42dba21f93c452e1971eb4"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1096"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1337"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1585"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2091"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1175"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1718"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb271"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2261"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb175"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb660"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb705"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1278"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1547"
        ],
        [
            "http://www.w3.org/2002/07/owl#equivalentClass",
            "https://brickschema.org/schema/1.1.0/Brick#Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1480"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2150"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1556"
        ]
    ],
    "success": true
}

```

### Performing an update

```shell
$ curl -X PUT -d '"PREFIX brick: <https://brickschema.org/schema/1.1.0/Brick#>\n\nDELETE { brick:Sensor rdfs:subClassOf ?obj .} WHERE {brick:Sensor rdfs:subClassOf ?obj .}"' http://localhost:8001/update
{
    "success": true
}
```

**Note:** rdflib has a few restrictions:
- namespace prefixes must be used
- a where clause must exist
- no slashes in entity names

### Inspecting a query

```shell
$ curl -X PUT -d '"! SELECT ?pred ?obj WHERE {?sub ?pred ?obj .}"' http://localhost:8001/inspect
{
    "success": false,
    "error": {
        "type": "lex",
        "data": "Lexer error: could not match text starting at 1:1 failing at 1:2.\n\tunmatched text: \"!\""
    }
}
$ curl -X PUT -d '"SELECT SELECT ?pred ?obj WHERE {?sub ?pred ?obj .}"' http://localhost:8001/inspect
{
    "success": false,
    "error": {
        "type": "parse",
        "data": "syntax error"
    }
}
$ curl -X PUT -d '"SELECT ?pred ?obj WHERE {?sub ?pred ?obj .}"' http://localhost:8001/inspect
{
    "success": true,
    "tokens": [
        "57352 "SELECT" 0 (1, 1)-(1, 6)",
        "57346 "?pred" 7 (1, 8)-(1, 12)",
        "57346 "?obj" 13 (1, 14)-(1, 17)",
        "57353 "WHERE" 18 (1, 19)-(1, 23)",
        "57355 "{" 24 (1, 25)-(1, 25)",
        "57346 "?sub" 25 (1, 26)-(1, 29)",
        "57346 "?pred" 30 (1, 31)-(1, 35)",
        "57346 "?obj" 36 (1, 37)-(1, 40)",
        "57359 "." 41 (1, 42)-(1, 42)",
        "57356 "}" 42 (1, 43)-(1, 43)"
    ],
    "sexp": "(select \"SELECT\" (list \"SELECT\") (list \"?pred\" (var \"?pred\") (var \"?obj\")) (list \"?sub\" (restriction \"?sub\" (var \"?sub\") (var \"?pred\") (var \"?obj\"))))"
}
```

**Note:** This may provide useful insights while debugging.

### Send valid data over MQTT

```shell
$ mosquitto_pub -t "test" -m "{\"time\": 1234.5678, \"value\": 42.56}"
```

### Send invalid data over MQTT

```shell
$ mosquitto_pub -t "test" -m "blah"
```

## Testing

### SparQL Parsing

```shell
$ cd test
$ make parse-sparql
go run parse-sparql.go
Case: garbage
[TOKENS]
 - 57354 "garbage" 0 (1, 1)-(1, 7)
[PARSE] Error parsing: syntax error

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57352 "?var3" 45 (1, 46)-(1, 50)
 - 57361 "." 51 (1, 52)-(1, 52)
 - 57358 "}" 53 (1, 54)-(1, 54)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 #
?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 41 (2, 1)-(2, 5)
 - 57352 "?var3" 47 (2, 7)-(2, 11)
 - 57361 "." 53 (2, 13)-(2, 13)
 - 57358 "}" 55 (2, 15)-(2, 15)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 #comment
?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 48 (2, 1)-(2, 5)
 - 57352 "?var3" 54 (2, 7)-(2, 11)
 - 57361 "." 60 (2, 13)-(2, 13)
 - 57358 "}" 62 (2, 15)-(2, 15)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2) ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57367 "(" 39 (1, 40)-(1, 40)
 - 57352 "?var2" 40 (1, 41)-(1, 45)
 - 57368 ")" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 47 (1, 48)-(1, 52)
 - 57361 "." 53 (1, 54)-(1, 54)
 - 57358 "}" 55 (1, 56)-(1, 56)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2/?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57362 "/" 44 (1, 45)-(1, 45)
 - 57352 "?var2" 45 (1, 46)-(1, 50)
 - 57352 "?var3" 51 (1, 52)-(1, 56)
 - 57361 "." 57 (1, 58)-(1, 58)
 - 57358 "}" 59 (1, 60)-(1, 60)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (sequence "/" (var "?var2") (var "?var2")) (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2/?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2/?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2|?var3 ?var4 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57363 "|" 44 (1, 45)-(1, 45)
 - 57352 "?var3" 45 (1, 46)-(1, 50)
 - 57352 "?var4" 51 (1, 52)-(1, 56)
 - 57361 "." 57 (1, 58)-(1, 58)
 - 57358 "}" 59 (1, 60)-(1, 60)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (choice "|" (var "?var2") (var "?var3")) (var "?var4")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3) ?var4 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57367 "(" 39 (1, 40)-(1, 40)
 - 57352 "?var2" 40 (1, 41)-(1, 45)
 - 57363 "|" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57368 ")" 51 (1, 52)-(1, 52)
 - 57352 "?var4" 53 (1, 54)-(1, 58)
 - 57361 "." 59 (1, 60)-(1, 60)
 - 57358 "}" 61 (1, 62)-(1, 62)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (choice "|" (var "?var2") (var "?var3")) (var "?var4")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2+ ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57364 "+" 44 (1, 45)-(1, 45)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57361 "." 52 (1, 53)-(1, 53)
 - 57358 "}" 54 (1, 55)-(1, 55)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (one-or-more "+" (var "?var2")) (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2+ ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2+ ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2++ ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57364 "+" 44 (1, 45)-(1, 45)
 - 57364 "+" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 47 (1, 48)-(1, 52)
 - 57361 "." 53 (1, 54)-(1, 54)
 - 57358 "}" 55 (1, 56)-(1, 56)
[PARSE] Error parsing: syntax error

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2* ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57365 "*" 44 (1, 45)-(1, 45)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57361 "." 52 (1, 53)-(1, 53)
 - 57358 "}" 54 (1, 55)-(1, 55)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (zero-or-more "*" (var "?var2")) (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2* ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2* ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3)/?var4 ?var5 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57367 "(" 39 (1, 40)-(1, 40)
 - 57352 "?var2" 40 (1, 41)-(1, 45)
 - 57363 "|" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57368 ")" 51 (1, 52)-(1, 52)
 - 57362 "/" 52 (1, 53)-(1, 53)
 - 57352 "?var4" 53 (1, 54)-(1, 58)
 - 57352 "?var5" 59 (1, 60)-(1, 64)
 - 57361 "." 65 (1, 66)-(1, 66)
 - 57358 "}" 67 (1, 68)-(1, 68)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (sequence "/" (choice "|" (var "?var2") (var "?var3")) (var "?var4")) (var "?var5")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3)/?var4 ?var5 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3)/?var4 ?var5 .
}

Case: SELECT ?var1 ?var2 WHERE { { ?var1 ?var1 ?var2 } UNION { ?var1 ?var2 ?var2 } . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57357 "{" 27 (1, 28)-(1, 28)
 - 57352 "?var1" 29 (1, 30)-(1, 34)
 - 57352 "?var1" 35 (1, 36)-(1, 40)
 - 57352 "?var2" 41 (1, 42)-(1, 46)
 - 57358 "}" 47 (1, 48)-(1, 48)
 - 57348 "UNION" 49 (1, 50)-(1, 54)
 - 57357 "{" 55 (1, 56)-(1, 56)
 - 57352 "?var1" 57 (1, 58)-(1, 62)
 - 57352 "?var2" 63 (1, 64)-(1, 68)
 - 57352 "?var2" 69 (1, 70)-(1, 74)
 - 57358 "}" 75 (1, 76)-(1, 76)
 - 57361 "." 77 (1, 78)-(1, 78)
 - 57358 "}" 79 (1, 80)-(1, 80)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "{" (union "UNION" (restriction "?var1" (var "?var1") (var "?var1") (var "?var2")) (restriction "?var1" (var "?var1") (var "?var2") (var "?var2"))))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    {
        ?var1 ?var1 ?var2 .
    } UNION {
        ?var1 ?var2 ?var2 .
    } .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    {
        ?var1 ?var1 ?var2 .
    } UNION {
        ?var1 ?var2 ?var2 .
    } .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 < http://www.google.com#test > . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57359 "<" 39 (1, 40)-(1, 40)
 - 57353 "http://www.google.com#test" 41 (1, 42)-(1, 67)
 - 57360 ">" 68 (1, 69)-(1, 69)
 - 57361 "." 70 (1, 71)-(1, 71)
 - 57358 "}" 72 (1, 73)-(1, 73)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (uri "http://www.google.com#test")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 <http://www.google.com#test> . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57359 "<" 39 (1, 40)-(1, 40)
 - 57353 "http://www.google.com#test" 40 (1, 41)-(1, 66)
 - 57360 ">" 66 (1, 67)-(1, 67)
 - 57361 "." 68 (1, 69)-(1, 69)
 - 57358 "}" 70 (1, 71)-(1, 71)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (uri "http://www.google.com#test")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 "" . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57355 "\"\"" 39 (1, 40)-(1, 41)
 - 57361 "." 42 (1, 43)-(1, 43)
 - 57358 "}" 44 (1, 45)-(1, 45)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (string "\"\"")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "" .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "" .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 "a b" . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57355 "\"a b\"" 39 (1, 40)-(1, 44)
 - 57361 "." 45 (1, 46)-(1, 46)
 - 57358 "}" 47 (1, 48)-(1, 48)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (string "\"a b\"")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a b" .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a b" .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 "a \" b" . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57355 "\"a \\\" b\"" 39 (1, 40)-(1, 47)
 - 57361 "." 48 (1, 49)-(1, 49)
 - 57358 "}" 50 (1, 51)-(1, 51)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (string "\"a \\\" b\"")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a \" b" .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a \" b" .
}

Case: PREFIX a : <http://b> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 9 (1, 10)-(1, 10)
 - 57359 "<" 11 (1, 12)-(1, 12)
 - 57353 "http://b" 12 (1, 13)-(1, 20)
 - 57360 ">" 20 (1, 21)-(1, 21)
 - 57346 "SELECT" 22 (1, 23)-(1, 28)
 - 57352 "?var1" 29 (1, 30)-(1, 34)
 - 57352 "?var2" 35 (1, 36)-(1, 40)
 - 57347 "WHERE" 41 (1, 42)-(1, 46)
 - 57357 "{" 47 (1, 48)-(1, 48)
 - 57352 "?var1" 49 (1, 50)-(1, 54)
 - 57352 "?var2" 55 (1, 56)-(1, 60)
 - 57352 "?var3" 61 (1, 62)-(1, 66)
 - 57361 "." 67 (1, 68)-(1, 68)
 - 57358 "}" 69 (1, 70)-(1, 70)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "PREFIX") (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57346 "SELECT" 20 (1, 21)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57347 "WHERE" 39 (1, 40)-(1, 44)
 - 57357 "{" 45 (1, 46)-(1, 46)
 - 57352 "?var1" 47 (1, 48)-(1, 52)
 - 57352 "?var2" 53 (1, 54)-(1, 58)
 - 57352 "?var3" 59 (1, 60)-(1, 64)
 - 57361 "." 65 (1, 66)-(1, 66)
 - 57358 "}" 67 (1, 68)-(1, 68)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "PREFIX") (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> PREFIX c:<http://d> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57351 "PREFIX" 20 (1, 21)-(1, 26)
 - 57354 "c" 27 (1, 28)-(1, 28)
 - 57366 ":" 28 (1, 29)-(1, 29)
 - 57359 "<" 29 (1, 30)-(1, 30)
 - 57353 "http://d" 30 (1, 31)-(1, 38)
 - 57360 ">" 38 (1, 39)-(1, 39)
 - 57346 "SELECT" 40 (1, 41)-(1, 46)
 - 57352 "?var1" 47 (1, 48)-(1, 52)
 - 57352 "?var2" 53 (1, 54)-(1, 58)
 - 57347 "WHERE" 59 (1, 60)-(1, 64)
 - 57357 "{" 65 (1, 66)-(1, 66)
 - 57352 "?var1" 67 (1, 68)-(1, 72)
 - 57352 "?var2" 73 (1, 74)-(1, 78)
 - 57352 "?var3" 79 (1, 80)-(1, 84)
 - 57361 "." 85 (1, 86)-(1, 86)
 - 57358 "}" 87 (1, 88)-(1, 88)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b")) (prefix "PREFIX" (id "c") (uri "http://d"))) (list "PREFIX") (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>
PREFIX c: <http://d>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>
PREFIX c: <http://d>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { a:b ?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57354 "a" 33 (1, 34)-(1, 34)
 - 57366 ":" 34 (1, 35)-(1, 35)
 - 57354 "b" 35 (1, 36)-(1, 36)
 - 57352 "?var2" 37 (1, 38)-(1, 42)
 - 57352 "?var3" 43 (1, 44)-(1, 48)
 - 57361 "." 49 (1, 50)-(1, 50)
 - 57358 "}" 51 (1, 52)-(1, 52)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "a" (restriction "a" (prefixed ":" (id "a") (id "b")) (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    a:b ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    a:b ?var2 ?var3 .
}

Case: DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57349 "DATA" 0 (1, 1)-(1, 4)
 - 57352 "?var1" 5 (1, 6)-(1, 10)
 - 57352 "?var3" 11 (1, 12)-(1, 16)
 - 57346 "SELECT" 17 (1, 18)-(1, 23)
 - 57352 "?var1" 24 (1, 25)-(1, 29)
 - 57352 "?var2" 30 (1, 31)-(1, 35)
 - 57352 "?var3" 36 (1, 37)-(1, 41)
 - 57347 "WHERE" 42 (1, 43)-(1, 47)
 - 57357 "{" 48 (1, 49)-(1, 49)
 - 57352 "?var1" 50 (1, 51)-(1, 55)
 - 57352 "?var2" 56 (1, 57)-(1, 61)
 - 57352 "?var3" 62 (1, 63)-(1, 67)
 - 57361 "." 68 (1, 69)-(1, 69)
 - 57358 "}" 70 (1, 71)-(1, 71)
[PARSE] (query "DATA" (list "DATA") (list "?var1" (var "?var1") (var "?var3")) (list "DATA") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
DATA
    ?var1
    ?var3
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57349 "DATA" 20 (1, 21)-(1, 24)
 - 57352 "?var1" 25 (1, 26)-(1, 30)
 - 57352 "?var3" 31 (1, 32)-(1, 36)
 - 57346 "SELECT" 37 (1, 38)-(1, 43)
 - 57352 "?var1" 44 (1, 45)-(1, 49)
 - 57352 "?var2" 50 (1, 51)-(1, 55)
 - 57352 "?var3" 56 (1, 57)-(1, 61)
 - 57347 "WHERE" 62 (1, 63)-(1, 67)
 - 57357 "{" 68 (1, 69)-(1, 69)
 - 57352 "?var1" 70 (1, 71)-(1, 75)
 - 57352 "?var2" 76 (1, 77)-(1, 81)
 - 57352 "?var3" 82 (1, 83)-(1, 87)
 - 57361 "." 88 (1, 89)-(1, 89)
 - 57358 "}" 90 (1, 91)-(1, 91)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "?var1" (var "?var1") (var "?var3")) (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

DATA
    ?var1
    ?var3
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57349 "DATA" 20 (1, 21)-(1, 24)
 - 57352 "?var1" 25 (1, 26)-(1, 30)
 - 57352 "?var3" 31 (1, 32)-(1, 36)
 - 57346 "SELECT" 37 (1, 38)-(1, 43)
 - 57352 "?var1" 44 (1, 45)-(1, 49)
 - 57352 "?var2" 50 (1, 51)-(1, 55)
 - 57352 "?var3" 56 (1, 57)-(1, 61)
 - 57347 "WHERE" 62 (1, 63)-(1, 67)
 - 57357 "{" 68 (1, 69)-(1, 69)
 - 57352 "?var1" 70 (1, 71)-(1, 75)
 - 57352 "?var2" 76 (1, 77)-(1, 81)
 - 57352 "?var3" 82 (1, 83)-(1, 87)
 - 57361 "." 88 (1, 89)-(1, 89)
 - 57358 "}" 90 (1, 91)-(1, 91)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "?var1" (var "?var1") (var "?var3")) (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

DATA
    ?var1
    ?var3
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> DATA ?var1 ?var3 UNITS mod:temp->unit:degc mod:dist->unit:m SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57349 "DATA" 20 (1, 21)-(1, 24)
 - 57352 "?var1" 25 (1, 26)-(1, 30)
 - 57352 "?var3" 31 (1, 32)-(1, 36)
 - 57350 "UNITS" 37 (1, 38)-(1, 42)
 - 57354 "mod" 43 (1, 44)-(1, 46)
 - 57366 ":" 46 (1, 47)-(1, 47)
 - 57354 "temp" 47 (1, 48)-(1, 51)
 - 57356 "->" 51 (1, 52)-(1, 53)
 - 57354 "unit" 53 (1, 54)-(1, 57)
 - 57366 ":" 57 (1, 58)-(1, 58)
 - 57354 "degc" 58 (1, 59)-(1, 62)
 - 57354 "mod" 63 (1, 64)-(1, 66)
 - 57366 ":" 66 (1, 67)-(1, 67)
 - 57354 "dist" 67 (1, 68)-(1, 71)
 - 57356 "->" 71 (1, 72)-(1, 73)
 - 57354 "unit" 73 (1, 74)-(1, 77)
 - 57366 ":" 77 (1, 78)-(1, 78)
 - 57354 "m" 78 (1, 79)-(1, 79)
 - 57346 "SELECT" 80 (1, 81)-(1, 86)
 - 57352 "?var1" 87 (1, 88)-(1, 92)
 - 57352 "?var2" 93 (1, 94)-(1, 98)
 - 57352 "?var3" 99 (1, 100)-(1, 104)
 - 57347 "WHERE" 105 (1, 106)-(1, 110)
 - 57357 "{" 111 (1, 112)-(1, 112)
 - 57352 "?var1" 113 (1, 114)-(1, 118)
 - 57352 "?var2" 119 (1, 120)-(1, 124)
 - 57352 "?var3" 125 (1, 126)-(1, 130)
 - 57361 "." 131 (1, 132)-(1, 132)
 - 57358 "}" 133 (1, 134)-(1, 134)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "?var1" (var "?var1") (var "?var3")) (list "mod" (mapping "->" (prefixed ":" (id "mod") (id "temp")) (prefixed ":" (id "unit") (id "degc"))) (mapping "->" (prefixed ":" (id "mod") (id "dist")) (prefixed ":" (id "unit") (id "m")))) (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

DATA
    ?var1
    ?var3
UNITS
    mod:temp -> unit:degc
    mod:dist -> unit:m
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
```

### Data Dispatching

```shell
$ cd test
$ make data-dispatch-registration
go run data-dispatch-registration.go
Registering:
~~~~~~~~~~~
Dispatcher
 - a: 4 entries
 - b: 3 entries
 - c: 2 entries
 - d: 1 entries

Unregistering (not registered):
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - c: 2 entries
 - d: 1 entries
 - a: 4 entries
 - b: 3 entries

Unregistering (beginning):
~~~~~~~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - b: 3 entries
 - c: 2 entries
 - d: 1 entries
 - a: 3 entries

Unregistering (middle):
~~~~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - a: 2 entries
 - b: 3 entries
 - c: 2 entries
 - d: 1 entries

Unregistering (end):
~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - c: 2 entries
 - d: 1 entries
 - a: 2 entries
 - b: 2 entries
```

