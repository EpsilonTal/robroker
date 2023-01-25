package utils

type BrokerSettings struct {
	ID          string               `json:"id"`
	Username    string               `json:"username,omitempty"`
	Password    string               `json:"password,omitempty"`
	Catalog     interface{}          `json:"catalog"`
	Name        string               `json:"name"`
	Title       string               `json:"title,omitempty"`
	URL         *string              `json:"url,omitempty"`
	CreateTime  *string              `json:"createTime,omitempty"`
	Provision   MethodConfig         `json:"provision,omitempty"`
	Deprovision MethodConfig         `json:"deprovision,omitempty"`
	Bind        MethodConfig         `json:"bind,omitempty"`
	Unbind      MethodConfig         `json:"unbind,omitempty"`
	Instances   map[string]*Instance `json:"instances,omitempty"`
}

type MethodConfig struct {
	Status int64       `json:"status"`
	Body   interface{} `json:"body"`
}

type Catalog struct {
	Services interface{} `json:"services"`
}

type Instance struct {
	ID         string              `json:"instance_id"`
	ServiceID  string              `json:"service_id"`
	PlanID     string              `json:"plan_id"`
	CreateTime string              `json:"createTime"`
	Bindings   map[string]*Binding `json:"instances,omitempty"`
}

type Credentials struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}
type BindRequestBody struct {
	// ServiceID is the ID of the service to use for the service binding.
	ServiceID string `json:"service_id"`
	// PlanID is the ID of the plan to use for the service binding.
	PlanID string `json:"plan_id"`
	// Context is the contextual information under which the service binding is to be created.
	Context map[string]interface{} `json:"context,omitempty"`
	// AppGUID is the GUID of an application associated with the binding to be created. Optional.
	AppGUID string `json:"app_guid,omitempty"`
	// BindResource holds extra information about platform resources associated with the binding to
	// be created. CF-specific. Optional.
	BindResource map[string]interface{} `json:"bind_resource,omitempty"`
	// Parameters is a set of configuration options for the service binding. Optional.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}
type BindResponseBody struct {
	// Credentials is a free-form hash of credentials that can be used by applications or users to
	// access the service.
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	// SyslogDrainURl is a URL to which logs must be streamed. CF-specific. May only be supplied by a
	// service that declares a requirement for the 'syslog_drain' permission.
	SyslogDrainURL *string `json:"syslog_drain_url,omitempty"`
	// RouteServiceURL is a URL to which the platform must proxy requests to the application the
	// binding is for. CF-specific. May only be supplied by a service that declares a requirement for
	// the 'route_service' permission.
	RouteServiceURL *string `json:"route_service_url,omitempty"`
	// VolumeMounts is an array of configuration string for mounting volumes. CF-specific. May only be
	// supplied by a service that declares a requirement for the 'volume_mount' permission.
	VolumeMounts []interface{} `json:"volume_mounts,omitempty"`
	// Operation is an extra identifier supplied by the broker to identify asynchronous operations.
	Operation string `json:"operation,omitempty"`
}
type Binding struct {
	ID string `json:"binding_id"`
}
type Broker struct {
	Name       string  `json:"name"`
	Title      string  `json:"title,omitempty"`
	URL        *string `json:"url,omitempty"`
	CreateTime *string `json:"createTime"`
}
type DashboardClient struct {
	// ID is the ID of the OAuth client that the dashboard will use.
	ID     *string `json:"id"`
	Secret *string `json:"secret"`
	// RedirectURI is a URI for the service dashboard.
	RedirectURI *string `json:"redirect_uri"`
}
type DeprovisionResponseBody struct {
	// Operation is an extra identifier supplied by the broker to identify asynchronous operations.
	Operation string `json:"operation,omitempty"`
}
type OperationResponseBody struct {
	// State is the state of the queried operation.
	State string `json:"state"`
	// Description is a message from the broker describing the current state of the operation.
	Description string `json:"description,omitempty"`
}
type Plan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Free is true if the plan is free and false is paid.
	// The OSB API specifies the default to be true so nil corresponds with true.
	Free *bool `json:"free"`
	// Bindable specifies whether the plan can be bound to. If not specified, then consumers
	// should use the Bindalbe property from the Service.
	Bindable *bool                  `json:"bindable"`
	Schemas  *Schemas               `json:"schemas"`
	Metadata map[string]interface{} `json:"metadata"`
}
type ProvisionRequestBody struct {
	// ServiceID is the ID of the service to use for the service instance.
	ServiceID string `json:"service_id"`
	// PlanID is the ID of the plan to use for the service instance.
	PlanID string `json:"plan_id"`
	// Context is platform-specific contextual information under which the service instance is to be
	// provisioned. Context was added in version 2.12 of the OSB API and is only sent for versions
	// 2.12 or later. Optional.
	Context map[string]interface{} `json:"context,omitempty"`
	// OrganizationGUID is the platform GUID for the organization under which the service plan is to
	// be provisioned. CF-specific. Optional.
	OrganizationGUID string `json:"organization_guid,omitempty"`
	// SpaceGUID is the identifier for the project space within the platform organization.
	// CF-specific. Optional.
	SpaceGUID string `json:"space_guid,omitempty"`
	// Parameters is a set of configuration options for the service instance. Optional.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}
type ProvisionResponseBody struct {
	// DashboardURL is the URL of a web-based management user interface for the service instance.
	DashboardURL string `json:"dashboard_url,omitempty"`
	// Operation is an extra identifier supplied by the broker to identify asynchronous operations.
	Operation string `json:"operation,omitempty"`
}
type Schemas struct {
	ServiceInstance *ServiceInstanceSchema `json:"service_instance"`
	ServiceBinding  *ServiceBindingSchema  `json:"service_binding"`
}
type Service struct {
	Name        string   `json:"name"`
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	// Requires is a list of permisisons a user would have to give to the service to provision
	// it.
	Requires []string `json:"requires"`
	// Bindable specifies the default value for the Bindable property of its Plans.
	Bindable        bool                   `json:"bindable"`
	Metadata        map[string]interface{} `json:"metadata"`
	DashboardClient *DashboardClient       `json:"dashboard_client"`
	// PlanUpdateable is true iff the service supports up/downgrade of some plans.
	PlanUpdateable bool `json:"plan_updateable"`
	// Plans is a list of plans for this service.
	Plans []Plan `json:"plans"`
}
type ServiceBindingSchema struct {
	// Create is the json schema describing binding creation.
	Create *map[string]interface{} `json:"create"`
}
type ServiceInstanceSchema struct {
	// Create is the json schema describing instance creation.
	Create *map[string]interface{} `json:"create"`
	// Update is the json schema describing instance update.
	Update *map[string]interface{} `json:"update"`
}
type UnbindResponseBody struct {
	// Operation is an extra identifier supplied by the broker to identify asynchronous operations.
	Operation string `json:"operation,omitempty"`
}
type UpdateInstancePreviousValues struct {
	// ServiceID is the ID of the service used by the service instance. This field is deprecated
	// because it should be immutable.
	ServiceID string `json:"service_id,omitempty"`
	// PlanID is the ID of the plan used by the service instance prior to the update. Optional.
	PlanID string `json:"plan_id,omitempty"`
	// OrganizationID is the ID of the organization specified for the service instance. CF-specific.
	// Optional.
	OrganizationID string `json:"organization_id,omitempty"`
	// SpaceID is the ID of the space specified for the service instance. CF-specific. Optional.
	SpaceID string `json:"space_id,omitempty"`
}
type UpdateInstanceRequestBody struct {
	// ServiceID is the ID of the service used by the service instance.
	ServiceID string `json:"service_id"`
	// PlanID is the ID of the plan to use for the service instance.
	PlanID string `json:"plan_id,omitempty"`
	// Context is platform-specific contextual information under which the service instance is to be
	// provisioned. Context was added in version 2.12 of the OSB API and is only sent for versions
	// 2.12 or later. Optional.
	Context map[string]interface{} `json:"context,omitempty"`
	// Parameters is a set of configuration options for the service instance. Optional.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// PreviousValues stores the information about the service instance prior to the update.
	PreviousValues *UpdateInstancePreviousValues `json:"previous_values,omitempty"`
}
type UpdateInstanceResponseBody struct {
	// Operation is an extra identifier supplied by the broker to identify asynchronous operations.
	Operation string `json:"operation,omitempty"`
}
