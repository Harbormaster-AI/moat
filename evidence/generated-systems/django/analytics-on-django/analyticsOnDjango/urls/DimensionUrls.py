from django.urls import path
from analyticsOnDjango.views import DimensionView

urlpatterns = [
    path('', DimensionView.index, name='index'),
	path('create', DimensionView.get, name='create'),
	path('get/<int:dimensionId>/', DimensionView.get, name='get'),
	path('save', DimensionView.save, name='save'),
	path('getAll', DimensionView.getAll, name='getAll'),
	path('delete/<int:dimensionId>/', DimensionView.delete, name='delete'),
	path('assignSemanticModel/<int:dimensionId>/<int:SemanticModelId>/', DimensionView.assignSemanticModel, name='assignSemanticModel'),
	path('unassignSemanticModel/<int:dimensionId>/', DimensionView.unassignSemanticModel, name='unassignSemanticModel'),
	path('addDatasets/<int:dimensionId>/<DatasetsIds>/', DimensionView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:dimensionId>/<DatasetsIds>/', DimensionView.removeDatasets, name='removeDatasets'),
	path('addGlossaryTerms/<int:dimensionId>/<GlossaryTermsIds>/', DimensionView.addGlossaryTerms, name='addGlossaryTerms'),
	path('removeGlossaryTerms/<int:dimensionId>/<GlossaryTermsIds>/', DimensionView.removeGlossaryTerms, name='removeGlossaryTerms'),
]
