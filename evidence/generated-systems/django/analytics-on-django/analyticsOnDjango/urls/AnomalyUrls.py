from django.urls import path
from analyticsOnDjango.views import AnomalyView

urlpatterns = [
    path('', AnomalyView.index, name='index'),
	path('create', AnomalyView.get, name='create'),
	path('get/<int:anomalyId>/', AnomalyView.get, name='get'),
	path('save', AnomalyView.save, name='save'),
	path('getAll', AnomalyView.getAll, name='getAll'),
	path('delete/<int:anomalyId>/', AnomalyView.delete, name='delete'),
	path('assignTimeSeries/<int:anomalyId>/<int:TimeSeriesId>/', AnomalyView.assignTimeSeries, name='assignTimeSeries'),
	path('unassignTimeSeries/<int:anomalyId>/', AnomalyView.unassignTimeSeries, name='unassignTimeSeries'),
	path('assignAlert/<int:anomalyId>/<int:AlertId>/', AnomalyView.assignAlert, name='assignAlert'),
	path('unassignAlert/<int:anomalyId>/', AnomalyView.unassignAlert, name='unassignAlert'),
	path('assignDataset/<int:anomalyId>/<int:DatasetId>/', AnomalyView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:anomalyId>/', AnomalyView.unassignDataset, name='unassignDataset'),
]
