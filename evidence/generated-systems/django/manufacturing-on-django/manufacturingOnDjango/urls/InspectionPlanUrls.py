from django.urls import path
from manufacturingOnDjango.views import InspectionPlanView

urlpatterns = [
    path('', InspectionPlanView.index, name='index'),
	path('create', InspectionPlanView.get, name='create'),
	path('get/<int:inspectionPlanId>/', InspectionPlanView.get, name='get'),
	path('save', InspectionPlanView.save, name='save'),
	path('getAll', InspectionPlanView.getAll, name='getAll'),
	path('delete/<int:inspectionPlanId>/', InspectionPlanView.delete, name='delete'),
	path('assignItem/<int:inspectionPlanId>/<int:ItemId>/', InspectionPlanView.assignItem, name='assignItem'),
	path('unassignItem/<int:inspectionPlanId>/', InspectionPlanView.unassignItem, name='unassignItem'),
	path('addCharacteristics/<int:inspectionPlanId>/<CharacteristicsIds>/', InspectionPlanView.addCharacteristics, name='addCharacteristics'),
	path('removeCharacteristics/<int:inspectionPlanId>/<CharacteristicsIds>/', InspectionPlanView.removeCharacteristics, name='removeCharacteristics'),
]
