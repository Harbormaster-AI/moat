from django.urls import path
from analyticsOnDjango.views import RunParameterView

urlpatterns = [
    path('', RunParameterView.index, name='index'),
	path('create', RunParameterView.get, name='create'),
	path('get/<int:runParameterId>/', RunParameterView.get, name='get'),
	path('save', RunParameterView.save, name='save'),
	path('getAll', RunParameterView.getAll, name='getAll'),
	path('delete/<int:runParameterId>/', RunParameterView.delete, name='delete'),
	path('assignTrainingRun/<int:runParameterId>/<int:TrainingRunId>/', RunParameterView.assignTrainingRun, name='assignTrainingRun'),
	path('unassignTrainingRun/<int:runParameterId>/', RunParameterView.unassignTrainingRun, name='unassignTrainingRun'),
]
