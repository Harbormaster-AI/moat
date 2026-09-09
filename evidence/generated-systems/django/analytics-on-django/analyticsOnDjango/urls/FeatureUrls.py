from django.urls import path
from analyticsOnDjango.views import FeatureView

urlpatterns = [
    path('', FeatureView.index, name='index'),
	path('create', FeatureView.get, name='create'),
	path('get/<int:featureId>/', FeatureView.get, name='get'),
	path('save', FeatureView.save, name='save'),
	path('getAll', FeatureView.getAll, name='getAll'),
	path('delete/<int:featureId>/', FeatureView.delete, name='delete'),
	path('assignFeatureSet/<int:featureId>/<int:FeatureSetId>/', FeatureView.assignFeatureSet, name='assignFeatureSet'),
	path('unassignFeatureSet/<int:featureId>/', FeatureView.unassignFeatureSet, name='unassignFeatureSet'),
	path('addSourceDatasets/<int:featureId>/<SourceDatasetsIds>/', FeatureView.addSourceDatasets, name='addSourceDatasets'),
	path('removeSourceDatasets/<int:featureId>/<SourceDatasetsIds>/', FeatureView.removeSourceDatasets, name='removeSourceDatasets'),
	path('addModels/<int:featureId>/<ModelsIds>/', FeatureView.addModels, name='addModels'),
	path('removeModels/<int:featureId>/<ModelsIds>/', FeatureView.removeModels, name='removeModels'),
	path('addTrainingRuns/<int:featureId>/<TrainingRunsIds>/', FeatureView.addTrainingRuns, name='addTrainingRuns'),
	path('removeTrainingRuns/<int:featureId>/<TrainingRunsIds>/', FeatureView.removeTrainingRuns, name='removeTrainingRuns'),
]
