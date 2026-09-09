from django.urls import path
from analyticsOnDjango.views import EvaluationMetricView

urlpatterns = [
    path('', EvaluationMetricView.index, name='index'),
	path('create', EvaluationMetricView.get, name='create'),
	path('get/<int:evaluationMetricId>/', EvaluationMetricView.get, name='get'),
	path('save', EvaluationMetricView.save, name='save'),
	path('getAll', EvaluationMetricView.getAll, name='getAll'),
	path('delete/<int:evaluationMetricId>/', EvaluationMetricView.delete, name='delete'),
	path('assignModelVersion/<int:evaluationMetricId>/<int:ModelVersionId>/', EvaluationMetricView.assignModelVersion, name='assignModelVersion'),
	path('unassignModelVersion/<int:evaluationMetricId>/', EvaluationMetricView.unassignModelVersion, name='unassignModelVersion'),
	path('assignMetric/<int:evaluationMetricId>/<int:MetricId>/', EvaluationMetricView.assignMetric, name='assignMetric'),
	path('unassignMetric/<int:evaluationMetricId>/', EvaluationMetricView.unassignMetric, name='unassignMetric'),
	path('assignDataset/<int:evaluationMetricId>/<int:DatasetId>/', EvaluationMetricView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:evaluationMetricId>/', EvaluationMetricView.unassignDataset, name='unassignDataset'),
]
