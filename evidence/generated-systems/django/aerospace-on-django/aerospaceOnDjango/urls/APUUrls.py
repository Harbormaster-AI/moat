from django.urls import path
from aerospaceOnDjango.views import APUView

urlpatterns = [
    path('', APUView.index, name='index'),
	path('create', APUView.get, name='create'),
	path('get/<int:aPUId>/', APUView.get, name='get'),
	path('save', APUView.save, name='save'),
	path('getAll', APUView.getAll, name='getAll'),
	path('delete/<int:aPUId>/', APUView.delete, name='delete'),
	path('assignSupplier/<int:aPUId>/<int:SupplierId>/', APUView.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:aPUId>/', APUView.unassignSupplier, name='unassignSupplier'),
	path('addVariants/<int:aPUId>/<VariantsIds>/', APUView.addVariants, name='addVariants'),
	path('removeVariants/<int:aPUId>/<VariantsIds>/', APUView.removeVariants, name='removeVariants'),
]
