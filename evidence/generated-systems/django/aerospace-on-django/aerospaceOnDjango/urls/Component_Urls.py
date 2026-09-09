from django.urls import path
from aerospaceOnDjango.views import Component_View

urlpatterns = [
    path('', Component_View.index, name='index'),
	path('create', Component_View.get, name='create'),
	path('get/<int:component_Id>/', Component_View.get, name='get'),
	path('save', Component_View.save, name='save'),
	path('getAll', Component_View.getAll, name='getAll'),
	path('delete/<int:component_Id>/', Component_View.delete, name='delete'),
	path('assignSupplier/<int:component_Id>/<int:SupplierId>/', Component_View.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:component_Id>/', Component_View.unassignSupplier, name='unassignSupplier'),
]
