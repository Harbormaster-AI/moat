from django.urls import path
from analyticsOnDjango.views import InferenceEndpointView

urlpatterns = [
    path('', InferenceEndpointView.index, name='index'),
	path('create', InferenceEndpointView.get, name='create'),
	path('get/<int:inferenceEndpointId>/', InferenceEndpointView.get, name='get'),
	path('save', InferenceEndpointView.save, name='save'),
	path('getAll', InferenceEndpointView.getAll, name='getAll'),
	path('delete/<int:inferenceEndpointId>/', InferenceEndpointView.delete, name='delete'),
	path('assignModelVersion/<int:inferenceEndpointId>/<int:ModelVersionId>/', InferenceEndpointView.assignModelVersion, name='assignModelVersion'),
	path('unassignModelVersion/<int:inferenceEndpointId>/', InferenceEndpointView.unassignModelVersion, name='unassignModelVersion'),
	path('assignWorkspace/<int:inferenceEndpointId>/<int:WorkspaceId>/', InferenceEndpointView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:inferenceEndpointId>/', InferenceEndpointView.unassignWorkspace, name='unassignWorkspace'),
	path('addPredictions/<int:inferenceEndpointId>/<PredictionsIds>/', InferenceEndpointView.addPredictions, name='addPredictions'),
	path('removePredictions/<int:inferenceEndpointId>/<PredictionsIds>/', InferenceEndpointView.removePredictions, name='removePredictions'),
]
