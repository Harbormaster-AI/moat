from django.urls import path
from aerospaceOnDjango.views import LandingGearView

urlpatterns = [
    path('', LandingGearView.index, name='index'),
	path('create', LandingGearView.get, name='create'),
	path('get/<int:landingGearId>/', LandingGearView.get, name='get'),
	path('save', LandingGearView.save, name='save'),
	path('getAll', LandingGearView.getAll, name='getAll'),
	path('delete/<int:landingGearId>/', LandingGearView.delete, name='delete'),
	path('assignSupplier/<int:landingGearId>/<int:SupplierId>/', LandingGearView.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:landingGearId>/', LandingGearView.unassignSupplier, name='unassignSupplier'),
	path('addVariants/<int:landingGearId>/<VariantsIds>/', LandingGearView.addVariants, name='addVariants'),
	path('removeVariants/<int:landingGearId>/<VariantsIds>/', LandingGearView.removeVariants, name='removeVariants'),
]
