/*
 * Copyright (c) 2019-2020 Datacequia LLC. All rights reserved.
 *
 * This program is licensed to you under the Apache License Version 2.0,
 * and you may not use this file except in compliance with the Apache License Version 2.0.
 * You may obtain a copy of the Apache License Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Apache License Version 2.0 is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the Apache License Version 2.0 for the specific language governing permissions and limitations there under.
 */

package repo

import (
	"context"

	rescom "github.com/datacequia/go-dogg3rz/resource/common"
)

// RepositoryResource is an interface the provides all the non-iterative interactions
// a user can perform against a repository
type RepositoryResource interface {
	InitRepo(ctxt context.Context, repoName string) error

 
	CreateDataset(ctxt context.Context, repoName string, datasetPath string) error

	AddNamespaceDataset(ctxt context.Context, repoName string, datasetPath string, term string, iri string) error

	AddNamespaceNode(ctxt context.Context, repoName string, datasetPath string, nodeID string, term string, iri string) error

	CreateSnapshot(ctxt context.Context, repoName string) error

	CreateTypeClass(ctxt context.Context, repoName string, datasetPath string, typeID string, subclassOf string,
		label string, comment string) error

	CreateTypeDatatype(ctxt context.Context, repoName string, datasetPath string, typeID string, subclassOf string,
		label string, comment string) error

	CreateTypeProperty(ctxt context.Context, repoName string, datasetPath string, typeID string, subPropertyOf string,
		domain string, _range string, label string, comment string) error

	// INSERT NEW NODE INTO DEFAULT-GRAPH (graphName="") or NAMED-GRAPH (graphName!="")
	// FOR DATASET datasetPath REPO repoName
	InsertNode(ctxt context.Context, repoName string, datasetPath string,
		nodeType string, nodeID string, graphName string,
		nodeProperties []string, nodeValues []string) error
}

type GetResourceItem interface {
	GetPath() string // PATH TO RESOURCE
	GetStatus() string
}

// ALLOWS USER TO STAGE (i.e. .Add() )  EXISTING (JSON-LD) WORKSPACE RESOURCES ITERATIVELY
// TO REPOSITORY IDENTIFIED BY VALUE RETURNED FROM  .Repository()
// BEFORE FLUSHING STAGED RESOURCES USING .Commit()  
//
// NOTE: CALLER SHOULD PASS THE SAME CANCELLABLE (context.Context.WithCancel()) CONTEXT OBJECT INSTANCE
// TO ResourceStager METHODS WHICH REQUIRE A CONTEXT AND CALL THE CANCEL FUNCTION WHEN DONE TO ENSURE
// ANY ALLOCATED GO-ROUTINES ALLOCATED DURING THE INTERACTION WITH THIS INTERFACE ARE DEALLOCATED

type ResourceStager interface {
	Add(ctxt context.Context, sr rescom.StagingResource)  error
        Commit(ctxt context.Context)  error 
        Repository() string 
} 


