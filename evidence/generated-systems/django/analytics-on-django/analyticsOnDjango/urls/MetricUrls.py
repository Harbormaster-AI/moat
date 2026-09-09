from django.urls import path
from analyticsOnDjango.views import MetricView

urlpatterns = [
    path('', MetricView.index, name='index'),
	path('create', MetricView.get, name='create'),
	path('get/<int:metricId>/', MetricView.get, name='get'),
	path('save', MetricView.save, name='save'),
	path('getAll', MetricView.getAll, name='getAll'),
	path('delete/<int:metricId>/', MetricView.delete, name='delete'),
	path('assignSemanticModel/<int:metricId>/<int:SemanticModelId>/', MetricView.assignSemanticModel, name='assignSemanticModel'),
	path('unassignSemanticModel/<int:metricId>/', MetricView.unassignSemanticModel, name='unassignSemanticModel'),
	path('addDatasets/<int:metricId>/<DatasetsIds>/', MetricView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:metricId>/<DatasetsIds>/', MetricView.removeDatasets, name='removeDatasets'),
	path('addGlossaryTerms/<int:metricId>/<GlossaryTermsIds>/', MetricView.addGlossaryTerms, name='addGlossaryTerms'),
	path('removeGlossaryTerms/<int:metricId>/<GlossaryTermsIds>/', MetricView.removeGlossaryTerms, name='removeGlossaryTerms'),
	path('addAlerts/<int:metricId>/<AlertsIds>/', MetricView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:metricId>/<AlertsIds>/', MetricView.removeAlerts, name='removeAlerts'),
	path('addVisualizations/<int:metricId>/<VisualizationsIds>/', MetricView.addVisualizations, name='addVisualizations'),
	path('removeVisualizations/<int:metricId>/<VisualizationsIds>/', MetricView.removeVisualizations, name='removeVisualizations'),
]
