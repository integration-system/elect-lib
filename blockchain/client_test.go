package blockchain

import (
	"github.com/valyala/fasthttp"
	"sync"
	"sync/atomic"
	"testing"
)

type stubTransport struct {
	invokeCounter int32
	authCount     int32
}

func (t *stubTransport) Invoke(uri string, header map[string]string, request []byte, respPtr interface{}) (int, error) {
	if uri == authenticateMethod {
		atomic.AddInt32(&t.authCount, 1)
		return 0, json.Unmarshal([]byte(`{"result": {"token": "123"}}`), respPtr)
	} else {
		val := atomic.AddInt32(&t.invokeCounter, 1)
		switch val {
		case 4, 7:
			return fasthttp.StatusUnauthorized, nil
		default:
			return 0, json.Unmarshal([]byte(`{"result": {"issuedFor": {"voterId": "123", "votingId": 123}}}`), respPtr)
		}
	}
}

func TestClient_IssueBallot(t *testing.T) {
	cli := NewClient(Config{
		Login: Login{
			Login:    "test",
			Password: "test",
		},
		Address: "test.com",
	})
	trans := &stubTransport{}
	cli.transport = trans

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ballot, err := cli.IssueBallot(IssueBallotRequest{})
			if err != nil {
				t.Fatal(err)
			}
			if ballot.IssuedFor.VoterId != "123" || ballot.IssuedFor.VotingId != 123 {
				t.Fatalf("received unexpected response %v", ballot)
			}
		}()
	}
	wg.Wait()

	if trans.authCount != 3 {
		t.Fatalf("ath count %d != %d", trans.authCount, 3)
	}
}
