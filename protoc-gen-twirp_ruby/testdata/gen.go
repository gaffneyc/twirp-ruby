// Copyright 2018 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package testdata

//go:generate protoc --descriptor_set_out=fileset.pb --include_imports --include_source_info ./rubytypes.proto
//go:generate protoc --descriptor_set_out=ruby_package.pb --include_imports --include_source_info ./ruby_package.proto
