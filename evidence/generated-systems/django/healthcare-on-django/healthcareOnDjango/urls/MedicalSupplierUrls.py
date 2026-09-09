from django.urls import path
from healthcareOnDjango.views import MedicalSupplierView

urlpatterns = [
    path('', MedicalSupplierView.index, name='index'),
	path('create', MedicalSupplierView.get, name='create'),
	path('get/<int:medicalSupplierId>/', MedicalSupplierView.get, name='get'),
	path('save', MedicalSupplierView.save, name='save'),
	path('getAll', MedicalSupplierView.getAll, name='getAll'),
	path('delete/<int:medicalSupplierId>/', MedicalSupplierView.delete, name='delete'),
	path('addFacilities/<int:medicalSupplierId>/<FacilitiesIds>/', MedicalSupplierView.addFacilities, name='addFacilities'),
	path('removeFacilities/<int:medicalSupplierId>/<FacilitiesIds>/', MedicalSupplierView.removeFacilities, name='removeFacilities'),
	path('addInventoryItems/<int:medicalSupplierId>/<InventoryItemsIds>/', MedicalSupplierView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:medicalSupplierId>/<InventoryItemsIds>/', MedicalSupplierView.removeInventoryItems, name='removeInventoryItems'),
]
