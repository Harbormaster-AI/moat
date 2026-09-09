from django.urls import path
from manufacturingOnDjango.views import ForecastLineView

urlpatterns = [
    path('', ForecastLineView.index, name='index'),
	path('create', ForecastLineView.get, name='create'),
	path('get/<int:forecastLineId>/', ForecastLineView.get, name='get'),
	path('save', ForecastLineView.save, name='save'),
	path('getAll', ForecastLineView.getAll, name='getAll'),
	path('delete/<int:forecastLineId>/', ForecastLineView.delete, name='delete'),
	path('assignForecast/<int:forecastLineId>/<int:ForecastId>/', ForecastLineView.assignForecast, name='assignForecast'),
	path('unassignForecast/<int:forecastLineId>/', ForecastLineView.unassignForecast, name='unassignForecast'),
	path('assignItem/<int:forecastLineId>/<int:ItemId>/', ForecastLineView.assignItem, name='assignItem'),
	path('unassignItem/<int:forecastLineId>/', ForecastLineView.unassignItem, name='unassignItem'),
]
