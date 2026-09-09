from django.urls import path
from analyticsOnDjango.views import SemanticModelView

urlpatterns = [
    path('', SemanticModelView.index, name='index'),
	path('create', SemanticModelView.get, name='create'),
	path('get/<int:semanticModelId>/', SemanticModelView.get, name='get'),
	path('save', SemanticModelView.save, name='save'),
	path('getAll', SemanticModelView.getAll, name='getAll'),
	path('delete/<int:semanticModelId>/', SemanticModelView.delete, name='delete'),
	path('addDatasets/<int:semanticModelId>/<DatasetsIds>/', SemanticModelView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:semanticModelId>/<DatasetsIds>/', SemanticModelView.removeDatasets, name='removeDatasets'),
	path('addMetrics/<int:semanticModelId>/<MetricsIds>/', SemanticModelView.addMetrics, name='addMetrics'),
	path('removeMetrics/<int:semanticModelId>/<MetricsIds>/', SemanticModelView.removeMetrics, name='removeMetrics'),
	path('addDimensions/<int:semanticModelId>/<DimensionsIds>/', SemanticModelView.addDimensions, name='addDimensions'),
	path('removeDimensions/<int:semanticModelId>/<DimensionsIds>/', SemanticModelView.removeDimensions, name='removeDimensions'),
	path('addMeasures/<int:semanticModelId>/<MeasuresIds>/', SemanticModelView.addMeasures, name='addMeasures'),
	path('removeMeasures/<int:semanticModelId>/<MeasuresIds>/', SemanticModelView.removeMeasures, name='removeMeasures'),
	path('addGlossaryTerms/<int:semanticModelId>/<GlossaryTermsIds>/', SemanticModelView.addGlossaryTerms, name='addGlossaryTerms'),
	path('removeGlossaryTerms/<int:semanticModelId>/<GlossaryTermsIds>/', SemanticModelView.removeGlossaryTerms, name='removeGlossaryTerms'),
]
