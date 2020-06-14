// -------------------------
// Decoupling With Interface
// -------------------------

// By looking at the API (functions), we need to decouple the API from the concrete implementation. The decoupling
// that we do must get all the way down into initialization. To do this right, the only piece of
// code that we need to change is initialization. Everything else should be able to act on the
// behavior that these types are gonna provide.

// pull is based on the concrete. It only knows how to work on Xenia. However, if we are able to
// decouple pull to use any system that know how to pull data, we can get the highest level of
// decoupling. Since the algorithm we have is already efficient, we don't need to add another level
// of generalization and destroy the work we did in the concrete. Same thing with store.

// It is nice to work from the concrete up. When we do this, not only we are solving problem
// efficiently and reducing technical debt but the contracts, they come to us. We already know what
// the contract is for pulling/storing data. We already validate that and this is what we need.

// Let's just decouple these 2 functions and add 2 interfaces. The Puller interface knows how to
// pull and the Storer knows how to store.
// Xenia already implemented the Puller interface and Pillar already implemented the Storer
// interface. Now we can come into pull/store, decouple this function from the concrete.
// Instead of passing Xenial and Pillar, we pass in the Puller and Storer. The algorithm doesn't
// change. All we doing is now calling pull/store indirectly through the interface value.

// Next step:
// ----------
// Copy also doesn't have to change because Xenia/Pillar already implemented the interfaces.
// However, we are not done because Copy is still bounded to the concrete. Copy can only work with
// pointer of type system. We need to decouple Copy so we can have a decoupled system that knows
// how to pull and store. We will do it in the next file.

package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Code will run on a timer. It needs to connect to Xenia, read the DB, identify all the data
// we haven't moved and pull it in.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data is the structure of the data we are copying.
// For simplicity, just pretend it is a string data.
type Data struct {
	Line string
}

// Puller declares behavior for pulling data
type Puller interface {
	Pull(d *Data) error
}

// Storer declares behavior for storing data
type Storer interface {
	Store(d *Data) error
}

// Xenia is a system we need to pull data from
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Pull knows how to pull data out of Xenia.
// We could do func (*Xenia) Pull() (*Data, error) that return the data and error. However, this
// would cost an allocation on every call and we don't want that.
// Using the function below, we know data is a struct type and its size ahead of time. Therefore
// they could be on the stack.
func (x *Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}

}

// Pillar is a system we need to store data into.
type Pillar struct {
	Host    string
	Timeout time.Duration
}

// Store knows how to store data into Pillar
// We are using pointer semantics for consistency
func (p *Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// System wraps Xenia and Pillar into a single system.
// We have the API based on Xenia and Pillar. We want to build another API on top of this
// and use it as it's foundation.
// One way is to have a type that have the behavior of being able to pull and store.
// We cannot do that through composition. System is based on the embedded value of Xenia
// and Pillar. And because of inner type promotion, System knows how to pull and store.
type System struct {
	Xenia
	Pillar
}

// pull knows how to pull bulks of data from any Puller
// decoupled
func pull(x Puller, data []Data) (int, error) {
	// Range over the slice of data and store each element with the Xenia's Pull method
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store knows how to store bulks of data from any Storer.
// decoupled
func store(p Storer, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy knows how to pull and store data from the System.
// Now we can call the pull and store functions, passing Xenia and Pillar though
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}

}

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8080",
			Timeout: time.Second,
		},

		Pillar: Pillar{
			Host:    "localhost:9090",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != nil {
		fmt.Println(err)
	}
}
