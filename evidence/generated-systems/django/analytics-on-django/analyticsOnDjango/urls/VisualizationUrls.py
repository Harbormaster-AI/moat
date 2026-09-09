from django.urls import path
from analyticsOnDjango.views import VisualizationView

urlpatterns = [
    path('', VisualizationView.index, name='index'),
	path('create', VisualizationView.get, name='create'),
	path('get/<int:visualizationId>/', VisualizationView.get, name='get'),
	path('save', VisualizationView.save, name='save'),
	path('getAll', VisualizationView.getAll, name='getAll'),
	path('delete/<int:visualizationId>/', VisualizationView.delete, name='delete'),
	path('assignDashboard/<int:visualizationId>/<int:DashboardId>/', VisualizationView.assignDashboard, name='assignDashboard'),
	path('unassignDashboard/<int:visualizationId>/', VisualizationView.unassignDashboard, name='unassignDashboard'),
	path('assignReport/<int:visualizationId>/<int:ReportId>/', VisualizationView.assignReport, name='assignReport'),
	path('unassignReport/<int:visualizationId>/', VisualizationView.unassignReport, name='unassignReport'),
	path('addMetrics/<int:visualizationId>/<MetricsIds>/', VisualizationView.addMetrics, name='addMetrics'),
	path('removeMetrics/<int:visualizationId>/<MetricsIds>/', VisualizationView.removeMetrics, name='removeMetrics'),
	path('addDimensions/<int:visualizationId>/<DimensionsIds>/', VisualizationView.addDimensions, name='addDimensions'),
	path('removeDimensions/<int:visualizationId>/<DimensionsIds>/', VisualizationView.removeDimensions, name='removeDimensions'),
	path('addDatasets/<int:visualizationId>/<DatasetsIds>/', VisualizationView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:visualizationId>/<DatasetsIds>/', VisualizationView.removeDatasets, name='removeDatasets'),
]
