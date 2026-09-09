from django.urls import path
from analyticsOnDjango.views import QualityCheckView

urlpatterns = [
    path('', QualityCheckView.index, name='index'),
	path('create', QualityCheckView.get, name='create'),
	path('get/<int:qualityCheckId>/', QualityCheckView.get, name='get'),
	path('save', QualityCheckView.save, name='save'),
	path('getAll', QualityCheckView.getAll, name='getAll'),
	path('delete/<int:qualityCheckId>/', QualityCheckView.delete, name='delete'),
	path('assignRule/<int:qualityCheckId>/<int:RuleId>/', QualityCheckView.assignRule, name='assignRule'),
	path('unassignRule/<int:qualityCheckId>/', QualityCheckView.unassignRule, name='unassignRule'),
	path('assignDataset/<int:qualityCheckId>/<int:DatasetId>/', QualityCheckView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:qualityCheckId>/', QualityCheckView.unassignDataset, name='unassignDataset'),
]
