from django.urls import path
from analyticsOnDjango.views import BusinessGlossaryTermView

urlpatterns = [
    path('', BusinessGlossaryTermView.index, name='index'),
	path('create', BusinessGlossaryTermView.get, name='create'),
	path('get/<int:businessGlossaryTermId>/', BusinessGlossaryTermView.get, name='get'),
	path('save', BusinessGlossaryTermView.save, name='save'),
	path('getAll', BusinessGlossaryTermView.getAll, name='getAll'),
	path('delete/<int:businessGlossaryTermId>/', BusinessGlossaryTermView.delete, name='delete'),
	path('addRelatedTerms/<int:businessGlossaryTermId>/<RelatedTermsIds>/', BusinessGlossaryTermView.addRelatedTerms, name='addRelatedTerms'),
	path('removeRelatedTerms/<int:businessGlossaryTermId>/<RelatedTermsIds>/', BusinessGlossaryTermView.removeRelatedTerms, name='removeRelatedTerms'),
	path('addMetrics/<int:businessGlossaryTermId>/<MetricsIds>/', BusinessGlossaryTermView.addMetrics, name='addMetrics'),
	path('removeMetrics/<int:businessGlossaryTermId>/<MetricsIds>/', BusinessGlossaryTermView.removeMetrics, name='removeMetrics'),
	path('addDatasets/<int:businessGlossaryTermId>/<DatasetsIds>/', BusinessGlossaryTermView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:businessGlossaryTermId>/<DatasetsIds>/', BusinessGlossaryTermView.removeDatasets, name='removeDatasets'),
	path('addDimensions/<int:businessGlossaryTermId>/<DimensionsIds>/', BusinessGlossaryTermView.addDimensions, name='addDimensions'),
	path('removeDimensions/<int:businessGlossaryTermId>/<DimensionsIds>/', BusinessGlossaryTermView.removeDimensions, name='removeDimensions'),
	path('addMeasures/<int:businessGlossaryTermId>/<MeasuresIds>/', BusinessGlossaryTermView.addMeasures, name='addMeasures'),
	path('removeMeasures/<int:businessGlossaryTermId>/<MeasuresIds>/', BusinessGlossaryTermView.removeMeasures, name='removeMeasures'),
]
