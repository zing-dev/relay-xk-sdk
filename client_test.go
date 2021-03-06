package relay

import (
	"fmt"
	"testing"
	"time"
)

var client = NewDefaultClient(handle)

func TestNewClient(t *testing.T) {
	_ = client.OnAll()
	time.Sleep(time.Second)
	_ = client.OffAll()
	time.Sleep(time.Second)
	_ = client.OnAll()
}

func TestClient_OffOne(t *testing.T) {
	client = NewClient(handle, DefaultBranchesLength)
	fmt.Println(client.OnAll())
	fmt.Println(client.OffOne(1))
	fmt.Println(client.OffOne(2))
	fmt.Println(client.OffOne(3))
	fmt.Println(client.OffOne(4))
	fmt.Println(client.OffOne(5))
	fmt.Println(client.OffOne(6))
	fmt.Println(client.OffOne(7))
	fmt.Println(client.OffOne(8))
	fmt.Println(client.OffOne(32))
}

func TestClient_OnOne(t *testing.T) {
	fmt.Println(client.OffAll())
	fmt.Println(client.OnOne(1))
	fmt.Println(client.OnOne(2))
	fmt.Println(client.OnOne(4))
	fmt.Println(client.OnOne(6))
	fmt.Println(client.OnOne(8))
	fmt.Println(client.OnOne(9))
	fmt.Println(client.OnOne(10))
}

func TestClient_FlipOne(t *testing.T) {
	fmt.Println(client.FlipOne(1))
	fmt.Println(client.StatusOne(1))
	fmt.Println(client.FlipOne(7))
	fmt.Println(client.StatusOne(7))
	fmt.Println(client.FlipOne(8))
	fmt.Println(client.StatusOne(8))
}

func TestClient_Status(t *testing.T) {
	fmt.Println(client.OffAll())
	fmt.Println(client.OnOne(1))
	fmt.Println(client.OnOne(7))
	fmt.Println(client.OnOne(8))
	fmt.Println(client.OnOne(10))
	fmt.Println(client.Status())
	fmt.Println(client.StatusOne(1))
	fmt.Println(client.StatusOne(6))
	fmt.Println(client.StatusOne(7))
	fmt.Println(client.StatusOne(8))
}

func TestClient_StatusOne(t *testing.T) {
	client.OffAll()
	fmt.Println(client.OnOne(0))
	fmt.Println(client.OnOne(1))
	fmt.Println(client.OnOne(2))
	fmt.Println(client.OnOne(7))
	fmt.Println(client.OnOne(8))
	fmt.Println(client.StatusOne(0))
	fmt.Println(client.StatusOne(1))
	fmt.Println(client.StatusOne(2))
	fmt.Println(client.StatusOne(7))
	fmt.Println(client.StatusOne(8))
}

func TestClient_OffGroup(t *testing.T) {
	client.OnAll()
	fmt.Println(client.OffGroup(0, 2, 4, 6))
	fmt.Println(client.OffGroup(0))
	fmt.Println(client.OffGroup(8))
}

func TestClient_OnGroup(t *testing.T) {
	client.OffAll()
	fmt.Println(client.OnGroup(0, 2, 4, 6))
	fmt.Println(client.OnGroup(0))
	fmt.Println(client.OnGroup(8))
}

func TestClient_FlipGroup(t *testing.T) {
	fmt.Println(client.FlipGroup(0, 2, 4, 6))
	fmt.Println(client.StatusOne(0))
	fmt.Println(client.StatusOne(2))
	fmt.Println(client.StatusOne(4))
	fmt.Println(client.StatusOne(6))
}

func TestClient_OffPoint(t *testing.T) {
	for i := 0; i < 9; i++ {
		fmt.Println(client.OffPoint(byte(i), 1000*(i+1)))
	}
}

func TestClient_OnPoint(t *testing.T) {
	for i := 0; i < 8; i++ {
		fmt.Println(client.OnPoint(byte(i), 1000*(i+1)))
	}
}

func TestClient_FlipOneNil(t *testing.T) {
	fmt.Println(client.FlipOneNil(1), client.FlipOneNil(2))
}

func TestClient_OffOneNil(t *testing.T) {
	fmt.Println(client.OnOne(1), client.OnOne(2))
	fmt.Println(client.OffOne(1), client.OffOne(2))
}

func TestGroupNil(t *testing.T) {
	client.FlipGroupNil(0, 1, 2, 3, 4, 5, 6, 7)
	time.Sleep(time.Second)
	client.OffGroup(0, 1, 2, 3, 4, 5, 6, 7)
	time.Sleep(time.Second)
	client.OnGroup(0, 1, 2, 3, 4, 5, 6, 7)
	time.Sleep(time.Second)
	client.OffGroup(0, 1, 2, 3, 4, 5, 6, 7)
	time.Sleep(time.Second)
	client.FlipGroupNil(0, 1, 2, 3, 4, 5, 6, 7)
}

func TestPointNil(t *testing.T) {
	fmt.Println(client.OffPoint(0, 1000))
	fmt.Println(client.OffPoint(1, 1000))
	fmt.Println(client.OffPoint(2, 2000))
	fmt.Println(client.OffPoint(8, 2000))
	time.Sleep(time.Second)
	fmt.Println(client.OnPoint(0, 1000))
	fmt.Println(client.OnPoint(1, 1000))
	fmt.Println(client.OnPoint(2, 2000))
	fmt.Println(client.OnPoint(8, 2000))
}
