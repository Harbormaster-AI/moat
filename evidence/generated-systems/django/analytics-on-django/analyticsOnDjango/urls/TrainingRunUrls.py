from django.urls import path
from analyticsOnDjango.views import TrainingRunView

urlpatterns = [
    path('', TrainingRunView.index, name='index'),
	path('create', TrainingRunView.get, name='create'),
	path('get/<int:trainingRunId>/', TrainingRunView.get, name='get'),
	path('save', TrainingRunView.save, name='save'),
	path('getAll', TrainingRunView.getAll, name='getAll'),
	path('delete/<int:trainingRunId>/', TrainingRunView.delete, name='delete'),
	path('assignExperiment/<int:trainingRunId>/<int:ExperimentId>/', TrainingRunView.assignExperiment, name='assignExperiment'),
	path('unassignExperiment/<int:trainingRunId>/', TrainingRunView.unassignExperiment, name='unassignExperiment'),
	path('assignModelVersion/<int:trainingRunId>/<int:ModelVersionId>/', TrainingRunView.assignModelVersion, name='assignModelVersion'),
	path('unassignModelVersion/<int:trainingRunId>/', TrainingRunView.unassignModelVersion, name='unassignModelVersion'),
	path('addInputDatasets/<int:trainingRunId>/<InputDatasetsIds>/', TrainingRunView.addInputDatasets, name='addInputDatasets'),
	path('removeInputDatasets/<int:trainingRunId>/<InputDatasetsIds>/', TrainingRunView.removeInputDatasets, name='removeInputDatasets'),
	path('addFeatures/<int:trainingRunId>/<FeaturesIds>/', TrainingRunView.addFeatures, name='addFeatures'),
	path('removeFeatures/<int:trainingRunId>/<FeaturesIds>/', TrainingRunView.removeFeatures, name='removeFeatures'),
	path('addRunMetrics/<int:trainingRunId>/<RunMetricsIds>/', TrainingRunView.addRunMetrics, name='addRunMetrics'),
	path('removeRunMetrics/<int:trainingRunId>/<RunMetricsIds>/', TrainingRunView.removeRunMetrics, name='removeRunMetrics'),
	path('addRunParameters/<int:trainingRunId>/<RunParametersIds>/', TrainingRunView.addRunParameters, name='addRunParameters'),
	path('removeRunParameters/<int:trainingRunId>/<RunParametersIds>/', TrainingRunView.removeRunParameters, name='removeRunParameters'),
]
