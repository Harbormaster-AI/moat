from django.urls import path
from analyticsOnDjango.views import NotebookView

urlpatterns = [
    path('', NotebookView.index, name='index'),
	path('create', NotebookView.get, name='create'),
	path('get/<int:notebookId>/', NotebookView.get, name='get'),
	path('save', NotebookView.save, name='save'),
	path('getAll', NotebookView.getAll, name='getAll'),
	path('delete/<int:notebookId>/', NotebookView.delete, name='delete'),
	path('assignWorkspace/<int:notebookId>/<int:WorkspaceId>/', NotebookView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:notebookId>/', NotebookView.unassignWorkspace, name='unassignWorkspace'),
	path('addDatasets/<int:notebookId>/<DatasetsIds>/', NotebookView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:notebookId>/<DatasetsIds>/', NotebookView.removeDatasets, name='removeDatasets'),
	path('addExperiments/<int:notebookId>/<ExperimentsIds>/', NotebookView.addExperiments, name='addExperiments'),
	path('removeExperiments/<int:notebookId>/<ExperimentsIds>/', NotebookView.removeExperiments, name='removeExperiments'),
	path('addQueries/<int:notebookId>/<QueriesIds>/', NotebookView.addQueries, name='addQueries'),
	path('removeQueries/<int:notebookId>/<QueriesIds>/', NotebookView.removeQueries, name='removeQueries'),
]
