/*
   nighthawkapi.routes.apiroutes;
*/

package routes

import (
	"fmt"
	"net/http"
	api "nighthawkapi/api/core"

	"nighthawkapi/api/handlers/analyzer"
	"nighthawkapi/api/handlers/audit"
	"nighthawkapi/api/handlers/auth"
	config "nighthawkapi/api/handlers/config"
	"nighthawkapi/api/handlers/delete"
	"nighthawkapi/api/handlers/search"
	"nighthawkapi/api/handlers/stacking"
	"nighthawkapi/api/handlers/upload"
	"nighthawkapi/api/handlers/watcher"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var _api = fmt.Sprintf("%s%s", api.API_NAME, api.API_VER)

var routes = Routes{
	Route{
		"Config",
		"GET",
		fmt.Sprintf("%s/config", _api),
		config.ReturnSystemConfig,
	},
	{
		"Config",
		"POST",
		fmt.Sprintf("%s/config", _api),
		config.UpdateSystemConfig,
	},
	{
		"Stats",
		"GET",
		fmt.Sprintf("%s/platformstats", _api),
		config.ReturnPlatformStats,
	},
	{
		"Upload",
		"POST",
		fmt.Sprintf("%s/upload", _api),
		upload.UploadFileHandler,
	},
	{
		"ListCases",
		"GET",
		fmt.Sprintf("%s/list/cases", _api),
		audit.GetCaseList,
	},
	{
		"ListEndpoints",
		"GET",
		fmt.Sprintf("%s/list/endpoints", _api),
		audit.GetEndpointList,
	},
	{
		"ListAuditType",
		"GET",
		fmt.Sprintf("%s/list/audittypes", _api),
		audit.GetAuditTypeList,
	},
	{
		"ListCompletedJobs",
		"GET",
		fmt.Sprintf("%s/list/completedjobs", _api),
		upload.ListCompletedJobs,
	},
	{
		"ListCompletedJobs",
		"GET",
		fmt.Sprintf("%s/list/jobs/completed", _api),
		upload.ListCompletedJobs,
	},
	{
		"SubscribeJobs",
		"GET",
		fmt.Sprintf("%s/subscribe/uploadjobs", _api),
		upload.SubscribeJobs,
	},
	{
		"SubscribeJobs",
		"GET",
		fmt.Sprintf("%s/subscribe/jobs/upload", _api),
		upload.SubscribeJobs,
	},
	{
		"ShowDocId",
		"GET",
		fmt.Sprintf("%s/show/doc/{doc_id}", _api),
		audit.GetDocById,
	},
	{
		"ShowEndpointByCase",
		"GET",
		fmt.Sprintf("%s/show/{case}", _api),
		audit.GetEndpointByCase,
	},
	{
		"ShowEndpointByCase",
		"GET",
		fmt.Sprintf("%s/show/{case}/endpoints", _api),
		audit.GetEndpointByCase,
	},
	{
		"ShowCasedateByEndpoint",
		"GET",
		fmt.Sprintf("%s/show/{case}/{endpoint}", _api),
		audit.GetCasedateByEndpoint,
	},
	{
		"ShowCasedateByEndpoint",
		"GET",
		fmt.Sprintf("%s/show/{case}/{endpoint}/info", _api),
		audit.GetCasedateByEndpoint,
	},
	{
		"ShowAuditTypeByCaseDate",
		"GET",
		fmt.Sprintf("%s/show/{case}/{endpoint}/{case_date}", _api),
		audit.GetAuditTypeByEndpointAndCase,
	},
	{
		"ShowAuditTypeByCaseDate",
		"GET",
		fmt.Sprintf("%s/show/{case}/{endpoint}/{case_date}/audits", _api),
		audit.GetAuditTypeByEndpointAndCase,
	},
	{
		"ShowAuditDataByAuditType",
		"GET",
		fmt.Sprintf("%s/show/{case}/{endpoint}/{case_date}/{audittype}", _api),
		audit.GetAuditDataByAuditGenerator,
	},
	{
		"ShowAuditDataByAuditType",
		"POST",
		fmt.Sprintf("%s/show/{case}/{endpoint}/{case_date}/{audittype}", _api),
		audit.GetAuditDataByAuditGenerator,
	},
	{
		"StackServices",
		"POST",
		fmt.Sprintf("%s/stacking/service", _api),
		stacking.StackServices,
	},
	{
		"StackPrefetch",
		"POST",
		fmt.Sprintf("%s/stacking/prefetch", _api),
		stacking.StackPrefetch,
	},
	{
		"StackTasks",
		"POST",
		fmt.Sprintf("%s/stacking/task", _api),
		stacking.StackTasks,
	},
	{
		"StackPersisence",
		"POST",
		fmt.Sprintf("%s/stacking/persistence", _api),
		stacking.StackPersistence,
	},
	{
		"StackLocalListenPort",
		"POST",
		fmt.Sprintf("%s/stacking/locallistenport", _api),
		stacking.StackLocalListenPort,
	},
	{
		"StackRunKey",
		"POST",
		fmt.Sprintf("%s/stacking/runkey", _api),
		stacking.StackRunKey,
	},
	{
		"StackDnsARequest",
		"POST",
		fmt.Sprintf("%s/stacking/dns/a", _api),
		stacking.StackDnsARequest,
	},
	{
		"StackUrlDomain",
		"POST",
		fmt.Sprintf("%s/stacking/url/domain", _api),
		stacking.StackUrlDomain,
	},
	{
		"StackContext",
		"POST",
		fmt.Sprintf("%s/stacking/context", _api),
		stacking.GetStackContext,
	},
	{
		"StackContextEndpointDoc",
		"POST",
		fmt.Sprintf("%s/stacking/context/endpoint", _api),
		stacking.GetStackContextByEndpoint,
	},
	{
		"GlobalSearch",
		"POST",
		fmt.Sprintf("%s/search", _api),
		search.GetGlobalSearch,
	},
	{
		"TimelineSearch",
		"POST",
		fmt.Sprintf("%s/search/timeline", _api),
		search.GetTimelineSearch,
	},
	{
		"DeleteCase",
		"GET",
		fmt.Sprintf("%s/delete/case/{case_name}", _api),
		delete.DeleteCase,
	},
	{
		"DeleteEndpoint",
		"GET",
		fmt.Sprintf("%s/delete/endpoint/{endpoint_name}", _api),
		delete.DeleteEndpoint,
	},
	{
		"DeleteCaseEndpoint",
		"GET",
		fmt.Sprintf("%s/delete/{case_name}/{endpoint_name}", _api),
		delete.DeleteCaseEndpoint,
	},
	{
		"DeleteCase",
		"POST",
		fmt.Sprintf("%s/delete/case", _api),
		delete.DeleteCase,
	},
	{
		"DeleteEndpoint",
		"POST",
		fmt.Sprintf("%s/delete/endpoint", _api),
		delete.DeleteEndpoint,
	},
	{
		"GetWatcherResults",
		"GET",
		fmt.Sprintf("%s/watcher/results", _api),
		watcher.GetWatcherResults,
	},
	{
		"GetWatcherMatchById",
		"GET",
		fmt.Sprintf("%s/watcher/results/{id}", _api),
		watcher.GetWatcherResultById,
	},
	{
		"GenWatcherRule",
		"POST",
		fmt.Sprintf("%s/watcher/generate/rule", _api),
		watcher.GenerateWatcherRule,
	},
	{
		"DiffEndpoint",
		"GET",
		fmt.Sprintf("%s/diff/{endpoint}", _api),
		stacking.TimelineEndpointDiff,
	},
	{
		"DiffEndpoint",
		"POST",
		fmt.Sprintf("%s/diff", _api),
		stacking.TimelineEndpointDiff,
	},
	{
		"AddBlacklistItem",
		"POST",
		fmt.Sprintf("%s/analyze/blacklist", _api),
		analyzer.AddBlacklistInformation,
	},
	{
		"AddWhitelistItem",
		"POST",
		fmt.Sprintf("%s/analyze/whitelist", _api),
		analyzer.AddWhitelistInformation,
	},
	{
		"AddStackCommonItem",
		"POST",
		fmt.Sprintf("%s/analyze/stack", _api),
		analyzer.AddStackInformation,
	},
	{
		"DeleteAnalyzerItemByID",
		"GET",
		fmt.Sprintf("%s/analyze/delete/{analyzer_type}/{analyzer_id}", _api),
		analyzer.DeleteAnalyzerItemByID,
	},
	{
		"DeleteAnalyzerItemByQuery",
		"POST",
		fmt.Sprintf("%s/analyze/delete/{analyzer_type}", _api),
		analyzer.DeleteAnalyzerItemByQuery,
	},
	{
		"Login",
		"POST",
		fmt.Sprintf("%s/auth/login", _api),
		auth.Login,
	},
	{
		"Logout",
		"POST",
		fmt.Sprintf("%s/auth/logout", _api),
		auth.Logout,
	},
	{
		"SetPassword",
		"POST",
		fmt.Sprintf("%s/admin/password/set", _api),
		auth.SetPassword,
	},
	{
		"ChangePassword",
		"POST",
		fmt.Sprintf("%s/auth/password/change", _api),
		auth.ChangePassword,
	},
	{
		"Comment::AddComment",
		"POST",
		fmt.Sprintf("%s/comment/add/{casename}/{endpoint}/{audit}/{doc_id}", _api),
		audit.AddComment,
	},
	{
		"GetAllComment",
		"GET",
		fmt.Sprintf("%s/comment/show", _api),
		audit.GetComment,
	},
	{
		"GetCommentByPost",
		"POST",
		fmt.Sprintf("%s/comment/show", _api),
		audit.GetComment,
	},
	{
		"GetCommentByCase",
		"GET",
		fmt.Sprintf("%s/comment/show/{casename}", _api),
		audit.GetComment,
	},
	{
		"GetCommentByCaseEndpoint",
		"GET",
		fmt.Sprintf("%s/comment/show/{casename}/{endpoint}", _api),
		audit.GetComment,
	},
	{
		"GetCommentByCaseEndpointAudit",
		"GET",
		fmt.Sprintf("%s/comment/show/{casename}/{endpoint}/{audit}", _api),
		audit.GetComment,
	},
	{
		"GetCommentByCaseEndpointAuditDocId",
		"GET",
		fmt.Sprintf("%s/comment/show/{casename}/{endpoint}/{audit}/{doc_id}", _api),
		audit.GetComment,
	},
	{
		"AddTag",
		"POST",
		fmt.Sprintf("%s/tag/add/{casename}/{endpoint}/{audit}/{doc_id}", _api),
		audit.AddTag,
	},
	{
		"GetAllTags",
		"GET",
		fmt.Sprintf("%s/tag/show", _api),
		audit.GetTagData,
	},
	{
		"GetTagByPost",
		"POST",
		fmt.Sprintf("%s/tag/show", _api),
		audit.GetTagData,
	},
	{
		"GetTagDataByCase",
		"GET",
		fmt.Sprintf("%s/tag/show/{casename}", _api),
		audit.GetTagData,
	},
	{
		"GetTagDataByCaseEndpoint",
		"GET",
		fmt.Sprintf("%s/tag/show/{casename}/{endpoint}", _api),
		audit.GetTagData,
	},
	{
		"GetTagByCaseEndpointAudit",
		"GET",
		fmt.Sprintf("%s/tag/show/{casename}/{endpoint}/{audit}", _api),
		audit.GetTagData,
	},
	{
		"GetTagDataByCaseEndpointAuditDocId",
		"GET",
		fmt.Sprintf("%s/tag/show/{casename}/{endpoint}/{audit}/{doc_id}", _api),
		audit.GetTagData,
	},
}
