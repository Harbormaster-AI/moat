from django.urls import path
from analyticsOnDjango.views import BIQueryView

urlpatterns = [
    path('', BIQueryView.index, name='index'),
	path('create', BIQueryView.get, name='create'),
	path('get/<int:bIQueryId>/', BIQueryView.get, name='get'),
	path('save', BIQueryView.save, name='save'),
	path('getAll', BIQueryView.getAll, name='getAll'),
	path('delete/<int:bIQueryId>/', BIQueryView.delete, name='delete'),
	path('assignWorkspace/<int:bIQueryId>/<int:WorkspaceId>/', BIQueryView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:bIQueryId>/', BIQueryView.unassignWorkspace, name='unassignWorkspace'),
	path('addDatasets/<int:bIQueryId>/<DatasetsIds>/', BIQueryView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:bIQueryId>/<DatasetsIds>/', BIQueryView.removeDatasets, name='removeDatasets'),
	path('addReports/<int:bIQueryId>/<ReportsIds>/', BIQueryView.addReports, name='addReports'),
	path('removeReports/<int:bIQueryId>/<ReportsIds>/', BIQueryView.removeReports, name='removeReports'),
	path('addDashboards/<int:bIQueryId>/<DashboardsIds>/', BIQueryView.addDashboards, name='addDashboards'),
	path('removeDashboards/<int:bIQueryId>/<DashboardsIds>/', BIQueryView.removeDashboards, name='removeDashboards'),
	path('addNotebooks/<int:bIQueryId>/<NotebooksIds>/', BIQueryView.addNotebooks, name='addNotebooks'),
	path('removeNotebooks/<int:bIQueryId>/<NotebooksIds>/', BIQueryView.removeNotebooks, name='removeNotebooks'),
]
