from django.urls import path
from inventoryOnDjango.views import UoMConversionView

urlpatterns = [
    path('', UoMConversionView.index, name='index'),
	path('create', UoMConversionView.get, name='create'),
	path('get/<int:uoMConversionId>/', UoMConversionView.get, name='get'),
	path('save', UoMConversionView.save, name='save'),
	path('getAll', UoMConversionView.getAll, name='getAll'),
	path('delete/<int:uoMConversionId>/', UoMConversionView.delete, name='delete'),
	path('assignSku/<int:uoMConversionId>/<int:SkuId>/', UoMConversionView.assignSku, name='assignSku'),
	path('unassignSku/<int:uoMConversionId>/', UoMConversionView.unassignSku, name='unassignSku'),
]
