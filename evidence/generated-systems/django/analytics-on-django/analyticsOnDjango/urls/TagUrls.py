from django.urls import path
from analyticsOnDjango.views import TagView

urlpatterns = [
    path('', TagView.index, name='index'),
	path('create', TagView.get, name='create'),
	path('get/<int:tagId>/', TagView.get, name='get'),
	path('save', TagView.save, name='save'),
	path('getAll', TagView.getAll, name='getAll'),
	path('delete/<int:tagId>/', TagView.delete, name='delete'),
	path('addDatasets/<int:tagId>/<DatasetsIds>/', TagView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:tagId>/<DatasetsIds>/', TagView.removeDatasets, name='removeDatasets'),
	path('addModels/<int:tagId>/<ModelsIds>/', TagView.addModels, name='addModels'),
	path('removeModels/<int:tagId>/<ModelsIds>/', TagView.removeModels, name='removeModels'),
	path('addModelVersions/<int:tagId>/<ModelVersionsIds>/', TagView.addModelVersions, name='addModelVersions'),
	path('removeModelVersions/<int:tagId>/<ModelVersionsIds>/', TagView.removeModelVersions, name='removeModelVersions'),
	path('addDashboards/<int:tagId>/<DashboardsIds>/', TagView.addDashboards, name='addDashboards'),
	path('removeDashboards/<int:tagId>/<DashboardsIds>/', TagView.removeDashboards, name='removeDashboards'),
	path('addReports/<int:tagId>/<ReportsIds>/', TagView.addReports, name='addReports'),
	path('removeReports/<int:tagId>/<ReportsIds>/', TagView.removeReports, name='removeReports'),
	path('addFeatureSets/<int:tagId>/<FeatureSetsIds>/', TagView.addFeatureSets, name='addFeatureSets'),
	path('removeFeatureSets/<int:tagId>/<FeatureSetsIds>/', TagView.removeFeatureSets, name='removeFeatureSets'),
	path('addMetrics/<int:tagId>/<MetricsIds>/', TagView.addMetrics, name='addMetrics'),
	path('removeMetrics/<int:tagId>/<MetricsIds>/', TagView.removeMetrics, name='removeMetrics'),
]
