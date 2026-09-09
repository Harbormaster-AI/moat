from django.urls import path
from aerospaceOnDjango.views import EngineTypeView

urlpatterns = [
    path('', EngineTypeView.index, name='index'),
	path('create', EngineTypeView.get, name='create'),
	path('get/<int:engineTypeId>/', EngineTypeView.get, name='get'),
	path('save', EngineTypeView.save, name='save'),
	path('getAll', EngineTypeView.getAll, name='getAll'),
	path('delete/<int:engineTypeId>/', EngineTypeView.delete, name='delete'),
	path('assignSupplier/<int:engineTypeId>/<int:SupplierId>/', EngineTypeView.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:engineTypeId>/', EngineTypeView.unassignSupplier, name='unassignSupplier'),
	path('addCompatibleModels/<int:engineTypeId>/<CompatibleModelsIds>/', EngineTypeView.addCompatibleModels, name='addCompatibleModels'),
	path('removeCompatibleModels/<int:engineTypeId>/<CompatibleModelsIds>/', EngineTypeView.removeCompatibleModels, name='removeCompatibleModels'),
]
