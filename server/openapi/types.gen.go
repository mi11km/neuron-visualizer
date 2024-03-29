// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package openapi

// Defines values for HealthCheckResponseStatus.
const (
	ERROR HealthCheckResponseStatus = "ERROR"
	OK    HealthCheckResponseStatus = "OK"
)

// Defines values for NeuronCompartmentTypeName.
const (
	APICALDENDRITE NeuronCompartmentTypeName = "APICAL_DENDRITE"
	AXON           NeuronCompartmentTypeName = "AXON"
	BASALDENDRITE  NeuronCompartmentTypeName = "BASAL_DENDRITE"
	SOMA           NeuronCompartmentTypeName = "SOMA"
)

// Coordinate Coordinate for example x, y, z
type Coordinate = float64

// GetNeuronCompartmentsMembranePotentialResponse defines model for GetNeuronCompartmentsMembranePotentialResponse.
type GetNeuronCompartmentsMembranePotentialResponse struct {
	MembranePotentials []float64 `json:"membranePotentials"`

	// TimeStep Step Time of the membrane potential.
	TimeStep float64 `json:"timeStep"`
}

// GetNeuronCompartmentsResponse defines model for GetNeuronCompartmentsResponse.
type GetNeuronCompartmentsResponse struct {
	Compartments []NeuronCompartment `json:"compartments"`
}

// GetNeuronsResponse defines model for GetNeuronsResponse.
type GetNeuronsResponse struct {
	Neurons []Neuron `json:"neurons"`
}

// HealthCheckResponse defines model for HealthCheckResponse.
type HealthCheckResponse struct {
	// Message Message of the health check
	Message string `json:"message"`

	// Status Status of the health check
	Status HealthCheckResponseStatus `json:"status"`
}

// HealthCheckResponseStatus Status of the health check
type HealthCheckResponseStatus string

// Neuron defines model for Neuron.
type Neuron struct {
	// Name Name of the neuron
	Name NeuronName `json:"name"`
}

// NeuronCompartment defines model for NeuronCompartment.
type NeuronCompartment struct {
	// Id ID of the compartment, unique for the same neuron.
	Id int64 `json:"id"`

	// ParentId ID of the parent compartment. -1 indicates no parent compartment.
	ParentId int64 `json:"parentId"`

	// PositionX Coordinate for example x, y, z
	PositionX Coordinate `json:"positionX"`

	// PositionY Coordinate for example x, y, z
	PositionY Coordinate `json:"positionY"`

	// PositionZ Coordinate for example x, y, z
	PositionZ Coordinate `json:"positionZ"`

	// Radius Radius of the compartment.
	Radius float64               `json:"radius"`
	Type   NeuronCompartmentType `json:"type"`
}

// NeuronCompartmentType defines model for NeuronCompartmentType.
type NeuronCompartmentType struct {
	// Id ID of the compartment type.
	Id int64 `json:"id"`

	// Name Type of the compartment.
	//   SOMA (細胞体).
	//   AXON (軸索).
	//   BASAL_DENDRITE (基底樹状突起).
	//   APICAL_DENDRITE (尖端樹状突起).
	Name NeuronCompartmentTypeName `json:"name"`
}

// NeuronCompartmentTypeName Type of the compartment.
//
//	SOMA (細胞体).
//	AXON (軸索).
//	BASAL_DENDRITE (基底樹状突起).
//	APICAL_DENDRITE (尖端樹状突起).
type NeuronCompartmentTypeName string

// NeuronName Name of the neuron
type NeuronName = string
