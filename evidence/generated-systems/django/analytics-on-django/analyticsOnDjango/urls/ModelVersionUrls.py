from django.urls import path
from analyticsOnDjango.views import ModelVersionView

urlpatterns = [
    path('', ModelVersionView.index, name='index'),
	path('create', ModelVersionView.get, name='create'),
	path('get/<int:modelVersionId>/', ModelVersionView.get, name='get'),
	path('save', ModelVersionView.save, name='save'),
	path('getAll', ModelVersionView.getAll, name='getAll'),
	path('delete/<int:modelVersionId>/', ModelVersionView.delete, name='delete'),
	path('assignModel/<int:modelVersionId>/<int:ModelId>/', ModelVersionView.assignModel, name='assignModel'),
	path('unassignModel/<int:modelVersionId>/', ModelVersionView.unassignModel, name='unassignModel'),
	path('assignTrainingRun/<int:modelVersionId>/<int:TrainingRunId>/', ModelVersionView.assignTrainingRun, name='assignTrainingRun'),
	path('unassignTrainingRun/<int:modelVersionId>/', ModelVersionView.unassignTrainingRun, name='unassignTrainingRun'),
	path('addEvaluationMetrics/<int:modelVersionId>/<EvaluationMetricsIds>/', ModelVersionView.addEvaluationMetrics, name='addEvaluationMetrics'),
	path('removeEvaluationMetrics/<int:modelVersionId>/<EvaluationMetricsIds>/', ModelVersionView.removeEvaluationMetrics, name='removeEvaluationMetrics'),
	path('addDeployments/<int:modelVersionId>/<DeploymentsIds>/', ModelVersionView.addDeployments, name='addDeployments'),
	path('removeDeployments/<int:modelVersionId>/<DeploymentsIds>/', ModelVersionView.removeDeployments, name='removeDeployments'),
	path('addFeatureSets/<int:modelVersionId>/<FeatureSetsIds>/', ModelVersionView.addFeatureSets, name='addFeatureSets'),
	path('removeFeatureSets/<int:modelVersionId>/<FeatureSetsIds>/', ModelVersionView.removeFeatureSets, name='removeFeatureSets'),
	path('addDatasets/<int:modelVersionId>/<DatasetsIds>/', ModelVersionView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:modelVersionId>/<DatasetsIds>/', ModelVersionView.removeDatasets, name='removeDatasets'),
]
