from django.urls import path
from analyticsOnDjango.views import DashboardView

urlpatterns = [
    path('', DashboardView.index, name='index'),
	path('create', DashboardView.get, name='create'),
	path('get/<int:dashboardId>/', DashboardView.get, name='get'),
	path('save', DashboardView.save, name='save'),
	path('getAll', DashboardView.getAll, name='getAll'),
	path('delete/<int:dashboardId>/', DashboardView.delete, name='delete'),
	path('assignWorkspace/<int:dashboardId>/<int:WorkspaceId>/', DashboardView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:dashboardId>/', DashboardView.unassignWorkspace, name='unassignWorkspace'),
	path('addVisualizations/<int:dashboardId>/<VisualizationsIds>/', DashboardView.addVisualizations, name='addVisualizations'),
	path('removeVisualizations/<int:dashboardId>/<VisualizationsIds>/', DashboardView.removeVisualizations, name='removeVisualizations'),
	path('addReports/<int:dashboardId>/<ReportsIds>/', DashboardView.addReports, name='addReports'),
	path('removeReports/<int:dashboardId>/<ReportsIds>/', DashboardView.removeReports, name='removeReports'),
	path('addDatasets/<int:dashboardId>/<DatasetsIds>/', DashboardView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:dashboardId>/<DatasetsIds>/', DashboardView.removeDatasets, name='removeDatasets'),
	path('addAlerts/<int:dashboardId>/<AlertsIds>/', DashboardView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:dashboardId>/<AlertsIds>/', DashboardView.removeAlerts, name='removeAlerts'),
	path('addQueries/<int:dashboardId>/<QueriesIds>/', DashboardView.addQueries, name='addQueries'),
	path('removeQueries/<int:dashboardId>/<QueriesIds>/', DashboardView.removeQueries, name='removeQueries'),
	path('addTags/<int:dashboardId>/<TagsIds>/', DashboardView.addTags, name='addTags'),
	path('removeTags/<int:dashboardId>/<TagsIds>/', DashboardView.removeTags, name='removeTags'),
]
