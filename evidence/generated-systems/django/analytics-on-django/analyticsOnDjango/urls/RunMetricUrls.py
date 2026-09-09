from django.urls import path
from analyticsOnDjango.views import RunMetricView

urlpatterns = [
    path('', RunMetricView.index, name='index'),
	path('create', RunMetricView.get, name='create'),
	path('get/<int:runMetricId>/', RunMetricView.get, name='get'),
	path('save', RunMetricView.save, name='save'),
	path('getAll', RunMetricView.getAll, name='getAll'),
	path('delete/<int:runMetricId>/', RunMetricView.delete, name='delete'),
	path('assignTrainingRun/<int:runMetricId>/<int:TrainingRunId>/', RunMetricView.assignTrainingRun, name='assignTrainingRun'),
	path('unassignTrainingRun/<int:runMetricId>/', RunMetricView.unassignTrainingRun, name='unassignTrainingRun'),
	path('assignMetric/<int:runMetricId>/<int:MetricId>/', RunMetricView.assignMetric, name='assignMetric'),
	path('unassignMetric/<int:runMetricId>/', RunMetricView.unassignMetric, name='unassignMetric'),
	path('assignDataset/<int:runMetricId>/<int:DatasetId>/', RunMetricView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:runMetricId>/', RunMetricView.unassignDataset, name='unassignDataset'),
]
