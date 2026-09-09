from django.urls import path
from analyticsOnDjango.views import ForecastView

urlpatterns = [
    path('', ForecastView.index, name='index'),
	path('create', ForecastView.get, name='create'),
	path('get/<int:forecastId>/', ForecastView.get, name='get'),
	path('save', ForecastView.save, name='save'),
	path('getAll', ForecastView.getAll, name='getAll'),
	path('delete/<int:forecastId>/', ForecastView.delete, name='delete'),
	path('assignModelVersion/<int:forecastId>/<int:ModelVersionId>/', ForecastView.assignModelVersion, name='assignModelVersion'),
	path('unassignModelVersion/<int:forecastId>/', ForecastView.unassignModelVersion, name='unassignModelVersion'),
	path('assignTimeSeries/<int:forecastId>/<int:TimeSeriesId>/', ForecastView.assignTimeSeries, name='assignTimeSeries'),
	path('unassignTimeSeries/<int:forecastId>/', ForecastView.unassignTimeSeries, name='unassignTimeSeries'),
	path('addDatasets/<int:forecastId>/<DatasetsIds>/', ForecastView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:forecastId>/<DatasetsIds>/', ForecastView.removeDatasets, name='removeDatasets'),
]
