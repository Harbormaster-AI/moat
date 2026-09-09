from django.urls import path
from analyticsOnDjango.views import RecommendationScenarioView

urlpatterns = [
    path('', RecommendationScenarioView.index, name='index'),
	path('create', RecommendationScenarioView.get, name='create'),
	path('get/<int:recommendationScenarioId>/', RecommendationScenarioView.get, name='get'),
	path('save', RecommendationScenarioView.save, name='save'),
	path('getAll', RecommendationScenarioView.getAll, name='getAll'),
	path('delete/<int:recommendationScenarioId>/', RecommendationScenarioView.delete, name='delete'),
	path('addModels/<int:recommendationScenarioId>/<ModelsIds>/', RecommendationScenarioView.addModels, name='addModels'),
	path('removeModels/<int:recommendationScenarioId>/<ModelsIds>/', RecommendationScenarioView.removeModels, name='removeModels'),
	path('addDatasets/<int:recommendationScenarioId>/<DatasetsIds>/', RecommendationScenarioView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:recommendationScenarioId>/<DatasetsIds>/', RecommendationScenarioView.removeDatasets, name='removeDatasets'),
	path('addExperiments/<int:recommendationScenarioId>/<ExperimentsIds>/', RecommendationScenarioView.addExperiments, name='addExperiments'),
	path('removeExperiments/<int:recommendationScenarioId>/<ExperimentsIds>/', RecommendationScenarioView.removeExperiments, name='removeExperiments'),
	path('addAlerts/<int:recommendationScenarioId>/<AlertsIds>/', RecommendationScenarioView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:recommendationScenarioId>/<AlertsIds>/', RecommendationScenarioView.removeAlerts, name='removeAlerts'),
]
