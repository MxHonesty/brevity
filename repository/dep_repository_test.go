package repository
import (
	"brevity/dependency"
	"brevity/task"
	"testing"
	"time"
)

// Test case for adding dependency.Dependency.
func TestAddDepRepository(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())
	con2 := task.NewContainer(2, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con2, 0)
	repo.Add(dep)

	// Test size
	if repo.Size() != 1 {
		t.Errorf("Expected size 1, got %d", repo.Size())
	}

	// Test it's the same element by value
	el, err := repo.Retrieve(0)
	if err != nil {
		t.Error("Did not find exisint item")
	} else {
		if el.GetId() != dep.GetId() {
			t.Errorf("Expected %d, got %d", dep.GetId(), el.GetId())
		}
	}

	// Make sure it's a copy
	dep.SetDependent(con1)
	el, err = repo.Retrieve(0)
	if err != nil {
		t.Error("Did not find existing item")
	} else {
		elDep := el.Dependent()
		if elDep.GetId() == con1.GetId() {
			t.Error("Expected to have different ids")
		}
	}
}

// Test case for removing dependency.Dependency.
func TestRemoveDepRepository(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())
	con2 := task.NewContainer(2, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con2, 0)
	dep2 := dependency.NewDependency([]task.Scheduable{con1}, con2, 1)
	repo.Add(dep)
	repo.Add(dep2)

	err := repo.Remove(0)
	if err != nil {
		t.Error("Expected no error")
	}
	err = repo.Remove(0)  // Try to remove same id
	if err == nil {
		t.Error("Expected error when deleting not existing item")
	}

	if repo.Size() != 1 {
		t.Errorf("Expected size 1, got %d", repo.Size())
	}

	found := repo.Find(0)  // Check if deleted item is found
	if found {
		t.Error("Found deleted item")
	}

	found = repo.Find(1)
	if !found {
		t.Error("Did not find exiting item")
	}
}

// Test case for taking the list of elements.
func TestGetAllRepository(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())
	con2 := task.NewContainer(2, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con2, 0)
	dep2 := dependency.NewDependency([]task.Scheduable{con1}, con2, 1)
	repo.Add(dep)
	repo.Add(dep2)

	list := repo.GetAll()
	values := list.Values()
	depRet1 := values[0].(dependency.Dependency)
	depRet2 := values[1].(dependency.Dependency)

	if depRet1.GetId() != dep.GetId() {
		t.Errorf("Expected it %d, got %d", dep.GetId(), depRet1.GetId())
	}

	if depRet2.GetId() != dep2.GetId() {
		t.Errorf("Expected it %d, got %d", dep2.GetId(), depRet2.GetId())
	}
}

// Test the update operation on DepRepository.
func TestUpdateDepRepository(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())
	con2 := task.NewContainer(2, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con1, 0)
	dep2 := dependency.NewDependency([]task.Scheduable{con1}, con2, 1)
	repo.Add(dep)

	// We expect id not to change and dependent to change
	err := repo.Update(0, dep2)
	if err != nil {
		t.Error("Expected no error")
	}

	el, err2 := repo.Retrieve(0)
	if err2 != nil {
		t.Error("Expected to find the same id")
	} else {
		depCon := el.Dependent()
		if depCon.GetId() != con2.GetId() {
			t.Error("Expected container to be the same")
		}
	}
}

func TestUpdateNotExistingDepRepository(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con1, 0)
	repo.Add(dep)

	err := repo.Update(10, dep)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestDeleteNonExisting(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())
	con2 := task.NewContainer(2, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con2, 0)
	dep2 := dependency.NewDependency([]task.Scheduable{con1}, con2, 1)
	repo.Add(dep)
	repo.Add(dep2)

	err := repo.Remove(0)
	err = repo.Remove(0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test case for removing all
func TestRemoveAllDepRepository(t *testing.T) {
	repo := NewDepRepository()
	con1 := task.NewContainer(0, time.Now(), time.Now())
	con2 := task.NewContainer(2, time.Now(), time.Now())

	dep := dependency.NewDependency([]task.Scheduable{con1}, con2, 0)
	dep2 := dependency.NewDependency([]task.Scheduable{con1}, con2, 1)
	repo.Add(dep)
	repo.Add(dep2)

	repo.RemoveAll()
	if repo.Size() != 0 {
		t.Errorf("Expected size 0, got %d", repo.Size())
	}
}

func TestRetrieveRemoveNotExistingItemDepRepository(t *testing.T) {
	repo := NewDepRepository()
	err := repo.Remove(0)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	_, err = repo.Retrieve(0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
