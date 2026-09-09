from django.urls import path
from manufacturingOnDjango.views import BusinessUnitView

urlpatterns = [
    path('', BusinessUnitView.index, name='index'),
	path('create', BusinessUnitView.get, name='create'),
	path('get/<int:businessUnitId>/', BusinessUnitView.get, name='get'),
	path('save', BusinessUnitView.save, name='save'),
	path('getAll', BusinessUnitView.getAll, name='getAll'),
	path('delete/<int:businessUnitId>/', BusinessUnitView.delete, name='delete'),
	path('assignEnterprise/<int:businessUnitId>/<int:EnterpriseId>/', BusinessUnitView.assignEnterprise, name='assignEnterprise'),
	path('unassignEnterprise/<int:businessUnitId>/', BusinessUnitView.unassignEnterprise, name='unassignEnterprise'),
	path('addItems/<int:businessUnitId>/<ItemsIds>/', BusinessUnitView.addItems, name='addItems'),
	path('removeItems/<int:businessUnitId>/<ItemsIds>/', BusinessUnitView.removeItems, name='removeItems'),
	path('addPlants/<int:businessUnitId>/<PlantsIds>/', BusinessUnitView.addPlants, name='addPlants'),
	path('removePlants/<int:businessUnitId>/<PlantsIds>/', BusinessUnitView.removePlants, name='removePlants'),
]
