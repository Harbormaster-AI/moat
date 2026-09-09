from django.urls import path
from analyticsOnDjango.views import PredictionView

urlpatterns = [
    path('', PredictionView.index, name='index'),
	path('create', PredictionView.get, name='create'),
	path('get/<int:predictionId>/', PredictionView.get, name='get'),
	path('save', PredictionView.save, name='save'),
	path('getAll', PredictionView.getAll, name='getAll'),
	path('delete/<int:predictionId>/', PredictionView.delete, name='delete'),
	path('assignEndpoint/<int:predictionId>/<int:EndpointId>/', PredictionView.assignEndpoint, name='assignEndpoint'),
	path('unassignEndpoint/<int:predictionId>/', PredictionView.unassignEndpoint, name='unassignEndpoint'),
	path('assignModelVersion/<int:predictionId>/<int:ModelVersionId>/', PredictionView.assignModelVersion, name='assignModelVersion'),
	path('unassignModelVersion/<int:predictionId>/', PredictionView.unassignModelVersion, name='unassignModelVersion'),
	path('assignDataset/<int:predictionId>/<int:DatasetId>/', PredictionView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:predictionId>/', PredictionView.unassignDataset, name='unassignDataset'),
]
