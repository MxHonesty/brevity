package dependency

import (
	"brevity/task"
	"testing"
	"time"
)

// Test case for adding dependentOn task.Scheduable.
func TestConcreteBuilder_AddDependentOn(t *testing.T) {
	builder := NewConcreteBuilder()
	builder.AddDependentOn(task.NewContainer(0, time.Now(), time.Now()))
	builder.AddDependentOn(task.NewContainer(1, time.Now(), time.Now()))

	dep := builder.GetResult(0)
	if dep.GetId() != 0 {
		t.Errorf("Expected Dependency id to be 0, got %d", dep.GetId())
	}

	if len(dep.DependentOn()) != 2 {
		t.Errorf("Expected the size of dependentOn to be 2, got %d", len(dep.DependentOn()))
	}

	firstId := dep.DependentOn()[0].GetId()
	secondId := dep.DependentOn()[1].GetId()

	if firstId != 0 {
		t.Errorf("Expected the id of first task.Scheduable to be 0, got %d.", firstId)
	}
	if secondId != 1 {
		t.Errorf("Expected the id of first task.Scheduable to be 0, got %d.", secondId)
	}
}

// Test case for removing task.Scheduable from dependentOn.
func TestConcreteBuilder_RemoveDependentOn(t *testing.T) {
	builder := NewConcreteBuilder()
	builder.AddDependentOn(task.NewContainer(0, time.Now(), time.Now()))
	builder.AddDependentOn(task.NewContainer(1, time.Now(), time.Now()))
	builder.AddDependentOn(task.NewContainer(2, time.Now(), time.Now()))
	builder.RemoveDependentOn(1)  // Remove the second dependentOn

	dep := builder.GetResult(0)
	size := len(dep.DependentOn())
	if size != 2 {
		t.Errorf("Expected to remove an item and get size 2, got %d", size)
	}

	// Check if the correct item was removed
	firstId := dep.DependentOn()[0].GetId()
	secondId := dep.DependentOn()[1].GetId()
	if firstId != 0 {
		t.Errorf("Expected item to have id 0, got %d", firstId)
	}
	if secondId != 2 {
		t.Errorf("Expected item to have id 0, got %d", secondId)
	}
}

// Test case for setting the dependent task.Scheduable
// Test setting the Dependent once.
// Test overriding a current Dependent.
func TestConcreteBuilder_SetDependent(t *testing.T) {
	builder := NewConcreteBuilder()
	builder.SetDependent(task.NewContainer(0, time.Now(), time.Now()))
	dep := builder.GetResult(0)

	if dep.Dependent().GetId() != 0 {
		t.Errorf("Expected the id of container to be 0, got %d", dep.Dependent().GetId())
	}

	builder.SetDependent(task.NewContainer(0, time.Now(), time.Now()))
	// Override the Dependent
	builder.SetDependent(task.NewContainer(1, time.Now(), time.Now()))
	dep = builder.GetResult(0)

	if dep.Dependent().GetId() != 1 {
		t.Errorf("Expected id to be 1, got %d", dep.Dependent().GetId())
	}
}

// Test case for GetResult method.
// It checks that the build property of the builder is
// reset after the call.
func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()
	builder.AddDependentOn(task.NewContainer(0, time.Now(), time.Now()))
	builder.AddDependentOn(task.NewContainer(1, time.Now(), time.Now()))

	dep := builder.GetResult(0)  // Should clear the builder.built
	dep = builder.GetResult(0)  // dep should be an uninitialized Dependency.

	size := len(dep.DependentOn())
	if size != 0 {
		t.Errorf("Expected Dependency to be empty, got size %d", size)
	}
}

// Test case for trying to remove an invalid id from the list
// of task.Scheduable.
func TestConcreteBuilder_RemoveDependentOnInvalidId(t *testing.T) {
	builder := NewConcreteBuilder()
	builder.AddDependentOn(task.NewContainer(0, time.Now(), time.Now()))
	builder.AddDependentOn(task.NewContainer(1, time.Now(), time.Now()))

	builder.RemoveDependentOn(2)  // Nothing should change
	builder.RemoveDependentOn(3)  // Nothing should change
	builder.RemoveDependentOn(4)  // Nothing should change
	builder.RemoveDependentOn(5)  // Nothing should change

	dep := builder.GetResult(0)
	size := len(dep.DependentOn())
	if size != 2 {
		t.Errorf("Expected size to be 2, got %d", size)
	}
}
