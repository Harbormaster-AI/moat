from django.urls import path
from healthcareOnDjango.views import LaboratoryOrderView

urlpatterns = [
    path('', LaboratoryOrderView.index, name='index'),
	path('create', LaboratoryOrderView.get, name='create'),
	path('get/<int:laboratoryOrderId>/', LaboratoryOrderView.get, name='get'),
	path('save', LaboratoryOrderView.save, name='save'),
	path('getAll', LaboratoryOrderView.getAll, name='getAll'),
	path('delete/<int:laboratoryOrderId>/', LaboratoryOrderView.delete, name='delete'),
	path('assignOrder/<int:laboratoryOrderId>/<int:OrderId>/', LaboratoryOrderView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:laboratoryOrderId>/', LaboratoryOrderView.unassignOrder, name='unassignOrder'),
	path('assignLaboratory/<int:laboratoryOrderId>/<int:LaboratoryId>/', LaboratoryOrderView.assignLaboratory, name='assignLaboratory'),
	path('unassignLaboratory/<int:laboratoryOrderId>/', LaboratoryOrderView.unassignLaboratory, name='unassignLaboratory'),
	path('addResults/<int:laboratoryOrderId>/<ResultsIds>/', LaboratoryOrderView.addResults, name='addResults'),
	path('removeResults/<int:laboratoryOrderId>/<ResultsIds>/', LaboratoryOrderView.removeResults, name='removeResults'),
]
