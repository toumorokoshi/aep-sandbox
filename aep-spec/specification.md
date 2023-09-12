# API Extention Proposal Specification

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [BCP 14](https://tools.ietf.org/html/bcp14) [RFC2119](https://tools.ietf.org/html/rfc2119) [RFC8174](https://tools.ietf.org/html/rfc8174) when, and only when, they appear in all capitals, as shown here.

This document is licensed under [The Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0.html).

## Introduction

The AEP Specification defines a set of requirements around HTTP APIs which
adhere to resource-oriented design principles.

APIs that adhere to the specification are highly uniform, thereby enabling
low-cost, high-quality integrations for multiple clients including:

- command-line interfaces
- visual user interfaces
- infrastructure-as-code / declarative tools
- asset inventories / catalogs
- dynamic clients.

## Specification

### Resources

Resources **MUST** have the following fields:

| Field Name | Type   | Description                                                                                |
| ---------- | ------ | ------------------------------------------------------------------------------------------ |
| path       | string | The [path](#resource-paths) of the resource. This value **MUST** be readOnly.              |
| version    | string | The version of the resource.                                                               |
| properties | object | The user-settable properties of the resource. The server **MUST** not modify these values. |
| status     | object | The server-settable properties of the resource. These values **MUST** be readOnly.         |

#### Resource Paths

Resources **MUST** have a unique identifier to specify the resource in a
collection.

#### Resource Versions

A resource can evolve over time: new fields can be added, while others removed.

The `version` field on a resource exists to help the user and the server
understand which version of the resource is being operated on. The version is
helpful for [Apply](#apply) operations, where a value not being populated could
either imply that the client is unaware of the field's existence, or is
purposefully setting it to empty to clear it.

#### Resource Properties

Properties contain the user-specified values of the resource. Clients that are
only concerned with user-specified state may look at the properties field
exclusively to determine the user's intent.

- Properties **must** match the payload sent by the user. Modification of user
  settable-properties can cause endless reconciliation for declarative clients,
  as the resource configuration will never match the original payload sent.
- New fields introduced in properties **must** be optional, with the null value
  representing that the value is unspecified.

### Methods

### Resource Standard Methods

Resource standard methods represent common operations against a resource. These
perform the function of retrieval and mutation of resource state, including it's
existence.

#### Create

Create methods:

- **MUST** have an operation ID of "Create", followed by the resource singular.
- **MUST** use the HTTP `POST` method.
- **MUST** have an HTTP path of the resource parent name, with the resource
  singular appended.
- **MUST** have an HTTP body that is the resource, using a content type
  supported by the API.

#### Read

Read methods:

- **MUST** have an operation ID of "Read", followed by the resource singular.
- **MUST** use the HTTP `GET` method.
- **MUST** have an HTTP path identical to the resource path.

##### Read Request

The read request **must** support the following fields:

| Field Name | HTTP field type | type   | description                              |
| ---------- | --------------- | ------ | ---------------------------------------- |
| version    | query parameter | string | the version of the resource to retrieve. |

#### Apply

Apply methods:

- **MUST** have an operation ID of "Apply", followed by the resource singular.
- **MUST** use the HTTP `PUT` method.
- **MUST** have an HTTP path identical to the resource path.
- **MUST** have an HTTP body that is the resource, using a content type
  supported by the API.

#### Delete

Delete methods:

- **MUST** have an operation ID of "Delete", followed by the resource singular.
- **MUST** use the HTTP `DELETE` method.
- **MUST** have an HTTP path identical to the resource path.
- **MUST** have an empty HTTP body for the request.
- **MUST** have an empty HTTP body for the response.

### Collection Standard Methods

#### List

List methods:

- **MUST** have an operation ID of "List", followed by the resource singular.
- **MUST** use the HTTP `GET` method.
- **MUST** have an HTTP path identical to the collection path.
- **MUST** have an empty HTTP body for the request.

##### List Response

The list response **MUST** contain the following fields:

| Field Name | Type       | Description                           |
| ---------- | ---------- | ------------------------------------- |
| items      | [resource] | The items that match the list request |

### Custom Methods

Custom methods are used for use case that are not handled by the retrieval and
mutation of resource configuration. Examples include:

- The dynamic calculation of resource state, based on more than one resource.
- The modification of resource state that does not modify it's configuration,
  such as pausing a job.

#### Custom Retrieval Methods

Custom retrieval methods:

- **MUST** use the HTTP `GET` method.
- **MUST** have an HTTP path identical to the resource path, with a colon and
  the custom method name appended (":fetch")

#### Custom Mutation Methods

Custom mutation methods:

- **MUST** use the HTTP `POST` method.
- **MUST** have an HTTP path identical to the resource path, with a colon and
  the custom method name appended (":trigger")