from django.urls import path
from aerospaceOnDjango.views import AvionicsSuiteView

urlpatterns = [
    path('', AvionicsSuiteView.index, name='index'),
	path('create', AvionicsSuiteView.get, name='create'),
	path('get/<int:avionicsSuiteId>/', AvionicsSuiteView.get, name='get'),
	path('save', AvionicsSuiteView.save, name='save'),
	path('getAll', AvionicsSuiteView.getAll, name='getAll'),
	path('delete/<int:avionicsSuiteId>/', AvionicsSuiteView.delete, name='delete'),
	path('assignSupplier/<int:avionicsSuiteId>/<int:SupplierId>/', AvionicsSuiteView.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:avionicsSuiteId>/', AvionicsSuiteView.unassignSupplier, name='unassignSupplier'),
	path('addVariants/<int:avionicsSuiteId>/<VariantsIds>/', AvionicsSuiteView.addVariants, name='addVariants'),
	path('removeVariants/<int:avionicsSuiteId>/<VariantsIds>/', AvionicsSuiteView.removeVariants, name='removeVariants'),
	path('addSoftwareLoads/<int:avionicsSuiteId>/<SoftwareLoadsIds>/', AvionicsSuiteView.addSoftwareLoads, name='addSoftwareLoads'),
	path('removeSoftwareLoads/<int:avionicsSuiteId>/<SoftwareLoadsIds>/', AvionicsSuiteView.removeSoftwareLoads, name='removeSoftwareLoads'),
]
