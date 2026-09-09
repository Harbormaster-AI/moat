from django.urls import path
from advertisingOnDjango.views import CreativeAssetView

urlpatterns = [
    path('', CreativeAssetView.index, name='index'),
	path('create', CreativeAssetView.get, name='create'),
	path('get/<int:creativeAssetId>/', CreativeAssetView.get, name='get'),
	path('save', CreativeAssetView.save, name='save'),
	path('getAll', CreativeAssetView.getAll, name='getAll'),
	path('delete/<int:creativeAssetId>/', CreativeAssetView.delete, name='delete'),
	path('addFiles/<int:creativeAssetId>/<FilesIds>/', CreativeAssetView.addFiles, name='addFiles'),
	path('removeFiles/<int:creativeAssetId>/<FilesIds>/', CreativeAssetView.removeFiles, name='removeFiles'),
	path('addApprovals/<int:creativeAssetId>/<ApprovalsIds>/', CreativeAssetView.addApprovals, name='addApprovals'),
	path('removeApprovals/<int:creativeAssetId>/<ApprovalsIds>/', CreativeAssetView.removeApprovals, name='removeApprovals'),
	path('addVariations/<int:creativeAssetId>/<VariationsIds>/', CreativeAssetView.addVariations, name='addVariations'),
	path('removeVariations/<int:creativeAssetId>/<VariationsIds>/', CreativeAssetView.removeVariations, name='removeVariations'),
	path('addLineItems/<int:creativeAssetId>/<LineItemsIds>/', CreativeAssetView.addLineItems, name='addLineItems'),
	path('removeLineItems/<int:creativeAssetId>/<LineItemsIds>/', CreativeAssetView.removeLineItems, name='removeLineItems'),
]
