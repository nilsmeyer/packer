/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// Project struct for Project
type Project struct {
	Id           string                `json:"id,omitempty"`
	Name         string                `json:"name,omitempty"`
	Services     []ProjectServices     `json:"services,omitempty"`
	Flavour      string                `json:"flavour,omitempty"`
	ModifiedOn   time.Time             `json:"modifiedOn,omitempty"`
	ModifiedBy   string                `json:"modifiedBy,omitempty"`
	CreatedBy    string                `json:"createdBy,omitempty"`
	CreatedOn    time.Time             `json:"createdOn,omitempty"`
	AccessRights []ProjectAccessRights `json:"accessRights,omitempty"`
	Processing   bool                  `json:"processing,omitempty"`
	Created      bool                  `json:"created,omitempty"`
	Queue        []Event               `json:"queue,omitempty"`
	State        string                `json:"state,omitempty"`
	Tag          map[string]string     `json:"tag,omitempty"`
	Project      string                `json:"project,omitempty"`
	BankAccount  string                `json:"bankAccount,omitempty"`
	Organisation string                `json:"organisation,omitempty"`
	Billing      ProjectBilling        `json:"billing,omitempty"`
	Invoices     []ProjectInvoices     `json:"invoices,omitempty"`
	Payments     []Payment             `json:"payments,omitempty"`
	Verified     string                `json:"verified,omitempty"`
	Active       bool                  `json:"active,omitempty"`
	Limit        ProjectLimit          `json:"limit,omitempty"`
	Threshold    ProjectThreshold      `json:"threshold,omitempty"`
	Roles        []ProjectRoles        `json:"roles,omitempty"`
	NetworkAcl   []UserNetworkAcl      `json:"networkAcl,omitempty"`
	Compliance   ProjectCompliance     `json:"compliance,omitempty"`
	Transfer     ProjectTransfer       `json:"transfer,omitempty"`
}
