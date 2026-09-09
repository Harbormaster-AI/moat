from django.urls import path
from healthcareOnDjango.views import LaboratoryView

urlpatterns = [
    path('', LaboratoryView.index, name='index'),
	path('create', LaboratoryView.get, name='create'),
	path('get/<int:laboratoryId>/', LaboratoryView.get, name='get'),
	path('save', LaboratoryView.save, name='save'),
	path('getAll', LaboratoryView.getAll, name='getAll'),
	path('delete/<int:laboratoryId>/', LaboratoryView.delete, name='delete'),
	path('assignFacility/<int:laboratoryId>/<int:FacilityId>/', LaboratoryView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:laboratoryId>/', LaboratoryView.unassignFacility, name='unassignFacility'),
	path('addLaboratoryOrders/<int:laboratoryId>/<LaboratoryOrdersIds>/', LaboratoryView.addLaboratoryOrders, name='addLaboratoryOrders'),
	path('removeLaboratoryOrders/<int:laboratoryId>/<LaboratoryOrdersIds>/', LaboratoryView.removeLaboratoryOrders, name='removeLaboratoryOrders'),
	path('addLabResults/<int:laboratoryId>/<LabResultsIds>/', LaboratoryView.addLabResults, name='addLabResults'),
	path('removeLabResults/<int:laboratoryId>/<LabResultsIds>/', LaboratoryView.removeLabResults, name='removeLabResults'),
]
