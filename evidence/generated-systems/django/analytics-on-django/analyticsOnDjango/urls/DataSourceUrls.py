from django.urls import path
from analyticsOnDjango.views import DataSourceView

urlpatterns = [
    path('', DataSourceView.index, name='index'),
	path('create', DataSourceView.get, name='create'),
	path('get/<int:dataSourceId>/', DataSourceView.get, name='get'),
	path('save', DataSourceView.save, name='save'),
	path('getAll', DataSourceView.getAll, name='getAll'),
	path('delete/<int:dataSourceId>/', DataSourceView.delete, name='delete'),
	path('assignWorkspace/<int:dataSourceId>/<int:WorkspaceId>/', DataSourceView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:dataSourceId>/', DataSourceView.unassignWorkspace, name='unassignWorkspace'),
	path('addProducedDatasets/<int:dataSourceId>/<ProducedDatasetsIds>/', DataSourceView.addProducedDatasets, name='addProducedDatasets'),
	path('removeProducedDatasets/<int:dataSourceId>/<ProducedDatasetsIds>/', DataSourceView.removeProducedDatasets, name='removeProducedDatasets'),
	path('addPipelines/<int:dataSourceId>/<PipelinesIds>/', DataSourceView.addPipelines, name='addPipelines'),
	path('removePipelines/<int:dataSourceId>/<PipelinesIds>/', DataSourceView.removePipelines, name='removePipelines'),
]
