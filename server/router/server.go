package router

import (
	"context"
	"fmt"
	"net/http"

	goswaggermiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/layer5io/meshery/server/handlers"
	"github.com/layer5io/meshery/server/models"
)

// Router represents Meshery router
type Router struct {
	Echo *echo.Echo
	port int
}

/*
// NewRouter returns a new ServeMux with app routes.
func NewRouter(_ context.Context, h models.HandlerInterface, port int, g http.Handler, gp http.Handler) *Router {
	gMux := mux.NewRouter()

	gMux.Handle("/api/system/graphql/query", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.MesheryControllersMiddleware(h.GraphqlMiddleware(g)))), models.ProviderAuth))).Methods("GET", "POST")
	gMux.Handle("/api/system/graphql/playground", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.MesheryControllersMiddleware(h.GraphqlMiddleware(gp)))), models.ProviderAuth))).Methods("GET", "POST")

	gMux.Handle("/api/system/database", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetSystemDatabase), models.ProviderAuth))).
		Methods("GET")

	gMux.Handle("/api/user/prefs", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UserPrefsHandler), models.ProviderAuth))).
		Methods("GET", "POST")

	gMux.Handle("/api/user/prefs/perf", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UserTestPreferenceHandler), models.ProviderAuth))).
		Methods("GET", "POST", "DELETE")

	gMux.Handle("/api/system/kubernetes", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.K8SConfigHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/system/kubernetes/ping", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesPingHandler), models.ProviderAuth))).
		Methods("GET")

	gMux.Handle("/api/system/kubernetes/contexts", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetContextsFromK8SConfig), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/system/kubernetes/register", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.K8sRegistrationHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/system/kubernetes/contexts", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetAllContexts), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/system/kubernetes/contexts/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetContext), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/system/kubernetes/contexts/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteContext), models.ProviderAuth))).
		Methods("DELETE")

	gMux.Handle("/api/perf/profile", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.LoadTestHandler), models.ProviderAuth))).
		Methods("GET", "POST")
	gMux.Handle("/api/perf/profile/result", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.FetchAllResultsHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/perf/profile/result/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetResultHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/mesh", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetSMPServiceMeshes), models.ProviderAuth))).
		Methods("GET")

	gMux.Handle("/api/smi/results", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.FetchSmiResultsHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/smi/results/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.FetchSingleSmiResultHandler), models.ProviderAuth))).
		Methods("GET")

	gMux.Handle("/api/system/adapter/manage", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.MeshAdapterConfigHandler)), models.ProviderAuth)))
	gMux.Handle("/api/system/adapter/operation", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.MeshOpsHandler)), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/system/adapters", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.AdaptersHandler), models.ProviderAuth)))
	gMux.Handle("/api/system/availableAdapters", http.HandlerFunc(h.AvailableAdaptersHandler)).
		Methods("GET")

	gMux.Handle("/api/events", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.EventStreamHandler), models.ProviderAuth))).
		Methods("GET")

	gMux.Handle("/api/telemetry/metrics/grafana/config", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GrafanaConfigHandler), models.ProviderAuth))).
		Methods("GET", "POST", "DELETE")
	gMux.Handle("/api/telemetry/metrics/grafana/boards", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GrafanaBoardsHandler), models.ProviderAuth))).
		Methods("GET", "POST")
	gMux.Handle("/api/telemetry/metrics/grafana/query", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GrafanaQueryHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/grafana/query_range", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GrafanaQueryRangeHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/telemetry/metrics/grafana/ping", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GrafanaPingHandler), models.ProviderAuth))).
		Methods("GET")
	// gMux.Handle("/api/telemetry/metrics/grafana/scan", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.ScanGrafanaHandler)), models.ProviderAuth)))

	gMux.Handle("/api/telemetry/metrics/config", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PrometheusConfigHandler), models.ProviderAuth))).
		Methods("GET", "POST", "DELETE")
	gMux.Handle("/api/telemetry/metrics/board_import", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GrafanaBoardImportForPrometheusHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/telemetry/metrics/query", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PrometheusQueryHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/prometheus/query_range", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.PrometheusQueryRangeHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/telemetry/metrics/ping", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PrometheusPingHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/telemetry/metrics/static-board", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PrometheusStaticBoardHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/telemetry/metrics/boards", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.SaveSelectedPrometheusBoardsHandler), models.ProviderAuth))).
		Methods("POST")
	// gMux.Handle("/api/system/meshsync/prometheus", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.ScanPrometheusHandler)), models.ProviderAuth)))

	// gMux.Handle("/api/system/meshsync/grafana", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.ScanPromGrafanaHandler)), models.ProviderAuth)))

	gMux.Handle("/api/pattern/deploy", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.PatternFileHandler)), models.ProviderAuth))).
		Methods("POST", "DELETE")
	gMux.Handle("/api/pattern", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PatternFileRequestHandler), models.ProviderAuth))).
		Methods("POST", "GET")
	gMux.Handle("/api/pattern/catalog", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetCatalogMesheryPatternsHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/pattern/catalog/unpublish", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UnPublishCatalogPatternHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/pattern/catalog/publish", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PublishCatalogPatternHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/pattern/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryPatternHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/pattern/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteMesheryPatternHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/pattern/clone/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.CloneMesheryPatternHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/pattern/download/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DownloadMesheryPatternHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/patterns/delete", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteMultiMesheryPatternsHandler), models.ProviderAuth))).
		Methods("POST")

	gMux.Handle("/api/meshmodels/validate", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.ValidationHandler), models.NoAuth))).Methods("POST")
	gMux.Handle("/api/meshmodels/components", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.RegisterMeshmodelComponents), models.NoAuth))).Methods("POST")         //This should also be left with NoAuth
	gMux.Handle("/api/meshmodel/components/register", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.RegisterMeshmodelComponents), models.NoAuth))).Methods("POST") //For backwards compatibility with previous registrants
	gMux.Handle("/api/meshmodels/components", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetAllMeshmodelComponents), models.NoAuth))).Methods("GET")

	gMux.Handle("/api/meshmodels/categories", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelCategories), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/models", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelModels), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/models/{model}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelModelsByName), models.NoAuth))).Methods("GET")

	gMux.Handle("/api/meshmodels/categories/{category}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelCategoriesByName), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/categories/{category}/models", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelModelsByCategories), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/categories/{category}/models/{model}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelModelsByCategoriesByModel), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/categories/{category}/models/{model}/components", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelComponentByModelByCategory), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/categories/{category}/models/{model}/components/{name}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelComponentsByNameByModelByCategory), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/categories/{category}/components", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelComponentByCategory), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/categories/{category}/components/{name}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelComponentsByNameByCategory), models.NoAuth))).Methods("GET")

	gMux.Handle("/api/meshmodels/components/{name}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetAllMeshmodelComponentsByName), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/generate", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.MeshModelGenerationHandler), models.NoAuth))).Methods("POST")
	gMux.Handle("/api/meshmodels/relationships", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetAllMeshmodelRelationships), models.NoAuth))).Methods("GET")

	gMux.Handle("/api/meshmodels/models/{model}/components", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelComponentByModel), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/models/{model}/components/{name}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelComponentsByNameByModel), models.NoAuth))).Methods("GET")

	gMux.Handle("/api/meshmodels/models/{model}/relationships", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetAllMeshmodelRelationships), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/models/{model}/relationships/{name}", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetMeshmodelRelationshipByName), models.NoAuth))).Methods("GET")
	gMux.Handle("/api/meshmodels/relationships", h.ProviderMiddleware(h.AuthMiddleware(http.HandlerFunc(h.RegisterMeshmodelRelationships), models.NoAuth))).Methods("POST") //This should also be left with NoAuth

	gMux.Handle("/api/filter/deploy", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.FilterFileHandler)), models.ProviderAuth))).
		Methods("POST", "DELETE")
	gMux.Handle("/api/filter", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.FilterFileRequestHandler), models.ProviderAuth))).
		Methods("POST", "GET")
	gMux.Handle("/api/filter/catalog", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetCatalogMesheryFiltersHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/filter/catalog/publish", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.PublishCatalogFilterHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/filter/catalog/unpublish", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UnPublishCatalogFilterHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/filter/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryFilterHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/filter/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteMesheryFilterHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/filter/download/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryFilterFileHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/filter/clone/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.CloneMesheryFilterHandler), models.ProviderAuth))).
		Methods("POST")

	gMux.Handle("/api/application/deploy", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.ApplicationFileHandler)), models.ProviderAuth))).
		Methods("POST", "DELETE")
	gMux.Handle("/api/application", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.ApplicationFileRequestHandler), models.ProviderAuth))).
		Methods("GET", "POST")
	gMux.Handle("/api/application/{sourcetype}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.ApplicationFileRequestHandler), models.ProviderAuth))).
		Methods("POST", "PUT")
	gMux.Handle("/api/application/types", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryApplicationTypesHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/application/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryApplicationHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/application/download/{id}/{sourcetype}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryApplicationSourceHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/application/download/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetMesheryApplicationFile), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/application/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteMesheryApplicationHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/content/design/share", h.ProviderMiddleware((h.AuthMiddleware(h.SessionInjectorMiddleware(h.ShareDesignHandler), models.ProviderAuth)))).
		Methods("POST")

	gMux.Handle("/api/user/performance/profiles", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetPerformanceProfilesHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/user/performance/profiles/results", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.FetchAllResultsHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/user/performance/profiles/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetPerformanceProfileHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/user/performance/profiles/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeletePerformanceProfileHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/user/performance/profiles", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.SavePerformanceProfileHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/user/performance/profiles/{id}/run", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.KubernetesMiddleware(h.LoadTestHandler)), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/user/performance/profiles/{id}/results", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.FetchResultsHandler), models.ProviderAuth))).
		Methods("GET")

	gMux.Handle("/api/user/schedules", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetSchedulesHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/user/schedules/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetScheduleHandler), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/user/schedules/{id}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteScheduleHandler), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/user/schedules", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.SaveScheduleHandler), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/policies/run_policy", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetRegoPolicyForDesignFile), models.ProviderAuth))).
		Methods("POST")

	// Handlers for User Credentials

	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetUserCredentials), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UpdateUserCredential), models.ProviderAuth))).
		Methods("PUT")
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteUserCredential), models.ProviderAuth))).
		Methods("DELETE")
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.SaveUserCredential), models.ProviderAuth))).
		Methods("POST")

	gMux.PathPrefix("/api/extensions").
		Handler(h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.ExtensionsHandler), models.ProviderAuth))).
		Methods("GET", "POST", "OPTIONS", "PUT", "DELETE")

	//gMux.PathPrefix("/api/system/graphql").Handler(h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GraphqlSystemHandler)))).Methods("GET", "POST")

	gMux.Handle("/api/token", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(
		func(w http.ResponseWriter, req *http.Request, _ *models.Preference, _ *models.User, provider models.Provider) {
			provider.ExtractToken(w, req)
		}), models.ProviderAuth))).Methods("GET")

	// TODO: have to change this too
	gMux.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=3600") // 1 hr
		http.ServeFile(w, r, "../ui/out/static/img/meshery-logo/meshery-logo.svg")
	}))

	// Handlers for User Credentials
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.SaveUserCredential), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetUserCredentials), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UpdateUserCredential), models.ProviderAuth))).
		Methods("PUT")
	gMux.Handle("/api/integrations/credentials", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteUserCredential), models.ProviderAuth))).
		Methods("DELETE")

	// Handlers for User Connections
	gMux.Handle("/api/integrations/connections", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.SaveConnection), models.ProviderAuth))).
		Methods("POST")
	gMux.Handle("/api/integrations/connections/{connectionKind}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.GetConnections), models.ProviderAuth))).
		Methods("GET")
	gMux.Handle("/api/integrations/connections/{connectionKind}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.UpdateConnection), models.ProviderAuth))).
		Methods("PUT")
	gMux.Handle("/api/integrations/connections/{connectionId}", h.ProviderMiddleware(h.AuthMiddleware(h.SessionInjectorMiddleware(h.DeleteConnection), models.ProviderAuth))).
		Methods("DELETE")

	return &Router{
		S:    gMux,
		port: port,
	}
}
*/

func NewRouter(_ context.Context, h models.HandlerInterface, port int, g http.Handler, gp http.Handler) *Router {
	e := echo.New()

	e.Use(middleware.Recover())

	// Group common routes under "/api" prefix
	apiGroup := e.Group("/api")
	apiGroup.Use(echo.WrapMiddleware(h.ProviderMiddleware))
	apiGroup.Use(echo.WrapMiddleware(h.AuthMiddleware))
	// apiGroup.Use(echo.WrapMiddleware(h.SessionInjectorMiddleware))
	// apiGroup.Use(echo.WrapMiddleware(h.KubernetesMiddleware))
	// apiGroup.Use(echo.Middleware(h.MesheryControllersMiddleware))

	// Register GET /api/extension/version route
	// apiGroup.GET("/extension/version", echo.WrapHandler(h.ExtensionsVersionHandler))

	// Register GET /api/session route
	apiGroup.GET("/session", sessionHandler)

	apiGroup.Any("/provider/extension/*", echo.WrapHandler(http.HandlerFunc(h.ProviderComponentsHandler)))

	// Register GET /api/provider/capabilities route
	apiGroup.GET("/provider/capabilities", echo.WrapHandler(http.HandlerFunc(h.ProviderCapabilityHandler)))

	// Register GET /api/identity/users route
	// apiGroup.GET("/identity/users", echo.WrapHandler(http.HandlerFunc(h.GetUsers)))

	// Register GET /api/user route
	apiGroup.GET("/user", echo.WrapHandler(http.HandlerFunc(h.UserHandler)))

	// Register GET /api/user/provider/:id route
	apiGroup.GET("/user/provider/:id", echo.WrapHandler(http.HandlerFunc(h.GetUserByIDHandler)))

	// Register GET /api/system/database
	// apiGroup.GET("/system/database", echo.WrapHandler(http.HandlerFunc(h.GetSystemDatabase)))

	// Register GET /api/system/sync
	apiGroup.GET("/system/sync", echo.WrapHandler(http.HandlerFunc(h.SessionSyncHandler)))

	// Group common routes under "/api/user/token" prefix
	tokenGroup := e.Group("/api/user/token")
	tokenGroup.Use(echo.WrapMiddleware(h.ProviderMiddleware))
	// tokenGroup.Use(echo.WrapMiddleware(h.SessionInjectorMiddleware))

	// Register GET /api/user/token route
	tokenGroup.GET("", tokenHandler)

	// Register POST /api/user/token route
	tokenGroup.GET("", tokenHandler)

	// Group common routes under "/user" prefix
	userGroup := e.Group("/user")
	userGroup.Use(echo.WrapMiddleware(h.ProviderMiddleware))
	userGroup.Use(echo.WrapMiddleware(h.AuthMiddleware))

	userGroup.GET("/login", echo.WrapHandler(http.HandlerFunc(loginHandler)))
	userGroup.GET("/logout", echo.WrapHandler(http.HandlerFunc(logoutHandler)))

	// Routes outside the common group
	// Serve swagger.yaml file
	e.Static("/swagger.yaml", "../helpers/")

	// Serve Swagger UI
	swaggerOpts := goswaggermiddleware.RedocOpts{
		SpecURL: "/swagger.yaml",
		Path:    "/docs",
	}
	redocHandler := goswaggermiddleware.Redoc(swaggerOpts, nil)
	e.GET("/docs", func(c echo.Context) error {
		redocHandler.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	// Register GET / route  - returns static page from Next.js page /
	// e.GET("/", echo.WrapHandler(http.HandlerFunc(h.NextProviderPageHandler)))
	// Serve static files
	e.Static("/ui/public/static/img/meshmodels", "../../ui")
	e.GET("/", func(e echo.Context) error {
		return e.File("../../ui/out/index.html")
	})

	// Register GET /provider route - returns static page from Next.js page /provider
	e.GET("/provider", echo.WrapHandler(http.HandlerFunc(h.ProviderUIHandler)))

	// Register GET /auth/login route
	e.GET("/auth/login", echo.WrapHandler(http.HandlerFunc(h.ProviderUIHandler)))

	// Register GET /api/system/version route
	e.GET("/api/system/version", echo.WrapHandler(http.HandlerFunc(versionHandler)))

	// Register GET /api/provider route
	e.GET("/api/provider", echo.WrapHandler(http.HandlerFunc(h.ProviderHandler)))

	// Register GET /api/providers route
	e.GET("/api/providers", echo.WrapHandler(http.HandlerFunc(h.ProvidersHandler)))

	return &Router{
		Echo: e,
		port: port,
	}
}

// Handler function for "/api/session" route
func sessionHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

// Handler function for "/api/system/version" route
func versionHandler(w http.ResponseWriter, r *http.Request) {
	h := &handlers.Handler{}
	h.ServerVersionHandler(w, r)
}

// Handler function for /user/login requests - will redirect to /provider
func loginHandler(w http.ResponseWriter, r *http.Request) {
	h := &handlers.Handler{}
	providerI := r.Context().Value(models.ProviderCtxKey)
	provider, ok := providerI.(models.Provider)
	if !ok {
		http.Redirect(w, r, "/provider", http.StatusFound)
		return
	}
	h.LoginHandler(w, r, provider, false)
}

// Handler function for /user/logout requests
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	h := &handlers.Handler{}
	providerI := r.Context().Value(models.ProviderCtxKey)
	provider, ok := providerI.(models.Provider)
	if !ok {
		http.Redirect(w, r, "/provider", http.StatusFound)
		return
	}
	h.LogoutHandler(w, r, provider)
}

func tokenHandler(c echo.Context) error {
	h := &handlers.Handler{}
	providerI := c.Request().Context().Value(models.ProviderCtxKey)
	provider, ok := providerI.(models.Provider)
	if !ok {
		return c.Redirect(http.StatusFound, "/provider")
	}

	h.TokenHandler(c.Response().Writer, c.Request(), provider, false)
	return nil
}

// Run starts the http server
func (r *Router) Run() error {
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", r.port),
	// 	Handler:        r.s,
	// 	ReadTimeout:    5 * time.Second,
	// 	WriteTimeout:   2 * time.Minute,
	// 	MaxHeaderBytes: 1 << 20,
	// 	IdleTimeout:    0, //time.Second,
	// }
	// return s.ListenAndServe()
	return http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.Echo)
}
