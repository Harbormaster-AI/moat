from django.urls import path
from advertisingOnDjango.views import CreativeVariationView

urlpatterns = [
    path('', CreativeVariationView.index, name='index'),
	path('create', CreativeVariationView.get, name='create'),
	path('get/<int:creativeVariationId>/', CreativeVariationView.get, name='get'),
	path('save', CreativeVariationView.save, name='save'),
	path('getAll', CreativeVariationView.getAll, name='getAll'),
	path('delete/<int:creativeVariationId>/', CreativeVariationView.delete, name='delete'),
	path('assignCreativeAsset/<int:creativeVariationId>/<int:CreativeAssetId>/', CreativeVariationView.assignCreativeAsset, name='assignCreativeAsset'),
	path('unassignCreativeAsset/<int:creativeVariationId>/', CreativeVariationView.unassignCreativeAsset, name='unassignCreativeAsset'),
]
