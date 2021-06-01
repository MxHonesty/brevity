package dependency
import (
	"brevity/task"
	"testing"
	"time"
)

// Test case for making sure the tasks are copied.
func TestCreateDependency(t *testing.T) {
	tm1 := time.Now()
	tm2 := tm1.Add(10)

	con1 := task.NewContainer(0, tm1, tm1)
	con2 := task.NewContainer(1, tm1, tm1)
	con3 := task.NewContainer(2, tm1, tm1)

	dep := NewDependency([]task.Scheduable{con1, con2}, con3, 0)
	con1.SetStartTime(tm2)
	if dep.dependentOn[0].GetStartTime() == tm2 {
		t.Error("Element not copied")
	}
}

// Test case for making sure the setters
// are making copies of the tasks.
func TestSettersDepOnDependency(t *testing.T) {
	tm1 := time.Now()
	tm2 := tm1.Add(10)

	con1 := task.NewContainer(0, tm1, tm1)
	con2 := task.NewContainer(1, tm1, tm1)
	con3 := task.NewContainer(2, tm1, tm1)

	dep := NewDependency([]task.Scheduable{con1, con2}, con3, 0)
	dep.SetDependent(con1)
	if dep.Dependent().GetId() != 0 {
		t.Errorf("Setter failed, expected id 0, got %d", dep.Dependent().GetId())
	}

	dep.Dependent().SetEndTime(tm2)  // Set the end time.
	if dep.Dependent().GetEndTime() != tm2 {
		t.Error("End time not set correctly")
	}

	con4 := task.NewContainer(4, tm2, tm2)
	dep.SetDependentOn([]task.Scheduable{con4})
	// Testing the new list was added.
	if len(dep.DependentOn()) != 1 {
		t.Errorf("Expected 1 item, got %d", len(dep.DependentOn()))
	}

	// Testing that the getter gives us the elements by ref
	tm4 := time.Now().Add(100)
	elems := dep.DependentOn()
	elems[0].SetEndTime(tm4)

	elems = dep.DependentOn()
	if elems[0].GetEndTime() != tm4 {
		t.Error("Expected EndTime to be updated.")
	}
}

// Test case for making sure that the setter for
// the dependent element creates a copy of that element.
func TestSetDependentCreatesCopy(t *testing.T) {
	tm1 := time.Now()
	tm2 := tm1.Add(10)

	con1 := task.NewContainer(0, tm1, tm1)
	con2 := task.NewContainer(1, tm1, tm1)
	con3 := task.NewContainer(2, tm1, tm1)

	dep := NewDependency([]task.Scheduable{con1, con2}, con3, 0)
	dep.Dependent().SetEndTime(tm2)
	if dep.Dependent().GetEndTime() != tm2 {
		t.Error("Expected end time to be different")
	}

	con4 := task.NewContainer(3, tm1, tm1)
	dep.SetDependent(con4)
	// Make sure it created a copy.
	con4.SetStartTime(tm2)
	if dep.Dependent().GetStartTime() != tm1 {
		t.Error("Expected StartTime not to change")
	}
}

// Test case for making a copy of the dependency.
func TestCopyDependency(t *testing.T) {
	tm1 := time.Now()
	tm2 := tm1.Add(10)

	con1 := task.NewContainer(0, tm1, tm1)
	con2 := task.NewContainer(1, tm1, tm1)
	con3 := task.NewContainer(2, tm1, tm1)

	dep := NewDependency([]task.Scheduable{con1, con2}, con3, 0)
	dep2 := dep.Copy()

	// Change Every attribute of dep and see if dep2 also changes.
	dep.SetDependentOn([]task.Scheduable{con1, con1, con1})
	if len(dep2.DependentOn()) != 2 {
		t.Errorf("Expected size to be 2, got %d", len(dep2.DependentOn()))
	}

	dep.Dependent().SetStartTime(tm2)
	if dep2.Dependent().GetStartTime() != tm1 {
		t.Error("Expected StartTime of copy not to change")
	}

	if dep.GetId() != dep2.GetId() {
		t.Error("Expected copy to have the same id")
	}
}
