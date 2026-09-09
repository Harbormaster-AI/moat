from django.urls import path
from insuranceOnDjango.views import PolicyCoverageView

urlpatterns = [
    path('', PolicyCoverageView.index, name='index'),
	path('create', PolicyCoverageView.get, name='create'),
	path('get/<int:policyCoverageId>/', PolicyCoverageView.get, name='get'),
	path('save', PolicyCoverageView.save, name='save'),
	path('getAll', PolicyCoverageView.getAll, name='getAll'),
	path('delete/<int:policyCoverageId>/', PolicyCoverageView.delete, name='delete'),
	path('assignPolicy/<int:policyCoverageId>/<int:PolicyId>/', PolicyCoverageView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:policyCoverageId>/', PolicyCoverageView.unassignPolicy, name='unassignPolicy'),
	path('addInsuredObjects/<int:policyCoverageId>/<InsuredObjectsIds>/', PolicyCoverageView.addInsuredObjects, name='addInsuredObjects'),
	path('removeInsuredObjects/<int:policyCoverageId>/<InsuredObjectsIds>/', PolicyCoverageView.removeInsuredObjects, name='removeInsuredObjects'),
]
