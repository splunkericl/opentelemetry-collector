// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package pdata

import "go.opentelemetry.io/collector/model/internal/pdata"

// InstrumentationLibrary is a message representing the instrumentation library information.
// Must use NewInstrumentationLibrary function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrary = pdata.InstrumentationLibrary

// NewInstrumentationLibrary creates a new empty InstrumentationLibrary.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
var NewInstrumentationLibrary = pdata.NewInstrumentationLibrary

// AttributeValueSlice logically represents a slice of AttributeValue.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewAttributeValueSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type AttributeValueSlice = pdata.AttributeValueSlice

// NewAttributeValueSlice creates a AttributeValueSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewAttributeValueSlice = pdata.NewAttributeValueSlice
