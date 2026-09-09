from django.urls import path
from analyticsOnDjango.views import DataTaskView

urlpatterns = [
    path('', DataTaskView.index, name='index'),
	path('create', DataTaskView.get, name='create'),
	path('get/<int:dataTaskId>/', DataTaskView.get, name='get'),
	path('save', DataTaskView.save, name='save'),
	path('getAll', DataTaskView.getAll, name='getAll'),
	path('delete/<int:dataTaskId>/', DataTaskView.delete, name='delete'),
	path('assignPipeline/<int:dataTaskId>/<int:PipelineId>/', DataTaskView.assignPipeline, name='assignPipeline'),
	path('unassignPipeline/<int:dataTaskId>/', DataTaskView.unassignPipeline, name='unassignPipeline'),
	path('addInputDatasets/<int:dataTaskId>/<InputDatasetsIds>/', DataTaskView.addInputDatasets, name='addInputDatasets'),
	path('removeInputDatasets/<int:dataTaskId>/<InputDatasetsIds>/', DataTaskView.removeInputDatasets, name='removeInputDatasets'),
	path('addOutputDatasets/<int:dataTaskId>/<OutputDatasetsIds>/', DataTaskView.addOutputDatasets, name='addOutputDatasets'),
	path('removeOutputDatasets/<int:dataTaskId>/<OutputDatasetsIds>/', DataTaskView.removeOutputDatasets, name='removeOutputDatasets'),
]
