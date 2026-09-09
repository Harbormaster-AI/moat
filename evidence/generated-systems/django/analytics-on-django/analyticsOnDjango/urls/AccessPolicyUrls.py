from django.urls import path
from analyticsOnDjango.views import AccessPolicyView

urlpatterns = [
    path('', AccessPolicyView.index, name='index'),
	path('create', AccessPolicyView.get, name='create'),
	path('get/<int:accessPolicyId>/', AccessPolicyView.get, name='get'),
	path('save', AccessPolicyView.save, name='save'),
	path('getAll', AccessPolicyView.getAll, name='getAll'),
	path('delete/<int:accessPolicyId>/', AccessPolicyView.delete, name='delete'),
	path('assignWorkspace/<int:accessPolicyId>/<int:WorkspaceId>/', AccessPolicyView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:accessPolicyId>/', AccessPolicyView.unassignWorkspace, name='unassignWorkspace'),
	path('addDatasets/<int:accessPolicyId>/<DatasetsIds>/', AccessPolicyView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:accessPolicyId>/<DatasetsIds>/', AccessPolicyView.removeDatasets, name='removeDatasets'),
	path('addDashboards/<int:accessPolicyId>/<DashboardsIds>/', AccessPolicyView.addDashboards, name='addDashboards'),
	path('removeDashboards/<int:accessPolicyId>/<DashboardsIds>/', AccessPolicyView.removeDashboards, name='removeDashboards'),
	path('addReports/<int:accessPolicyId>/<ReportsIds>/', AccessPolicyView.addReports, name='addReports'),
	path('removeReports/<int:accessPolicyId>/<ReportsIds>/', AccessPolicyView.removeReports, name='removeReports'),
	path('addModels/<int:accessPolicyId>/<ModelsIds>/', AccessPolicyView.addModels, name='addModels'),
	path('removeModels/<int:accessPolicyId>/<ModelsIds>/', AccessPolicyView.removeModels, name='removeModels'),
	path('addFeatureSets/<int:accessPolicyId>/<FeatureSetsIds>/', AccessPolicyView.addFeatureSets, name='addFeatureSets'),
	path('removeFeatureSets/<int:accessPolicyId>/<FeatureSetsIds>/', AccessPolicyView.removeFeatureSets, name='removeFeatureSets'),
]
