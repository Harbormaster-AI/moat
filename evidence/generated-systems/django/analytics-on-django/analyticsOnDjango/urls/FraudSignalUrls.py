from django.urls import path
from analyticsOnDjango.views import FraudSignalView

urlpatterns = [
    path('', FraudSignalView.index, name='index'),
	path('create', FraudSignalView.get, name='create'),
	path('get/<int:fraudSignalId>/', FraudSignalView.get, name='get'),
	path('save', FraudSignalView.save, name='save'),
	path('getAll', FraudSignalView.getAll, name='getAll'),
	path('delete/<int:fraudSignalId>/', FraudSignalView.delete, name='delete'),
	path('assignScenario/<int:fraudSignalId>/<int:ScenarioId>/', FraudSignalView.assignScenario, name='assignScenario'),
	path('unassignScenario/<int:fraudSignalId>/', FraudSignalView.unassignScenario, name='unassignScenario'),
	path('assignDataset/<int:fraudSignalId>/<int:DatasetId>/', FraudSignalView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:fraudSignalId>/', FraudSignalView.unassignDataset, name='unassignDataset'),
	path('assignModelVersion/<int:fraudSignalId>/<int:ModelVersionId>/', FraudSignalView.assignModelVersion, name='assignModelVersion'),
	path('unassignModelVersion/<int:fraudSignalId>/', FraudSignalView.unassignModelVersion, name='unassignModelVersion'),
]
