from django.urls import path
from analyticsOnDjango.views import ExperimentView

urlpatterns = [
    path('', ExperimentView.index, name='index'),
	path('create', ExperimentView.get, name='create'),
	path('get/<int:experimentId>/', ExperimentView.get, name='get'),
	path('save', ExperimentView.save, name='save'),
	path('getAll', ExperimentView.getAll, name='getAll'),
	path('delete/<int:experimentId>/', ExperimentView.delete, name='delete'),
	path('assignWorkspace/<int:experimentId>/<int:WorkspaceId>/', ExperimentView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:experimentId>/', ExperimentView.unassignWorkspace, name='unassignWorkspace'),
	path('addTrainingRuns/<int:experimentId>/<TrainingRunsIds>/', ExperimentView.addTrainingRuns, name='addTrainingRuns'),
	path('removeTrainingRuns/<int:experimentId>/<TrainingRunsIds>/', ExperimentView.removeTrainingRuns, name='removeTrainingRuns'),
	path('addModels/<int:experimentId>/<ModelsIds>/', ExperimentView.addModels, name='addModels'),
	path('removeModels/<int:experimentId>/<ModelsIds>/', ExperimentView.removeModels, name='removeModels'),
	path('addNotebooks/<int:experimentId>/<NotebooksIds>/', ExperimentView.addNotebooks, name='addNotebooks'),
	path('removeNotebooks/<int:experimentId>/<NotebooksIds>/', ExperimentView.removeNotebooks, name='removeNotebooks'),
]
