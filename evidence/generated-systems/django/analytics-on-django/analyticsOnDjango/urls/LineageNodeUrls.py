from django.urls import path
from analyticsOnDjango.views import LineageNodeView

urlpatterns = [
    path('', LineageNodeView.index, name='index'),
	path('create', LineageNodeView.get, name='create'),
	path('get/<int:lineageNodeId>/', LineageNodeView.get, name='get'),
	path('save', LineageNodeView.save, name='save'),
	path('getAll', LineageNodeView.getAll, name='getAll'),
	path('delete/<int:lineageNodeId>/', LineageNodeView.delete, name='delete'),
	path('assignWorkspace/<int:lineageNodeId>/<int:WorkspaceId>/', LineageNodeView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:lineageNodeId>/', LineageNodeView.unassignWorkspace, name='unassignWorkspace'),
	path('addInputs/<int:lineageNodeId>/<InputsIds>/', LineageNodeView.addInputs, name='addInputs'),
	path('removeInputs/<int:lineageNodeId>/<InputsIds>/', LineageNodeView.removeInputs, name='removeInputs'),
	path('addOutputs/<int:lineageNodeId>/<OutputsIds>/', LineageNodeView.addOutputs, name='addOutputs'),
	path('removeOutputs/<int:lineageNodeId>/<OutputsIds>/', LineageNodeView.removeOutputs, name='removeOutputs'),
	path('addDatasets/<int:lineageNodeId>/<DatasetsIds>/', LineageNodeView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:lineageNodeId>/<DatasetsIds>/', LineageNodeView.removeDatasets, name='removeDatasets'),
	path('addModels/<int:lineageNodeId>/<ModelsIds>/', LineageNodeView.addModels, name='addModels'),
	path('removeModels/<int:lineageNodeId>/<ModelsIds>/', LineageNodeView.removeModels, name='removeModels'),
	path('addPipelines/<int:lineageNodeId>/<PipelinesIds>/', LineageNodeView.addPipelines, name='addPipelines'),
	path('removePipelines/<int:lineageNodeId>/<PipelinesIds>/', LineageNodeView.removePipelines, name='removePipelines'),
	path('addDashboards/<int:lineageNodeId>/<DashboardsIds>/', LineageNodeView.addDashboards, name='addDashboards'),
	path('removeDashboards/<int:lineageNodeId>/<DashboardsIds>/', LineageNodeView.removeDashboards, name='removeDashboards'),
	path('addReports/<int:lineageNodeId>/<ReportsIds>/', LineageNodeView.addReports, name='addReports'),
	path('removeReports/<int:lineageNodeId>/<ReportsIds>/', LineageNodeView.removeReports, name='removeReports'),
]
