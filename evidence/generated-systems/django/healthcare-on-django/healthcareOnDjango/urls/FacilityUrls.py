from django.urls import path
from healthcareOnDjango.views import FacilityView

urlpatterns = [
    path('', FacilityView.index, name='index'),
	path('create', FacilityView.get, name='create'),
	path('get/<int:facilityId>/', FacilityView.get, name='get'),
	path('save', FacilityView.save, name='save'),
	path('getAll', FacilityView.getAll, name='getAll'),
	path('delete/<int:facilityId>/', FacilityView.delete, name='delete'),
	path('assignHealthSystem/<int:facilityId>/<int:HealthSystemId>/', FacilityView.assignHealthSystem, name='assignHealthSystem'),
	path('unassignHealthSystem/<int:facilityId>/', FacilityView.unassignHealthSystem, name='unassignHealthSystem'),
	path('addDepartments/<int:facilityId>/<DepartmentsIds>/', FacilityView.addDepartments, name='addDepartments'),
	path('removeDepartments/<int:facilityId>/<DepartmentsIds>/', FacilityView.removeDepartments, name='removeDepartments'),
	path('addCareTeams/<int:facilityId>/<CareTeamsIds>/', FacilityView.addCareTeams, name='addCareTeams'),
	path('removeCareTeams/<int:facilityId>/<CareTeamsIds>/', FacilityView.removeCareTeams, name='removeCareTeams'),
	path('addLaboratories/<int:facilityId>/<LaboratoriesIds>/', FacilityView.addLaboratories, name='addLaboratories'),
	path('removeLaboratories/<int:facilityId>/<LaboratoriesIds>/', FacilityView.removeLaboratories, name='removeLaboratories'),
	path('addImagingCenters/<int:facilityId>/<ImagingCentersIds>/', FacilityView.addImagingCenters, name='addImagingCenters'),
	path('removeImagingCenters/<int:facilityId>/<ImagingCentersIds>/', FacilityView.removeImagingCenters, name='removeImagingCenters'),
	path('addPharmacies/<int:facilityId>/<PharmaciesIds>/', FacilityView.addPharmacies, name='addPharmacies'),
	path('removePharmacies/<int:facilityId>/<PharmaciesIds>/', FacilityView.removePharmacies, name='removePharmacies'),
	path('addInventoryItems/<int:facilityId>/<InventoryItemsIds>/', FacilityView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:facilityId>/<InventoryItemsIds>/', FacilityView.removeInventoryItems, name='removeInventoryItems'),
]
