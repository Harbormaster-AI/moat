from django.urls import path
from insuranceOnDjango.views import InsuranceProductView

urlpatterns = [
    path('', InsuranceProductView.index, name='index'),
	path('create', InsuranceProductView.get, name='create'),
	path('get/<int:insuranceProductId>/', InsuranceProductView.get, name='get'),
	path('save', InsuranceProductView.save, name='save'),
	path('getAll', InsuranceProductView.getAll, name='getAll'),
	path('delete/<int:insuranceProductId>/', InsuranceProductView.delete, name='delete'),
	path('assignInsurer/<int:insuranceProductId>/<int:InsurerId>/', InsuranceProductView.assignInsurer, name='assignInsurer'),
	path('unassignInsurer/<int:insuranceProductId>/', InsuranceProductView.unassignInsurer, name='unassignInsurer'),
	path('addCoverageDefinitions/<int:insuranceProductId>/<CoverageDefinitionsIds>/', InsuranceProductView.addCoverageDefinitions, name='addCoverageDefinitions'),
	path('removeCoverageDefinitions/<int:insuranceProductId>/<CoverageDefinitionsIds>/', InsuranceProductView.removeCoverageDefinitions, name='removeCoverageDefinitions'),
]
