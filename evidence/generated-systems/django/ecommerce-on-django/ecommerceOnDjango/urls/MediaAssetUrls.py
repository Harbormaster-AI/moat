from django.urls import path
from ecommerceOnDjango.views import MediaAssetView

urlpatterns = [
    path('', MediaAssetView.index, name='index'),
	path('create', MediaAssetView.get, name='create'),
	path('get/<int:mediaAssetId>/', MediaAssetView.get, name='get'),
	path('save', MediaAssetView.save, name='save'),
	path('getAll', MediaAssetView.getAll, name='getAll'),
	path('delete/<int:mediaAssetId>/', MediaAssetView.delete, name='delete'),
	path('assignProduct/<int:mediaAssetId>/<int:ProductId>/', MediaAssetView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:mediaAssetId>/', MediaAssetView.unassignProduct, name='unassignProduct'),
	path('assignVariant/<int:mediaAssetId>/<int:VariantId>/', MediaAssetView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:mediaAssetId>/', MediaAssetView.unassignVariant, name='unassignVariant'),
]
