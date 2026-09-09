from django.urls import path
from manufacturingOnDjango.views import EnterpriseView

urlpatterns = [
    path('', EnterpriseView.index, name='index'),
	path('create', EnterpriseView.get, name='create'),
	path('get/<int:enterpriseId>/', EnterpriseView.get, name='get'),
	path('save', EnterpriseView.save, name='save'),
	path('getAll', EnterpriseView.getAll, name='getAll'),
	path('delete/<int:enterpriseId>/', EnterpriseView.delete, name='delete'),
	path('addBusinessUnits/<int:enterpriseId>/<BusinessUnitsIds>/', EnterpriseView.addBusinessUnits, name='addBusinessUnits'),
	path('removeBusinessUnits/<int:enterpriseId>/<BusinessUnitsIds>/', EnterpriseView.removeBusinessUnits, name='removeBusinessUnits'),
	path('addPlants/<int:enterpriseId>/<PlantsIds>/', EnterpriseView.addPlants, name='addPlants'),
	path('removePlants/<int:enterpriseId>/<PlantsIds>/', EnterpriseView.removePlants, name='removePlants'),
	path('addSuppliers/<int:enterpriseId>/<SuppliersIds>/', EnterpriseView.addSuppliers, name='addSuppliers'),
	path('removeSuppliers/<int:enterpriseId>/<SuppliersIds>/', EnterpriseView.removeSuppliers, name='removeSuppliers'),
	path('addCustomers/<int:enterpriseId>/<CustomersIds>/', EnterpriseView.addCustomers, name='addCustomers'),
	path('removeCustomers/<int:enterpriseId>/<CustomersIds>/', EnterpriseView.removeCustomers, name='removeCustomers'),
]
