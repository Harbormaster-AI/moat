from django.urls import path
from manufacturingOnDjango.views import ForecastView

urlpatterns = [
    path('', ForecastView.index, name='index'),
	path('create', ForecastView.get, name='create'),
	path('get/<int:forecastId>/', ForecastView.get, name='get'),
	path('save', ForecastView.save, name='save'),
	path('getAll', ForecastView.getAll, name='getAll'),
	path('delete/<int:forecastId>/', ForecastView.delete, name='delete'),
	path('addLines/<int:forecastId>/<LinesIds>/', ForecastView.addLines, name='addLines'),
	path('removeLines/<int:forecastId>/<LinesIds>/', ForecastView.removeLines, name='removeLines'),
]
