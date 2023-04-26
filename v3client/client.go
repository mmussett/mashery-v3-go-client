package v3client

import (
	"context"
	"errors"
	"fmt"
	"github.com/aliakseiyanchuk/mashery-v3-go-client/masherytypes"
	"github.com/aliakseiyanchuk/mashery-v3-go-client/transport"
	"net/url"
)

type WildcardClient interface {
	// FetchAny fetch an arbitrary resource from Mashery V3
	FetchAny(ctx context.Context, resource string, qs *url.Values) (*transport.WrappedResponse, error)
	// DeleteAny Delete an arbitrary resource from Mashery V3
	DeleteAny(ctx context.Context, resource string) (*transport.WrappedResponse, error)
	// PostAny post any value to an arbitrary resource
	PostAny(ctx context.Context, resource string, body interface{}) (*transport.WrappedResponse, error)
	// PutAny put any value to an arbitrary resource
	PutAny(ctx context.Context, resource string, body interface{}) (*transport.WrappedResponse, error)

	Close(ctx context.Context)
}

type Client interface {
	// GetPublicDomains Get public domains in this area
	GetPublicDomains(ctx context.Context) ([]string, error)
	GetSystemDomains(ctx context.Context) ([]string, error)

	// GetApplication Retrieve the details of the application
	GetApplication(ctx context.Context, appId masherytypes.ApplicationIdentifier) (*masherytypes.Application, error)
	GetApplicationPackageKeys(ctx context.Context, appId masherytypes.ApplicationIdentifier) ([]masherytypes.PackageKey, error)
	CountApplicationPackageKeys(ctx context.Context, appId masherytypes.ApplicationIdentifier) (int64, error)
	GetFullApplication(ctx context.Context, id masherytypes.ApplicationIdentifier) (*masherytypes.Application, error)
	CreateApplication(ctx context.Context, app masherytypes.Application, memberId masherytypes.MemberIdentifier) (*masherytypes.Application, error)
	UpdateApplication(ctx context.Context, app masherytypes.Application) (*masherytypes.Application, error)
	DeleteApplication(ctx context.Context, appId masherytypes.ApplicationIdentifier) error
	CountApplicationsOfMember(ctx context.Context, memberId masherytypes.MemberIdentifier) (int64, error)
	ListApplications(ctx context.Context) ([]masherytypes.Application, error)

	// Email template sets
	GetEmailTemplateSet(ctx context.Context, id string) (*masherytypes.EmailTemplateSet, error)
	ListEmailTemplateSets(ctx context.Context) ([]masherytypes.EmailTemplateSet, error)
	ListEmailTemplateSetsFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.EmailTemplateSet, error)

	// Endpoints
	ListEndpoints(ctx context.Context, serviceId masherytypes.ServiceIdentifier) ([]masherytypes.AddressableV3Object, error)
	ListEndpointsWithFullInfo(ctx context.Context, serviceId masherytypes.ServiceIdentifier) ([]masherytypes.Endpoint, error)
	CreateEndpoint(ctx context.Context, serviceId masherytypes.ServiceIdentifier, endp masherytypes.Endpoint) (*masherytypes.Endpoint, error)
	UpdateEndpoint(ctx context.Context, endp masherytypes.Endpoint) (*masherytypes.Endpoint, error)
	GetEndpoint(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) (*masherytypes.Endpoint, error)
	DeleteEndpoint(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) error
	CountEndpointsOf(ctx context.Context, serviceId masherytypes.ServiceIdentifier) (int64, error)

	// Endpoint methods
	ListEndpointMethods(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) ([]masherytypes.ServiceEndpointMethod, error)
	ListEndpointMethodsWithFullInfo(ctx context.Context, dent masherytypes.ServiceEndpointIdentifier) ([]masherytypes.ServiceEndpointMethod, error)
	CreateEndpointMethod(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, methodUpsert masherytypes.ServiceEndpointMethod) (*masherytypes.ServiceEndpointMethod, error)
	UpdateEndpointMethod(ctx context.Context, methUpsert masherytypes.ServiceEndpointMethod) (*masherytypes.ServiceEndpointMethod, error)
	GetEndpointMethod(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) (*masherytypes.ServiceEndpointMethod, error)
	DeleteEndpointMethod(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) error
	CountEndpointsMethodsOf(ctx context.Context, identifier masherytypes.ServiceEndpointIdentifier) (int64, error)

	// Endpoint method filters
	ListEndpointMethodFilters(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) ([]masherytypes.ServiceEndpointMethodFilter, error)
	ListEndpointMethodFiltersWithFullInfo(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) ([]masherytypes.ServiceEndpointMethodFilter, error)
	CreateEndpointMethodFilter(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, filterUpsert masherytypes.ServiceEndpointMethodFilter) (*masherytypes.ServiceEndpointMethodFilter, error)
	UpdateEndpointMethodFilter(ctx context.Context, methUpsert masherytypes.ServiceEndpointMethodFilter) (*masherytypes.ServiceEndpointMethodFilter, error)
	GetEndpointMethodFilter(ctx context.Context, ident masherytypes.ServiceEndpointMethodFilterIdentifier) (*masherytypes.ServiceEndpointMethodFilter, error)
	DeleteEndpointMethodFilter(ctx context.Context, ident masherytypes.ServiceEndpointMethodFilterIdentifier) error
	CountEndpointsMethodsFiltersOf(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) (int64, error)

	// Member
	GetMember(ctx context.Context, id masherytypes.MemberIdentifier) (*masherytypes.Member, error)
	GetFullMember(ctx context.Context, id masherytypes.MemberIdentifier) (*masherytypes.Member, error)
	CreateMember(ctx context.Context, member masherytypes.Member) (*masherytypes.Member, error)
	UpdateMember(ctx context.Context, member masherytypes.Member) (*masherytypes.Member, error)
	DeleteMember(ctx context.Context, memberId masherytypes.MemberIdentifier) error
	ListMembers(ctx context.Context) ([]masherytypes.Member, error)

	// Packages
	GetPackage(ctx context.Context, id masherytypes.PackageIdentifier) (*masherytypes.Package, error)
	CreatePackage(ctx context.Context, pack masherytypes.Package) (*masherytypes.Package, error)
	UpdatePackage(ctx context.Context, pack masherytypes.Package) (*masherytypes.Package, error)
	DeletePackage(ctx context.Context, packId masherytypes.PackageIdentifier) error
	ListPackages(ctx context.Context) ([]masherytypes.Package, error)

	// Package plans
	CreatePlanService(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) (*masherytypes.AddressableV3Object, error)
	CheckPlanServiceExists(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) (bool, error)
	DeletePlanService(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) error
	CreatePlanEndpoint(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier) (*masherytypes.AddressableV3Object, error)
	CheckPlanEndpointExists(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier) (bool, error)
	DeletePlanEndpoint(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier) error
	ListPlanEndpoints(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) ([]masherytypes.AddressableV3Object, error)
	CountPlanService(ctx context.Context, ident masherytypes.PackagePlanIdentifier) (int64, error)
	GetPlan(ctx context.Context, ident masherytypes.PackagePlanIdentifier) (*masherytypes.Plan, error)
	CreatePlan(ctx context.Context, packageId masherytypes.PackageIdentifier, plan masherytypes.Plan) (*masherytypes.Plan, error)
	UpdatePlan(ctx context.Context, plan masherytypes.Plan) (*masherytypes.Plan, error)
	DeletePlan(ctx context.Context, ident masherytypes.PackagePlanIdentifier) error
	CountPlans(ctx context.Context, packageId masherytypes.PackageIdentifier) (int64, error)
	ListPlans(ctx context.Context, packageId masherytypes.PackageIdentifier) ([]masherytypes.Plan, error)
	ListPlanServices(ctx context.Context, ident masherytypes.PackagePlanIdentifier) ([]masherytypes.Service, error)

	CountPlanEndpoints(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) (int64, error)

	// Plan methods
	ListPackagePlanMethods(ctx context.Context, id masherytypes.PackagePlanServiceEndpointIdentifier) ([]masherytypes.PackagePlanServiceEndpointMethod, error)
	GetPackagePlanMethod(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) (*masherytypes.PackagePlanServiceEndpointMethod, error)
	CreatePackagePlanMethod(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) (*masherytypes.PackagePlanServiceEndpointMethod, error)
	DeletePackagePlanMethod(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) error

	// PLan method filter
	GetPackagePlanMethodFilter(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) (*masherytypes.PackagePlanServiceEndpointMethodFilter, error)
	CreatePackagePlanMethodFilter(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodFilterIdentifier) (*masherytypes.PackagePlanServiceEndpointMethodFilter, error)
	DeletePackagePlanMethodFilter(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) error

	// Package key
	GetPackageKey(ctx context.Context, id masherytypes.PackageKeyIdentifier) (*masherytypes.PackageKey, error)
	CreatePackageKey(ctx context.Context, appId masherytypes.ApplicationIdentifier, packageKey masherytypes.PackageKey) (*masherytypes.PackageKey, error)
	UpdatePackageKey(ctx context.Context, packageKey masherytypes.PackageKey) (*masherytypes.PackageKey, error)
	DeletePackageKey(ctx context.Context, keyId masherytypes.PackageKeyIdentifier) error
	ListPackageKeysFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.PackageKey, error)
	ListPackageKeys(ctx context.Context) ([]masherytypes.PackageKey, error)

	// Roles
	GetRole(ctx context.Context, id string) (*masherytypes.Role, error)
	ListRoles(ctx context.Context) ([]masherytypes.Role, error)
	ListRolesFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.Role, error)

	// GetService retrieves service based on the service identifier
	GetService(ctx context.Context, id masherytypes.ServiceIdentifier) (*masherytypes.Service, error)
	CreateService(ctx context.Context, service masherytypes.Service) (*masherytypes.Service, error)
	UpdateService(ctx context.Context, service masherytypes.Service) (*masherytypes.Service, error)
	DeleteService(ctx context.Context, serviceId masherytypes.ServiceIdentifier) error
	ListServicesFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.Service, error)
	ListServices(ctx context.Context) ([]masherytypes.Service, error)
	CountServices(ctx context.Context, params map[string]string) (int64, error)

	ListErrorSets(ctx context.Context, serviceId masherytypes.ServiceIdentifier, qs url.Values) ([]masherytypes.ErrorSet, error)
	GetErrorSet(ctx context.Context, ident masherytypes.ErrorSetIdentifier) (*masherytypes.ErrorSet, error)
	CreateErrorSet(ctx context.Context, serviceId masherytypes.ServiceIdentifier, set masherytypes.ErrorSet) (*masherytypes.ErrorSet, error)
	UpdateErrorSet(ctx context.Context, setData masherytypes.ErrorSet) (*masherytypes.ErrorSet, error)
	DeleteErrorSet(ctx context.Context, ident masherytypes.ErrorSetIdentifier) error
	UpdateErrorSetMessage(ctx context.Context, ident masherytypes.ErrorSetIdentifier, msg masherytypes.MasheryErrorMessage) (*masherytypes.MasheryErrorMessage, error)

	GetServiceRoles(ctx context.Context, serviceId masherytypes.ServiceIdentifier) ([]masherytypes.RolePermission, error)
	SetServiceRoles(ctx context.Context, id masherytypes.ServiceIdentifier, roles []masherytypes.RolePermission) error
	DeleteServiceRoles(ctx context.Context, id masherytypes.ServiceIdentifier) error

	// Service cache
	GetServiceCache(ctx context.Context, id string) (*masherytypes.ServiceCache, error)
	CreateServiceCache(ctx context.Context, id string, service masherytypes.ServiceCache) (*masherytypes.ServiceCache, error)
	UpdateServiceCache(ctx context.Context, id string, service masherytypes.ServiceCache) (*masherytypes.ServiceCache, error)
	DeleteServiceCache(ctx context.Context, id string) error

	// GetServiceOAuthSecurityProfile Service OAuth
	GetServiceOAuthSecurityProfile(ctx context.Context, id masherytypes.ServiceIdentifier) (*masherytypes.MasheryOAuth, error)
	CreateServiceOAuthSecurityProfile(ctx context.Context, service masherytypes.MasheryOAuth) (*masherytypes.MasheryOAuth, error)
	UpdateServiceOAuthSecurityProfile(ctx context.Context, service masherytypes.MasheryOAuth) (*masherytypes.MasheryOAuth, error)
	DeleteServiceOAuthSecurityProfile(ctx context.Context, id masherytypes.ServiceIdentifier) error
}

type PluggableClient struct {
	schema    *ClientMethodSchema
	transport *transport.V3Transport
}

// FixedSchemeClient Fixed method scheme client that will not allow changing schema after it was created.
type FixedSchemeClient struct {
	PluggableClient
}

func (fsc *FixedSchemeClient) AssumeSchema(_ *ClientMethodSchema) {
	panic("This client cannot change method schema after it was created")
}

// AssumeSchema Allows a pluggable client to assume a schema. This method should be mainly used in the test contexts.
func (c *PluggableClient) AssumeSchema(sch *ClientMethodSchema) {
	c.schema = sch
}

func (c *PluggableClient) notImplemented(meth string) error {
	return errors.New(fmt.Sprintf("No implementation method was supplied for method %s", meth))
}

// -----------------------------------------------------------------------------------------------------------------
// Application methods
// -----------------------------------------------------------------------------------------------------------------

type ClientMethodSchema struct {
	// Public and System Domains
	GetPublicDomains func(ctx context.Context, transport *transport.V3Transport) ([]string, error)
	GetSystemDomains func(ctx context.Context, transport *transport.V3Transport) ([]string, error)

	// Applications
	GetApplicationContext       func(ctx context.Context, appId masherytypes.ApplicationIdentifier, transport *transport.V3Transport) (*masherytypes.Application, error)
	GetApplicationPackageKeys   func(ctx context.Context, appId masherytypes.ApplicationIdentifier, transport *transport.V3Transport) ([]masherytypes.PackageKey, error)
	CountApplicationPackageKeys func(ctx context.Context, appId masherytypes.ApplicationIdentifier, c *transport.V3Transport) (int64, error)
	GetFullApplication          func(ctx context.Context, id masherytypes.ApplicationIdentifier, c *transport.V3Transport) (*masherytypes.Application, error)
	CreateApplication           func(ctx context.Context, app masherytypes.Application, memberId masherytypes.MemberIdentifier, c *transport.V3Transport) (*masherytypes.Application, error)
	UpdateApplication           func(ctx context.Context, app masherytypes.Application, c *transport.V3Transport) (*masherytypes.Application, error)
	DeleteApplication           func(ctx context.Context, appId masherytypes.ApplicationIdentifier, c *transport.V3Transport) error
	CountApplicationsOfMember   func(ctx context.Context, memberId masherytypes.MemberIdentifier, c *transport.V3Transport) (int64, error)
	ListApplications            func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.Application, error)

	// Email template set
	GetEmailTemplateSet           func(ctx context.Context, id string, c *transport.V3Transport) (*masherytypes.EmailTemplateSet, error)
	ListEmailTemplateSets         func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.EmailTemplateSet, error)
	ListEmailTemplateSetsFiltered func(ctx context.Context, params map[string]string, fields []string, c *transport.V3Transport) ([]masherytypes.EmailTemplateSet, error)

	// Endpoints
	ListEndpoints             func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, c *transport.V3Transport) ([]masherytypes.AddressableV3Object, error)
	ListEndpointsWithFullInfo func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, c *transport.V3Transport) ([]masherytypes.Endpoint, error)
	CreateEndpoint            func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, endp masherytypes.Endpoint, c *transport.V3Transport) (*masherytypes.Endpoint, error)
	UpdateEndpoint            func(ctx context.Context, endpoint masherytypes.Endpoint, c *transport.V3Transport) (*masherytypes.Endpoint, error)
	GetEndpoint               func(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, c *transport.V3Transport) (*masherytypes.Endpoint, error)
	DeleteEndpoint            func(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, c *transport.V3Transport) error
	CountEndpointsOf          func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, c *transport.V3Transport) (int64, error)

	// Endpoint Methods
	ListEndpointMethods             func(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, c *transport.V3Transport) ([]masherytypes.ServiceEndpointMethod, error)
	ListEndpointMethodsWithFullInfo func(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, c *transport.V3Transport) ([]masherytypes.ServiceEndpointMethod, error)
	CreateEndpointMethod            func(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, methoUpsert masherytypes.ServiceEndpointMethod, c *transport.V3Transport) (*masherytypes.ServiceEndpointMethod, error)
	UpdateEndpointMethod            func(ctx context.Context, methUpsert masherytypes.ServiceEndpointMethod, c *transport.V3Transport) (*masherytypes.ServiceEndpointMethod, error)
	GetEndpointMethod               func(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, c *transport.V3Transport) (*masherytypes.ServiceEndpointMethod, error)
	DeleteEndpointMethod            func(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, c *transport.V3Transport) error
	CountEndpointsMethodsOf         func(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, c *transport.V3Transport) (int64, error)

	// Endpoint method filters
	ListEndpointMethodFilters             func(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, c *transport.V3Transport) ([]masherytypes.ServiceEndpointMethodFilter, error)
	ListEndpointMethodFiltersWithFullInfo func(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, c *transport.V3Transport) ([]masherytypes.ServiceEndpointMethodFilter, error)
	CreateEndpointMethodFilter            func(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, filterUpsert masherytypes.ServiceEndpointMethodFilter, c *transport.V3Transport) (*masherytypes.ServiceEndpointMethodFilter, error)
	UpdateEndpointMethodFilter            func(ctx context.Context, methUpsert masherytypes.ServiceEndpointMethodFilter, c *transport.V3Transport) (*masherytypes.ServiceEndpointMethodFilter, error)
	GetEndpointMethodFilter               func(ctx context.Context, ident masherytypes.ServiceEndpointMethodFilterIdentifier, c *transport.V3Transport) (*masherytypes.ServiceEndpointMethodFilter, error)
	DeleteEndpointMethodFilter            func(ctx context.Context, dent masherytypes.ServiceEndpointMethodFilterIdentifier, c *transport.V3Transport) error
	CountEndpointsMethodsFiltersOf        func(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, c *transport.V3Transport) (int64, error)

	// Member
	GetMember     func(ctx context.Context, id masherytypes.MemberIdentifier, c *transport.V3Transport) (*masherytypes.Member, error)
	GetFullMember func(ctx context.Context, id masherytypes.MemberIdentifier, c *transport.V3Transport) (*masherytypes.Member, error)
	CreateMember  func(ctx context.Context, member masherytypes.Member, c *transport.V3Transport) (*masherytypes.Member, error)
	UpdateMember  func(ctx context.Context, member masherytypes.Member, c *transport.V3Transport) (*masherytypes.Member, error)
	DeleteMember  func(ctx context.Context, memberId masherytypes.MemberIdentifier, c *transport.V3Transport) error
	ListMembers   func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.Member, error)

	// Packages
	GetPackage    func(ctx context.Context, id masherytypes.PackageIdentifier, c *transport.V3Transport) (*masherytypes.Package, error)
	CreatePackage func(ctx context.Context, pack masherytypes.Package, c *transport.V3Transport) (*masherytypes.Package, error)
	UpdatePackage func(ctx context.Context, pack masherytypes.Package, c *transport.V3Transport) (*masherytypes.Package, error)
	DeletePackage func(ctx context.Context, packId masherytypes.PackageIdentifier, c *transport.V3Transport) error
	ListPackages  func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.Package, error)

	// Package plans
	CreatePlanService       func(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier, c *transport.V3Transport) (*masherytypes.AddressableV3Object, error)
	CheckPlanServiceExists  func(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier, c *transport.V3Transport) (bool, error)
	DeletePlanService       func(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier, c *transport.V3Transport) error
	CreatePlanEndpoint      func(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier, c *transport.V3Transport) (*masherytypes.AddressableV3Object, error)
	CheckPlanEndpointExists func(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier, c *transport.V3Transport) (bool, error)
	DeletePlanEndpoint      func(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier, c *transport.V3Transport) error
	ListPlanEndpoints       func(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier, c *transport.V3Transport) ([]masherytypes.AddressableV3Object, error)

	CountPlanEndpoints func(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier, c *transport.V3Transport) (int64, error)
	CountPlanService   func(ctx context.Context, ident masherytypes.PackagePlanIdentifier, c *transport.V3Transport) (int64, error)
	GetPlan            func(ctx context.Context, ident masherytypes.PackagePlanIdentifier, c *transport.V3Transport) (*masherytypes.Plan, error)
	CreatePlan         func(ctx context.Context, packageId masherytypes.PackageIdentifier, plan masherytypes.Plan, c *transport.V3Transport) (*masherytypes.Plan, error)
	UpdatePlan         func(ctx context.Context, plan masherytypes.Plan, c *transport.V3Transport) (*masherytypes.Plan, error)
	DeletePlan         func(ctx context.Context, dent masherytypes.PackagePlanIdentifier, c *transport.V3Transport) error
	CountPlans         func(ctx context.Context, packageId masherytypes.PackageIdentifier, c *transport.V3Transport) (int64, error)
	ListPlans          func(ctx context.Context, packageId masherytypes.PackageIdentifier, c *transport.V3Transport) ([]masherytypes.Plan, error)
	ListPlanServices   func(ctx context.Context, dent masherytypes.PackagePlanIdentifier, c *transport.V3Transport) ([]masherytypes.Service, error)

	// Plan methods
	ListPackagePlanMethods  func(ctx context.Context, id masherytypes.PackagePlanServiceEndpointIdentifier, c *transport.V3Transport) ([]masherytypes.PackagePlanServiceEndpointMethod, error)
	GetPackagePlanMethod    func(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier, c *transport.V3Transport) (*masherytypes.PackagePlanServiceEndpointMethod, error)
	CreatePackagePlanMethod func(ctx context.Context, ident masherytypes.PackagePlanServiceEndpointMethodIdentifier, c *transport.V3Transport) (*masherytypes.PackagePlanServiceEndpointMethod, error)
	DeletePackagePlanMethod func(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier, c *transport.V3Transport) error

	// Plan method filter
	GetPackagePlanMethodFilter    func(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier, c *transport.V3Transport) (*masherytypes.PackagePlanServiceEndpointMethodFilter, error)
	CreatePackagePlanMethodFilter func(ctx context.Context, ident masherytypes.PackagePlanServiceEndpointMethodFilterIdentifier, c *transport.V3Transport) (*masherytypes.PackagePlanServiceEndpointMethodFilter, error)
	DeletePackagePlanMethodFilter func(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier, c *transport.V3Transport) error

	// Package key
	GetPackageKey           func(ctx context.Context, id masherytypes.PackageKeyIdentifier, c *transport.V3Transport) (*masherytypes.PackageKey, error)
	CreatePackageKey        func(ctx context.Context, appId masherytypes.ApplicationIdentifier, packageKey masherytypes.PackageKey, c *transport.V3Transport) (*masherytypes.PackageKey, error)
	UpdatePackageKey        func(ctx context.Context, packageKey masherytypes.PackageKey, c *transport.V3Transport) (*masherytypes.PackageKey, error)
	DeletePackageKey        func(ctx context.Context, keyId masherytypes.PackageKeyIdentifier, c *transport.V3Transport) error
	ListPackageKeysFiltered func(ctx context.Context, params map[string]string, fields []string, c *transport.V3Transport) ([]masherytypes.PackageKey, error)
	ListPackageKeys         func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.PackageKey, error)

	// Roles
	GetRole           func(ctx context.Context, id string, c *transport.V3Transport) (*masherytypes.Role, error)
	ListRoles         func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.Role, error)
	ListRolesFiltered func(ctx context.Context, params map[string]string, fields []string, c *transport.V3Transport) ([]masherytypes.Role, error)

	// Services
	GetService           func(ctx context.Context, id masherytypes.ServiceIdentifier, c *transport.V3Transport) (*masherytypes.Service, error)
	CreateService        func(ctx context.Context, service masherytypes.Service, c *transport.V3Transport) (*masherytypes.Service, error)
	UpdateService        func(ctx context.Context, service masherytypes.Service, c *transport.V3Transport) (*masherytypes.Service, error)
	DeleteService        func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, c *transport.V3Transport) error
	ListServicesFiltered func(ctx context.Context, params map[string]string, fields []string, c *transport.V3Transport) ([]masherytypes.Service, error)
	ListServices         func(ctx context.Context, c *transport.V3Transport) ([]masherytypes.Service, error)
	CountServices        func(ctx context.Context, params map[string]string, c *transport.V3Transport) (int64, error)

	ListErrorSets         func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, qs url.Values, c *transport.V3Transport) ([]masherytypes.ErrorSet, error)
	GetErrorSet           func(ctx context.Context, ident masherytypes.ErrorSetIdentifier, c *transport.V3Transport) (*masherytypes.ErrorSet, error)
	CreateErrorSet        func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, set masherytypes.ErrorSet, c *transport.V3Transport) (*masherytypes.ErrorSet, error)
	UpdateErrorSet        func(ctx context.Context, setData masherytypes.ErrorSet, c *transport.V3Transport) (*masherytypes.ErrorSet, error)
	DeleteErrorSet        func(ctx context.Context, ident masherytypes.ErrorSetIdentifier, c *transport.V3Transport) error
	UpdateErrorSetMessage func(ctx context.Context, sident masherytypes.ErrorSetIdentifier, msg masherytypes.MasheryErrorMessage, c *transport.V3Transport) (*masherytypes.MasheryErrorMessage, error)

	GetServiceRoles    func(ctx context.Context, serviceId masherytypes.ServiceIdentifier, c *transport.V3Transport) ([]masherytypes.RolePermission, error)
	SetServiceRoles    func(ctx context.Context, id masherytypes.ServiceIdentifier, roles []masherytypes.RolePermission, c *transport.V3Transport) error
	DeleteServiceRoles func(ctx context.Context, id masherytypes.ServiceIdentifier, c *transport.V3Transport) error

	// Service cache
	GetServiceCache    func(ctx context.Context, id string, c *transport.V3Transport) (*masherytypes.ServiceCache, error)
	CreateServiceCache func(ctx context.Context, id string, service masherytypes.ServiceCache, c *transport.V3Transport) (*masherytypes.ServiceCache, error)
	UpdateServiceCache func(ctx context.Context, id string, service masherytypes.ServiceCache, c *transport.V3Transport) (*masherytypes.ServiceCache, error)
	DeleteServiceCache func(ctx context.Context, id string, c *transport.V3Transport) error

	// Service OAuth
	GetServiceOAuthSecurityProfile    func(ctx context.Context, id masherytypes.ServiceIdentifier, c *transport.V3Transport) (*masherytypes.MasheryOAuth, error)
	CreateServiceOAuthSecurityProfile func(ctx context.Context, service masherytypes.MasheryOAuth, c *transport.V3Transport) (*masherytypes.MasheryOAuth, error)
	UpdateServiceOAuthSecurityProfile func(ctx context.Context, service masherytypes.MasheryOAuth, c *transport.V3Transport) (*masherytypes.MasheryOAuth, error)
	DeleteServiceOAuthSecurityProfile func(ctx context.Context, id masherytypes.ServiceIdentifier, c *transport.V3Transport) error
}

func (c *PluggableClient) ListErrorSets(ctx context.Context, serviceId masherytypes.ServiceIdentifier, qs url.Values) ([]masherytypes.ErrorSet, error) {
	if c.schema.ListErrorSets != nil {
		return c.schema.ListErrorSets(ctx, serviceId, qs, c.transport)
	} else {
		return []masherytypes.ErrorSet{}, c.notImplemented("ListErrorSets")
	}
}

func (c *PluggableClient) GetErrorSet(ctx context.Context, ident masherytypes.ErrorSetIdentifier) (*masherytypes.ErrorSet, error) {
	if c.schema.GetErrorSet != nil {
		return c.schema.GetErrorSet(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("GetErrorSet")
	}
}

func (c *PluggableClient) CreateErrorSet(ctx context.Context, serviceId masherytypes.ServiceIdentifier, set masherytypes.ErrorSet) (*masherytypes.ErrorSet, error) {
	if c.schema.CreateErrorSet != nil {
		return c.schema.CreateErrorSet(ctx, serviceId, set, c.transport)
	} else {
		return nil, c.notImplemented("CreateErrorSet")
	}
}

func (c *PluggableClient) UpdateErrorSet(ctx context.Context, setData masherytypes.ErrorSet) (*masherytypes.ErrorSet, error) {
	if c.schema.UpdateErrorSet != nil {
		return c.schema.UpdateErrorSet(ctx, setData, c.transport)
	} else {
		return nil, c.notImplemented("UpdateErrorSet")
	}
}

func (c *PluggableClient) DeleteErrorSet(ctx context.Context, ident masherytypes.ErrorSetIdentifier) error {
	if c.schema.DeleteErrorSet != nil {
		return c.schema.DeleteErrorSet(ctx, ident, c.transport)
	} else {
		return c.notImplemented("DeleteErrorSet")
	}
}

func (c *PluggableClient) UpdateErrorSetMessage(ctx context.Context, ident masherytypes.ErrorSetIdentifier, msg masherytypes.MasheryErrorMessage) (*masherytypes.MasheryErrorMessage, error) {
	if c.schema.UpdateErrorSetMessage != nil {
		return c.schema.UpdateErrorSetMessage(ctx, ident, msg, c.transport)
	} else {
		return nil, c.notImplemented("UpdateErrorSetMessage")
	}
}

func (c *PluggableClient) GetServiceRoles(ctx context.Context, serviceId masherytypes.ServiceIdentifier) ([]masherytypes.RolePermission, error) {
	if c.schema.GetServiceRoles != nil {
		return c.schema.GetServiceRoles(ctx, serviceId, c.transport)
	} else {
		return []masherytypes.RolePermission{}, c.notImplemented("GetServiceRoles")
	}
}

func (c *PluggableClient) SetServiceRoles(ctx context.Context, serviceId masherytypes.ServiceIdentifier, perms []masherytypes.RolePermission) error {
	if c.schema.SetServiceRoles != nil {
		return c.schema.SetServiceRoles(ctx, serviceId, perms, c.transport)
	} else {
		return c.notImplemented("SetServiceRoles")
	}
}
func (c *PluggableClient) DeleteServiceRoles(ctx context.Context, serviceId masherytypes.ServiceIdentifier) error {
	if c.schema.DeleteServiceRoles != nil {
		return c.schema.DeleteServiceRoles(ctx, serviceId, c.transport)
	} else {
		return c.notImplemented("DeleteServiceRoles")
	}
}

func (c *PluggableClient) GetPublicDomains(ctx context.Context) ([]string, error) {
	if c.schema.GetPublicDomains != nil {
		return c.schema.GetPublicDomains(ctx, c.transport)
	} else {
		return []string{}, c.notImplemented("GetPublicDomains")
	}
}

func (c *PluggableClient) GetSystemDomains(ctx context.Context) ([]string, error) {
	if c.schema.GetSystemDomains != nil {
		return c.schema.GetSystemDomains(ctx, c.transport)
	} else {
		return []string{}, c.notImplemented("GetSystemDomains")
	}
}

func (c *PluggableClient) GetApplication(ctx context.Context, appId masherytypes.ApplicationIdentifier) (*masherytypes.Application, error) {
	if c.schema.GetApplicationContext != nil {
		return c.schema.GetApplicationContext(ctx, appId, c.transport)
	} else {
		return nil, c.notImplemented("GetApplication")
	}
}

func (c *PluggableClient) GetApplicationPackageKeys(ctx context.Context, appId masherytypes.ApplicationIdentifier) ([]masherytypes.PackageKey, error) {
	if c.schema.GetApplicationContext != nil {
		return c.schema.GetApplicationPackageKeys(ctx, appId, c.transport)
	} else {
		return nil, c.notImplemented("GetApplicationPackageKeys")
	}
}

func (c *PluggableClient) CountApplicationPackageKeys(ctx context.Context, appId masherytypes.ApplicationIdentifier) (int64, error) {
	if c.schema.CountApplicationPackageKeys != nil {
		return c.schema.CountApplicationPackageKeys(ctx, appId, c.transport)
	} else {
		return 0, c.notImplemented("CountApplicationPackageKeys")
	}
}

func (c *PluggableClient) GetFullApplication(ctx context.Context, id masherytypes.ApplicationIdentifier) (*masherytypes.Application, error) {
	if c.schema.GetFullApplication != nil {
		return c.schema.GetFullApplication(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetFullApplication")
	}
}

func (c *PluggableClient) CreateApplication(ctx context.Context, app masherytypes.Application, memberId masherytypes.MemberIdentifier) (*masherytypes.Application, error) {
	if c.schema.CreateApplication != nil {
		return c.schema.CreateApplication(ctx, app, memberId, c.transport)
	} else {
		return nil, c.notImplemented("CreateApplication")
	}
}

func (c *PluggableClient) UpdateApplication(ctx context.Context, app masherytypes.Application) (*masherytypes.Application, error) {
	if c.schema.UpdateApplication != nil {
		return c.schema.UpdateApplication(ctx, app, c.transport)
	} else {
		return nil, c.notImplemented("UpdateApplication")
	}
}

func (c *PluggableClient) DeleteApplication(ctx context.Context, appId masherytypes.ApplicationIdentifier) error {
	if c.schema.DeleteApplication != nil {
		return c.schema.DeleteApplication(ctx, appId, c.transport)
	} else {
		return c.notImplemented("DeleteApplication")
	}
}

func (c *PluggableClient) CountApplicationsOfMember(ctx context.Context, memberId masherytypes.MemberIdentifier) (int64, error) {
	if c.schema.CountApplicationsOfMember != nil {
		return c.schema.CountApplicationsOfMember(ctx, memberId, c.transport)
	} else {
		return 0, c.notImplemented("DeleteApplication")
	}
}

func (c *PluggableClient) ListApplications(ctx context.Context) ([]masherytypes.Application, error) {
	if c.schema.ListApplications != nil {
		return c.schema.ListApplications(ctx, c.transport)
	} else {
		return []masherytypes.Application{}, c.notImplemented("ListApplications")
	}
}

// -----------------------------------------------------------------------------------------------------------------
// Email template set
// -----------------------------------------------------------------------------------------------------------------

func (c *PluggableClient) GetEmailTemplateSet(ctx context.Context, id string) (*masherytypes.EmailTemplateSet, error) {
	if c.schema.GetEmailTemplateSet != nil {
		return c.schema.GetEmailTemplateSet(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetEmailTemplateSet")
	}
}

func (c *PluggableClient) ListEmailTemplateSets(ctx context.Context) ([]masherytypes.EmailTemplateSet, error) {
	if c.schema.ListEmailTemplateSets != nil {
		return c.schema.ListEmailTemplateSets(ctx, c.transport)
	} else {
		return nil, c.notImplemented("ListEmailTemplateSets")
	}
}

func (c *PluggableClient) ListEmailTemplateSetsFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.EmailTemplateSet, error) {
	if c.schema.ListEmailTemplateSetsFiltered != nil {
		return c.schema.ListEmailTemplateSetsFiltered(ctx, params, fields, c.transport)
	} else {
		return nil, c.notImplemented("ListEmailTemplateSetsFiltered")
	}
}

// -----------------------------------------------------------------------------------------------------------------
// Endpoints
// -----------------------------------------------------------------------------------------------------------------

func (c *PluggableClient) ListEndpoints(ctx context.Context, serviceId masherytypes.ServiceIdentifier) ([]masherytypes.AddressableV3Object, error) {
	if c.schema.ListEndpoints != nil {
		return c.schema.ListEndpoints(ctx, serviceId, c.transport)
	} else {
		return nil, c.notImplemented("ListEndpoints")
	}
}

func (c *PluggableClient) ListEndpointsWithFullInfo(ctx context.Context, serviceId masherytypes.ServiceIdentifier) ([]masherytypes.Endpoint, error) {
	if c.schema.ListEndpointsWithFullInfo != nil {
		return c.schema.ListEndpointsWithFullInfo(ctx, serviceId, c.transport)
	} else {
		return nil, c.notImplemented("ListEndpointsWithFullInfo")
	}
}

func (c *PluggableClient) CreateEndpoint(ctx context.Context, serviceId masherytypes.ServiceIdentifier, endp masherytypes.Endpoint) (*masherytypes.Endpoint, error) {
	if c.schema.CreateEndpoint != nil {
		return c.schema.CreateEndpoint(ctx, serviceId, endp, c.transport)
	} else {
		return nil, c.notImplemented("CreateEndpoint")
	}
}

func (c *PluggableClient) UpdateEndpoint(ctx context.Context, endp masherytypes.Endpoint) (*masherytypes.Endpoint, error) {
	if c.schema.UpdateEndpoint != nil {
		return c.schema.UpdateEndpoint(ctx, endp, c.transport)
	} else {
		return nil, c.notImplemented("UpdateEndpoint")
	}
}

func (c *PluggableClient) GetEndpoint(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) (*masherytypes.Endpoint, error) {
	if c.schema.GetEndpoint != nil {
		return c.schema.GetEndpoint(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("GetEndpoint")
	}
}

func (c *PluggableClient) DeleteEndpoint(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) error {
	if c.schema.DeleteEndpoint != nil {
		return c.schema.DeleteEndpoint(ctx, ident, c.transport)
	} else {
		return c.notImplemented("DeleteEndpoint")
	}
}

func (c *PluggableClient) CountEndpointsOf(ctx context.Context, serviceId masherytypes.ServiceIdentifier) (int64, error) {
	if c.schema.CountEndpointsOf != nil {
		return c.schema.CountEndpointsOf(ctx, serviceId, c.transport)
	} else {
		return 0, c.notImplemented("CountEndpointsOf")
	}
}

// -------------------------------------
// Endpoint methods

func (c *PluggableClient) ListEndpointMethods(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) ([]masherytypes.ServiceEndpointMethod, error) {
	if c.schema.ListEndpointMethods != nil {
		return c.schema.ListEndpointMethods(ctx, ident, c.transport)
	} else {
		return []masherytypes.ServiceEndpointMethod{}, c.notImplemented("ListEndpointMethods")
	}
}

func (c *PluggableClient) ListEndpointMethodsWithFullInfo(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) ([]masherytypes.ServiceEndpointMethod, error) {
	if c.schema.ListEndpointMethodsWithFullInfo != nil {
		return c.schema.ListEndpointMethodsWithFullInfo(ctx, ident, c.transport)
	} else {
		return []masherytypes.ServiceEndpointMethod{}, c.notImplemented("ListEndpointMethodsWithFullInfo")
	}
}

func (c *PluggableClient) CreateEndpointMethod(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier, methoUpsert masherytypes.ServiceEndpointMethod) (*masherytypes.ServiceEndpointMethod, error) {
	if c.schema.CreateEndpointMethod != nil {
		return c.schema.CreateEndpointMethod(ctx, ident, methoUpsert, c.transport)
	} else {
		return nil, c.notImplemented("CreateEndpointMethod")
	}
}

func (c *PluggableClient) UpdateEndpointMethod(ctx context.Context, methUpsert masherytypes.ServiceEndpointMethod) (*masherytypes.ServiceEndpointMethod, error) {
	if c.schema.UpdateEndpointMethod != nil {
		return c.schema.UpdateEndpointMethod(ctx, methUpsert, c.transport)
	} else {
		return nil, c.notImplemented("UpdateEndpointMethod")
	}
}

func (c *PluggableClient) GetEndpointMethod(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) (*masherytypes.ServiceEndpointMethod, error) {
	if c.schema.GetEndpointMethod != nil {
		return c.schema.GetEndpointMethod(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("GetEndpointMethod")
	}
}

func (c *PluggableClient) DeleteEndpointMethod(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) error {
	if c.schema.DeleteEndpointMethod != nil {
		return c.schema.DeleteEndpointMethod(ctx, ident, c.transport)
	} else {
		return c.notImplemented("DeleteEndpointMethod")
	}
}

func (c *PluggableClient) CountEndpointsMethodsOf(ctx context.Context, ident masherytypes.ServiceEndpointIdentifier) (int64, error) {
	if c.schema.CountEndpointsMethodsOf != nil {
		return c.schema.CountEndpointsMethodsOf(ctx, ident, c.transport)
	} else {
		return 0, c.notImplemented("CountEndpointsMethodsOf")
	}
}

// -----------------------------------------------------------------------------------
// Endpoint method filters.

func (c *PluggableClient) ListEndpointMethodFilters(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) ([]masherytypes.ServiceEndpointMethodFilter, error) {
	if c.schema.ListEndpointMethodFilters != nil {
		return c.schema.ListEndpointMethodFilters(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("ListEndpointMethodFilters")
	}
}

func (c *PluggableClient) ListEndpointMethodFiltersWithFullInfo(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) ([]masherytypes.ServiceEndpointMethodFilter, error) {
	if c.schema.ListEndpointMethodFiltersWithFullInfo != nil {
		return c.schema.ListEndpointMethodFiltersWithFullInfo(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("ListEndpointMethodFiltersWithFullInfo")
	}
}

func (c *PluggableClient) CreateEndpointMethodFilter(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier, filterUpsert masherytypes.ServiceEndpointMethodFilter) (*masherytypes.ServiceEndpointMethodFilter, error) {
	if c.schema.CreateEndpointMethodFilter != nil {
		return c.schema.CreateEndpointMethodFilter(ctx, ident, filterUpsert, c.transport)
	} else {
		return nil, c.notImplemented("CreateEndpointMethodFilter")
	}
}

func (c *PluggableClient) UpdateEndpointMethodFilter(ctx context.Context, filterUpsert masherytypes.ServiceEndpointMethodFilter) (*masherytypes.ServiceEndpointMethodFilter, error) {
	if c.schema.UpdateEndpointMethodFilter != nil {
		return c.schema.UpdateEndpointMethodFilter(ctx, filterUpsert, c.transport)
	} else {
		return nil, c.notImplemented("UpdateEndpointMethodFilter")
	}
}

func (c *PluggableClient) GetEndpointMethodFilter(ctx context.Context, ident masherytypes.ServiceEndpointMethodFilterIdentifier) (*masherytypes.ServiceEndpointMethodFilter, error) {
	if c.schema.GetEndpointMethodFilter != nil {
		return c.schema.GetEndpointMethodFilter(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("GetEndpointMethodFilter")
	}
}

func (c *PluggableClient) DeleteEndpointMethodFilter(ctx context.Context, ident masherytypes.ServiceEndpointMethodFilterIdentifier) error {
	if c.schema.DeleteEndpointMethodFilter != nil {
		return c.schema.DeleteEndpointMethodFilter(ctx, ident, c.transport)
	} else {
		return c.notImplemented("DeleteEndpointMethodFilter")
	}
}

func (c *PluggableClient) CountEndpointsMethodsFiltersOf(ctx context.Context, ident masherytypes.ServiceEndpointMethodIdentifier) (int64, error) {
	if c.schema.CountEndpointsMethodsFiltersOf != nil {
		return c.schema.CountEndpointsMethodsFiltersOf(ctx, ident, c.transport)
	} else {
		return 0, c.notImplemented("CountEndpointsMethodsFiltersOf")
	}
}

func (c *PluggableClient) GetMember(ctx context.Context, id masherytypes.MemberIdentifier) (*masherytypes.Member, error) {
	if c.schema.GetMember != nil {
		return c.schema.GetMember(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetMember")
	}
}

func (c *PluggableClient) GetFullMember(ctx context.Context, id masherytypes.MemberIdentifier) (*masherytypes.Member, error) {
	if c.schema.GetFullMember != nil {
		return c.schema.GetFullMember(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetFullMember")
	}
}

func (c *PluggableClient) CreateMember(ctx context.Context, member masherytypes.Member) (*masherytypes.Member, error) {
	if c.schema.CreateMember != nil {
		return c.schema.CreateMember(ctx, member, c.transport)
	} else {
		return nil, c.notImplemented("CreateMember")
	}
}

func (c *PluggableClient) UpdateMember(ctx context.Context, member masherytypes.Member) (*masherytypes.Member, error) {
	if c.schema.UpdateMember != nil {
		return c.schema.UpdateMember(ctx, member, c.transport)
	} else {
		return nil, c.notImplemented("UpdateMember")
	}
}

func (c *PluggableClient) DeleteMember(ctx context.Context, memberId masherytypes.MemberIdentifier) error {
	if c.schema.DeleteMember != nil {
		return c.schema.DeleteMember(ctx, memberId, c.transport)
	} else {
		return c.notImplemented("DeleteMember")
	}
}

func (c *PluggableClient) ListMembers(ctx context.Context) ([]masherytypes.Member, error) {
	if c.schema.ListMembers != nil {
		return c.schema.ListMembers(ctx, c.transport)
	} else {
		return []masherytypes.Member{}, c.notImplemented("ListMembers")
	}
}

// ---------------------------------------------
// Packages

func (c *PluggableClient) GetPackage(ctx context.Context, id masherytypes.PackageIdentifier) (*masherytypes.Package, error) {
	if c.schema.GetPackage != nil {
		return c.schema.GetPackage(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetPackage")
	}
}

func (c *PluggableClient) CreatePackage(ctx context.Context, pack masherytypes.Package) (*masherytypes.Package, error) {
	if c.schema.CreatePackage != nil {
		return c.schema.CreatePackage(ctx, pack, c.transport)
	} else {
		return nil, c.notImplemented("CreatePackage")
	}
}

func (c *PluggableClient) UpdatePackage(ctx context.Context, pack masherytypes.Package) (*masherytypes.Package, error) {
	if c.schema.UpdatePackage != nil {
		return c.schema.UpdatePackage(ctx, pack, c.transport)
	} else {
		return nil, c.notImplemented("UpdatePackage")
	}
}

func (c *PluggableClient) DeletePackage(ctx context.Context, packId masherytypes.PackageIdentifier) error {
	if c.schema.DeletePackage != nil {
		return c.schema.DeletePackage(ctx, packId, c.transport)
	} else {
		return c.notImplemented("DeletePackage")
	}
}

func (c *PluggableClient) ListPackages(ctx context.Context) ([]masherytypes.Package, error) {
	if c.schema.ListPackages != nil {
		return c.schema.ListPackages(ctx, c.transport)
	} else {
		return []masherytypes.Package{}, c.notImplemented("ListPackages")
	}
}

// --------------------------------------
// Package plans

func (c *PluggableClient) CreatePlanService(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) (*masherytypes.AddressableV3Object, error) {
	if c.schema.CreatePlanService != nil {
		return c.schema.CreatePlanService(ctx, planService, c.transport)
	} else {
		return nil, c.notImplemented("CreatePlanService")
	}
}

func (c *PluggableClient) CheckPlanServiceExists(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) (bool, error) {
	if c.schema.CheckPlanServiceExists != nil {
		return c.schema.CheckPlanServiceExists(ctx, planService, c.transport)
	} else {
		return false, c.notImplemented("CheckPlanServiceExists")
	}
}

func (c *PluggableClient) DeletePlanService(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) error {
	if c.schema.DeletePlanService != nil {
		return c.schema.DeletePlanService(ctx, planService, c.transport)
	} else {
		return c.notImplemented("CreatePlanService")
	}
}

func (c *PluggableClient) CreatePlanEndpoint(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier) (*masherytypes.AddressableV3Object, error) {
	if c.schema.CreatePlanEndpoint != nil {
		return c.schema.CreatePlanEndpoint(ctx, planEndp, c.transport)
	} else {
		return nil, c.notImplemented("CreatePlanEndpoint")
	}
}

func (c *PluggableClient) CheckPlanEndpointExists(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier) (bool, error) {
	if c.schema.CheckPlanEndpointExists != nil {
		return c.schema.CheckPlanEndpointExists(ctx, planEndp, c.transport)
	} else {
		return false, c.notImplemented("CheckPlanEndpointExists")
	}
}

func (c *PluggableClient) DeletePlanEndpoint(ctx context.Context, planEndp masherytypes.PackagePlanServiceEndpointIdentifier) error {
	if c.schema.DeletePlanEndpoint != nil {
		return c.schema.DeletePlanEndpoint(ctx, planEndp, c.transport)
	} else {
		return c.notImplemented("DeletePlanEndpoint")
	}
}

func (c *PluggableClient) ListPlanEndpoints(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) ([]masherytypes.AddressableV3Object, error) {
	if c.schema.ListPlanEndpoints != nil {
		return c.schema.ListPlanEndpoints(ctx, planService, c.transport)
	} else {
		return []masherytypes.AddressableV3Object{}, c.notImplemented("ListPlanEndpoints")
	}
}

func (c *PluggableClient) ListPlanServices(ctx context.Context, ident masherytypes.PackagePlanIdentifier) ([]masherytypes.Service, error) {
	if c.schema.ListPlanServices != nil {
		return c.schema.ListPlanServices(ctx, ident, c.transport)
	} else {
		return []masherytypes.Service{}, c.notImplemented("ListPlanServices")
	}
}

func (c *PluggableClient) ListPlans(ctx context.Context, packageId masherytypes.PackageIdentifier) ([]masherytypes.Plan, error) {
	if c.schema.ListPlans != nil {
		return c.schema.ListPlans(ctx, packageId, c.transport)
	} else {
		return []masherytypes.Plan{}, c.notImplemented("ListPlans")
	}
}

func (c *PluggableClient) CountPlans(ctx context.Context, packageId masherytypes.PackageIdentifier) (int64, error) {
	if c.schema.CountPlans != nil {
		return c.schema.CountPlans(ctx, packageId, c.transport)
	} else {
		return 0, c.notImplemented("CountPlans")
	}
}

func (c *PluggableClient) DeletePlan(ctx context.Context, ident masherytypes.PackagePlanIdentifier) error {
	if c.schema.DeletePlan != nil {
		return c.schema.DeletePlan(ctx, ident, c.transport)
	} else {
		return c.notImplemented("DeletePlan")
	}
}

func (c *PluggableClient) UpdatePlan(ctx context.Context, plan masherytypes.Plan) (*masherytypes.Plan, error) {
	if c.schema.UpdatePlan != nil {
		return c.schema.UpdatePlan(ctx, plan, c.transport)
	} else {
		return nil, c.notImplemented("UpdatePlan")
	}
}

func (c *PluggableClient) CreatePlan(ctx context.Context, packageId masherytypes.PackageIdentifier, plan masherytypes.Plan) (*masherytypes.Plan, error) {
	if c.schema.CreatePlan != nil {
		return c.schema.CreatePlan(ctx, packageId, plan, c.transport)
	} else {
		return nil, c.notImplemented("CreatePlan")
	}
}

func (c *PluggableClient) GetPlan(ctx context.Context, ident masherytypes.PackagePlanIdentifier) (*masherytypes.Plan, error) {
	if c.schema.GetPlan != nil {
		return c.schema.GetPlan(ctx, ident, c.transport)
	} else {
		return nil, c.notImplemented("GetPlan")
	}
}

func (c *PluggableClient) CountPlanService(ctx context.Context, ident masherytypes.PackagePlanIdentifier) (int64, error) {
	if c.schema.CountPlanService != nil {
		return c.schema.CountPlanService(ctx, ident, c.transport)
	} else {
		return 0, c.notImplemented("ListPlanEndpoints")
	}
}

func (c *PluggableClient) CountPlanEndpoints(ctx context.Context, planService masherytypes.PackagePlanServiceIdentifier) (int64, error) {
	if c.schema.CountPlanEndpoints != nil {
		return c.schema.CountPlanEndpoints(ctx, planService, c.transport)
	} else {
		return 0, c.notImplemented("CountPlanEndpoints")
	}
}

//---------------------------------------------
// Package plan methods

func (c *PluggableClient) ListPackagePlanMethods(ctx context.Context, id masherytypes.PackagePlanServiceEndpointIdentifier) ([]masherytypes.PackagePlanServiceEndpointMethod, error) {
	if c.schema.ListPackagePlanMethods != nil {
		return c.schema.ListPackagePlanMethods(ctx, id, c.transport)
	} else {
		return []masherytypes.PackagePlanServiceEndpointMethod{}, c.notImplemented("ListPackagePlanMethods")
	}
}

func (c *PluggableClient) GetPackagePlanMethod(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) (*masherytypes.PackagePlanServiceEndpointMethod, error) {
	if c.schema.GetPackagePlanMethod != nil {
		return c.schema.GetPackagePlanMethod(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetPackagePlanMethod")
	}
}

func (c *PluggableClient) CreatePackagePlanMethod(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) (*masherytypes.PackagePlanServiceEndpointMethod, error) {
	if c.schema.CreatePackagePlanMethod != nil {
		return c.schema.CreatePackagePlanMethod(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("CreatePackagePlanServiceEndpointMethod")
	}
}

func (c *PluggableClient) DeletePackagePlanMethod(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) error {
	if c.schema.DeletePackagePlanMethod != nil {
		return c.schema.DeletePackagePlanMethod(ctx, id, c.transport)
	} else {
		return c.notImplemented("CreatePackagePlanServiceEndpointMethod")
	}
}

// ----------------------------------------
// Plan method filter

func (c *PluggableClient) GetPackagePlanMethodFilter(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) (*masherytypes.PackagePlanServiceEndpointMethodFilter, error) {
	if c.schema.GetPackagePlanMethodFilter != nil {
		return c.schema.GetPackagePlanMethodFilter(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetPackagePlanMethodFilter")
	}
}

func (c *PluggableClient) CreatePackagePlanMethodFilter(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodFilterIdentifier) (*masherytypes.PackagePlanServiceEndpointMethodFilter, error) {
	if c.schema.CreatePackagePlanMethodFilter != nil {
		return c.schema.CreatePackagePlanMethodFilter(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("CreatePackagePlanMethodFilter")
	}
}

func (c *PluggableClient) DeletePackagePlanMethodFilter(ctx context.Context, id masherytypes.PackagePlanServiceEndpointMethodIdentifier) error {
	if c.schema.DeletePackagePlanMethodFilter != nil {
		return c.schema.DeletePackagePlanMethodFilter(ctx, id, c.transport)
	} else {
		return c.notImplemented("DeletePackagePlanMethodFilter")
	}
}

// ------------------------------------------------------------
// Package key

func (c *PluggableClient) GetPackageKey(ctx context.Context, id masherytypes.PackageKeyIdentifier) (*masherytypes.PackageKey, error) {
	if c.schema.GetPackageKey != nil {
		return c.schema.GetPackageKey(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetPackageKey")
	}
}

func (c *PluggableClient) CreatePackageKey(ctx context.Context, appId masherytypes.ApplicationIdentifier, packageKey masherytypes.PackageKey) (*masherytypes.PackageKey, error) {
	if c.schema.CreatePackageKey != nil {
		return c.schema.CreatePackageKey(ctx, appId, packageKey, c.transport)
	} else {
		return nil, c.notImplemented("CreatePackageKey")
	}
}

func (c *PluggableClient) UpdatePackageKey(ctx context.Context, packageKey masherytypes.PackageKey) (*masherytypes.PackageKey, error) {
	if c.schema.UpdatePackageKey != nil {
		return c.schema.UpdatePackageKey(ctx, packageKey, c.transport)
	} else {
		return nil, c.notImplemented("UpdatePackageKey")
	}
}

func (c *PluggableClient) DeletePackageKey(ctx context.Context, keyId masherytypes.PackageKeyIdentifier) error {
	if c.schema.DeletePackageKey != nil {
		return c.schema.DeletePackageKey(ctx, keyId, c.transport)
	} else {
		return c.notImplemented("DeletePackageKey")
	}
}

func (c *PluggableClient) ListPackageKeysFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.PackageKey, error) {
	if c.schema.ListPackageKeysFiltered != nil {
		return c.schema.ListPackageKeysFiltered(ctx, params, fields, c.transport)
	} else {
		return []masherytypes.PackageKey{}, c.notImplemented("ListPackageKeysFiltered")
	}
}

func (c *PluggableClient) ListPackageKeys(ctx context.Context) ([]masherytypes.PackageKey, error) {
	if c.schema.ListPackageKeys != nil {
		return c.schema.ListPackageKeys(ctx, c.transport)
	} else {
		return []masherytypes.PackageKey{}, c.notImplemented("ListPackageKeys")
	}
}

// ---------------------
// Roles

func (c *PluggableClient) GetRole(ctx context.Context, id string) (*masherytypes.Role, error) {
	if c.schema.GetRole != nil {
		return c.schema.GetRole(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetRole")
	}
}

func (c *PluggableClient) ListRoles(ctx context.Context) ([]masherytypes.Role, error) {
	if c.schema.ListRoles != nil {
		return c.schema.ListRoles(ctx, c.transport)
	} else {
		return []masherytypes.Role{}, c.notImplemented("ListRoles")
	}
}

func (c *PluggableClient) ListRolesFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.Role, error) {
	if c.schema.ListRolesFiltered != nil {
		return c.schema.ListRolesFiltered(ctx, params, fields, c.transport)
	} else {
		return []masherytypes.Role{}, c.notImplemented("ListRolesFiltered")
	}
}

// ------------------------------
// Service

func (c *PluggableClient) GetService(ctx context.Context, id masherytypes.ServiceIdentifier) (*masherytypes.Service, error) {
	if c.schema.GetService != nil {
		return c.schema.GetService(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetService")
	}
}

func (c *PluggableClient) CreateService(ctx context.Context, service masherytypes.Service) (*masherytypes.Service, error) {
	if c.schema.CreateService != nil {
		return c.schema.CreateService(ctx, service, c.transport)
	} else {
		return nil, c.notImplemented("CreateService")
	}
}

func (c *PluggableClient) UpdateService(ctx context.Context, service masherytypes.Service) (*masherytypes.Service, error) {
	if c.schema.UpdateService != nil {
		return c.schema.UpdateService(ctx, service, c.transport)
	} else {
		return nil, c.notImplemented("UpdateService")
	}
}

func (c *PluggableClient) DeleteService(ctx context.Context, serviceId masherytypes.ServiceIdentifier) error {
	if c.schema.DeleteService != nil {
		return c.schema.DeleteService(ctx, serviceId, c.transport)
	} else {
		return c.notImplemented("UpdateService")
	}
}

func (c *PluggableClient) ListServicesFiltered(ctx context.Context, params map[string]string, fields []string) ([]masherytypes.Service, error) {
	if c.schema.ListServicesFiltered != nil {
		return c.schema.ListServicesFiltered(ctx, params, fields, c.transport)
	} else {
		return []masherytypes.Service{}, c.notImplemented("ListServicesFiltered")
	}
}

func (c *PluggableClient) ListServices(ctx context.Context) ([]masherytypes.Service, error) {
	if c.schema.ListServices != nil {
		return c.schema.ListServices(ctx, c.transport)
	} else {
		return []masherytypes.Service{}, c.notImplemented("ListServices")
	}
}

func (c *PluggableClient) CountServices(ctx context.Context, params map[string]string) (int64, error) {
	if c.schema.CountServices != nil {
		return c.schema.CountServices(ctx, params, c.transport)
	} else {
		return 0, c.notImplemented("CountServices")
	}
}

// --------------------------------
// Service cache

func (c *PluggableClient) GetServiceCache(ctx context.Context, id string) (*masherytypes.ServiceCache, error) {
	if c.schema.GetServiceCache != nil {
		return c.schema.GetServiceCache(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetServiceCache")
	}
}

func (c *PluggableClient) CreateServiceCache(ctx context.Context, id string, service masherytypes.ServiceCache) (*masherytypes.ServiceCache, error) {
	if c.schema.CreateServiceCache != nil {
		return c.schema.CreateServiceCache(ctx, id, service, c.transport)
	} else {
		return nil, c.notImplemented("CreateServiceCache")
	}
}

func (c *PluggableClient) UpdateServiceCache(ctx context.Context, id string, service masherytypes.ServiceCache) (*masherytypes.ServiceCache, error) {
	if c.schema.UpdateServiceCache != nil {
		return c.schema.UpdateServiceCache(ctx, id, service, c.transport)
	} else {
		return nil, c.notImplemented("UpdateServiceCache")
	}
}

func (c *PluggableClient) DeleteServiceCache(ctx context.Context, id string) error {
	if c.schema.DeleteServiceCache != nil {
		return c.schema.DeleteServiceCache(ctx, id, c.transport)
	} else {
		return c.notImplemented("DeleteServiceCache")
	}
}

// Service OAtuh
func (c *PluggableClient) GetServiceOAuthSecurityProfile(ctx context.Context, id masherytypes.ServiceIdentifier) (*masherytypes.MasheryOAuth, error) {
	if c.schema.GetServiceOAuthSecurityProfile != nil {
		return c.schema.GetServiceOAuthSecurityProfile(ctx, id, c.transport)
	} else {
		return nil, c.notImplemented("GetServiceOAuthSecurityProfile")
	}
}

func (c *PluggableClient) CreateServiceOAuthSecurityProfile(ctx context.Context, service masherytypes.MasheryOAuth) (*masherytypes.MasheryOAuth, error) {
	if c.schema.CreateServiceOAuthSecurityProfile != nil {
		return c.schema.CreateServiceOAuthSecurityProfile(ctx, service, c.transport)
	} else {
		return nil, c.notImplemented("CreateServiceOAuthSecurityProfile")
	}
}

func (c *PluggableClient) UpdateServiceOAuthSecurityProfile(ctx context.Context, service masherytypes.MasheryOAuth) (*masherytypes.MasheryOAuth, error) {
	if c.schema.UpdateServiceOAuthSecurityProfile != nil {
		return c.schema.UpdateServiceOAuthSecurityProfile(ctx, service, c.transport)
	} else {
		return nil, c.notImplemented("UpdateServiceOAuthSecurityProfile")
	}
}

func (c *PluggableClient) DeleteServiceOAuthSecurityProfile(ctx context.Context, id masherytypes.ServiceIdentifier) error {
	if c.schema.DeleteServiceOAuthSecurityProfile != nil {
		return c.schema.DeleteServiceOAuthSecurityProfile(ctx, id, c.transport)
	} else {
		return c.notImplemented("DeleteServiceOAuthSecurityProfile")
	}
}
