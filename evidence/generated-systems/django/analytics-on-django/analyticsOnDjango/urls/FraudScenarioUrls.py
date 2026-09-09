from django.urls import path
from analyticsOnDjango.views import FraudScenarioView

urlpatterns = [
    path('', FraudScenarioView.index, name='index'),
	path('create', FraudScenarioView.get, name='create'),
	path('get/<int:fraudScenarioId>/', FraudScenarioView.get, name='get'),
	path('save', FraudScenarioView.save, name='save'),
	path('getAll', FraudScenarioView.getAll, name='getAll'),
	path('delete/<int:fraudScenarioId>/', FraudScenarioView.delete, name='delete'),
	path('addModels/<int:fraudScenarioId>/<ModelsIds>/', FraudScenarioView.addModels, name='addModels'),
	path('removeModels/<int:fraudScenarioId>/<ModelsIds>/', FraudScenarioView.removeModels, name='removeModels'),
	path('addDatasets/<int:fraudScenarioId>/<DatasetsIds>/', FraudScenarioView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:fraudScenarioId>/<DatasetsIds>/', FraudScenarioView.removeDatasets, name='removeDatasets'),
	path('addAlerts/<int:fraudScenarioId>/<AlertsIds>/', FraudScenarioView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:fraudScenarioId>/<AlertsIds>/', FraudScenarioView.removeAlerts, name='removeAlerts'),
	path('addSignals/<int:fraudScenarioId>/<SignalsIds>/', FraudScenarioView.addSignals, name='addSignals'),
	path('removeSignals/<int:fraudScenarioId>/<SignalsIds>/', FraudScenarioView.removeSignals, name='removeSignals'),
]
