from django.urls import path
from analyticsOnDjango.views import QualityRuleView

urlpatterns = [
    path('', QualityRuleView.index, name='index'),
	path('create', QualityRuleView.get, name='create'),
	path('get/<int:qualityRuleId>/', QualityRuleView.get, name='get'),
	path('save', QualityRuleView.save, name='save'),
	path('getAll', QualityRuleView.getAll, name='getAll'),
	path('delete/<int:qualityRuleId>/', QualityRuleView.delete, name='delete'),
	path('assignDataset/<int:qualityRuleId>/<int:DatasetId>/', QualityRuleView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:qualityRuleId>/', QualityRuleView.unassignDataset, name='unassignDataset'),
	path('addChecks/<int:qualityRuleId>/<ChecksIds>/', QualityRuleView.addChecks, name='addChecks'),
	path('removeChecks/<int:qualityRuleId>/<ChecksIds>/', QualityRuleView.removeChecks, name='removeChecks'),
]
