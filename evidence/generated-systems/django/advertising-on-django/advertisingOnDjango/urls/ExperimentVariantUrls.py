from django.urls import path
from advertisingOnDjango.views import ExperimentVariantView

urlpatterns = [
    path('', ExperimentVariantView.index, name='index'),
	path('create', ExperimentVariantView.get, name='create'),
	path('get/<int:experimentVariantId>/', ExperimentVariantView.get, name='get'),
	path('save', ExperimentVariantView.save, name='save'),
	path('getAll', ExperimentVariantView.getAll, name='getAll'),
	path('delete/<int:experimentVariantId>/', ExperimentVariantView.delete, name='delete'),
	path('assignExperiment/<int:experimentVariantId>/<int:ExperimentId>/', ExperimentVariantView.assignExperiment, name='assignExperiment'),
	path('unassignExperiment/<int:experimentVariantId>/', ExperimentVariantView.unassignExperiment, name='unassignExperiment'),
	path('assignCreativeVariation/<int:experimentVariantId>/<int:CreativeVariationId>/', ExperimentVariantView.assignCreativeVariation, name='assignCreativeVariation'),
	path('unassignCreativeVariation/<int:experimentVariantId>/', ExperimentVariantView.unassignCreativeVariation, name='unassignCreativeVariation'),
	path('assignLineItem/<int:experimentVariantId>/<int:LineItemId>/', ExperimentVariantView.assignLineItem, name='assignLineItem'),
	path('unassignLineItem/<int:experimentVariantId>/', ExperimentVariantView.unassignLineItem, name='unassignLineItem'),
]
