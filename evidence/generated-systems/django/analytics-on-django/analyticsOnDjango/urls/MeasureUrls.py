from django.urls import path
from analyticsOnDjango.views import MeasureView

urlpatterns = [
    path('', MeasureView.index, name='index'),
	path('create', MeasureView.get, name='create'),
	path('get/<int:measureId>/', MeasureView.get, name='get'),
	path('save', MeasureView.save, name='save'),
	path('getAll', MeasureView.getAll, name='getAll'),
	path('delete/<int:measureId>/', MeasureView.delete, name='delete'),
	path('assignSemanticModel/<int:measureId>/<int:SemanticModelId>/', MeasureView.assignSemanticModel, name='assignSemanticModel'),
	path('unassignSemanticModel/<int:measureId>/', MeasureView.unassignSemanticModel, name='unassignSemanticModel'),
	path('addDatasets/<int:measureId>/<DatasetsIds>/', MeasureView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:measureId>/<DatasetsIds>/', MeasureView.removeDatasets, name='removeDatasets'),
	path('addGlossaryTerms/<int:measureId>/<GlossaryTermsIds>/', MeasureView.addGlossaryTerms, name='addGlossaryTerms'),
	path('removeGlossaryTerms/<int:measureId>/<GlossaryTermsIds>/', MeasureView.removeGlossaryTerms, name='removeGlossaryTerms'),
]
