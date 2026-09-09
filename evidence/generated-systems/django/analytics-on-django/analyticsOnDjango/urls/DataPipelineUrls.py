from django.urls import path
from analyticsOnDjango.views import DataPipelineView

urlpatterns = [
    path('', DataPipelineView.index, name='index'),
	path('create', DataPipelineView.get, name='create'),
	path('get/<int:dataPipelineId>/', DataPipelineView.get, name='get'),
	path('save', DataPipelineView.save, name='save'),
	path('getAll', DataPipelineView.getAll, name='getAll'),
	path('delete/<int:dataPipelineId>/', DataPipelineView.delete, name='delete'),
	path('assignWorkspace/<int:dataPipelineId>/<int:WorkspaceId>/', DataPipelineView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:dataPipelineId>/', DataPipelineView.unassignWorkspace, name='unassignWorkspace'),
	path('assignLineageNode/<int:dataPipelineId>/<int:LineageNodeId>/', DataPipelineView.assignLineageNode, name='assignLineageNode'),
	path('unassignLineageNode/<int:dataPipelineId>/', DataPipelineView.unassignLineageNode, name='unassignLineageNode'),
	path('addTasks/<int:dataPipelineId>/<TasksIds>/', DataPipelineView.addTasks, name='addTasks'),
	path('removeTasks/<int:dataPipelineId>/<TasksIds>/', DataPipelineView.removeTasks, name='removeTasks'),
	path('addSources/<int:dataPipelineId>/<SourcesIds>/', DataPipelineView.addSources, name='addSources'),
	path('removeSources/<int:dataPipelineId>/<SourcesIds>/', DataPipelineView.removeSources, name='removeSources'),
	path('addOutputs/<int:dataPipelineId>/<OutputsIds>/', DataPipelineView.addOutputs, name='addOutputs'),
	path('removeOutputs/<int:dataPipelineId>/<OutputsIds>/', DataPipelineView.removeOutputs, name='removeOutputs'),
]
