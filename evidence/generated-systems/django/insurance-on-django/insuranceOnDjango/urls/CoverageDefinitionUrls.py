from django.urls import path
from insuranceOnDjango.views import CoverageDefinitionView

urlpatterns = [
    path('', CoverageDefinitionView.index, name='index'),
	path('create', CoverageDefinitionView.get, name='create'),
	path('get/<int:coverageDefinitionId>/', CoverageDefinitionView.get, name='get'),
	path('save', CoverageDefinitionView.save, name='save'),
	path('getAll', CoverageDefinitionView.getAll, name='getAll'),
	path('delete/<int:coverageDefinitionId>/', CoverageDefinitionView.delete, name='delete'),
	path('assignProduct/<int:coverageDefinitionId>/<int:ProductId>/', CoverageDefinitionView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:coverageDefinitionId>/', CoverageDefinitionView.unassignProduct, name='unassignProduct'),
]
