from django.urls import path
from analyticsOnDjango.views import TimeSeriesView

urlpatterns = [
    path('', TimeSeriesView.index, name='index'),
	path('create', TimeSeriesView.get, name='create'),
	path('get/<int:timeSeriesId>/', TimeSeriesView.get, name='get'),
	path('save', TimeSeriesView.save, name='save'),
	path('getAll', TimeSeriesView.getAll, name='getAll'),
	path('delete/<int:timeSeriesId>/', TimeSeriesView.delete, name='delete'),
	path('addDatasets/<int:timeSeriesId>/<DatasetsIds>/', TimeSeriesView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:timeSeriesId>/<DatasetsIds>/', TimeSeriesView.removeDatasets, name='removeDatasets'),
	path('addForecasts/<int:timeSeriesId>/<ForecastsIds>/', TimeSeriesView.addForecasts, name='addForecasts'),
	path('removeForecasts/<int:timeSeriesId>/<ForecastsIds>/', TimeSeriesView.removeForecasts, name='removeForecasts'),
	path('addAnomalies/<int:timeSeriesId>/<AnomaliesIds>/', TimeSeriesView.addAnomalies, name='addAnomalies'),
	path('removeAnomalies/<int:timeSeriesId>/<AnomaliesIds>/', TimeSeriesView.removeAnomalies, name='removeAnomalies'),
]
