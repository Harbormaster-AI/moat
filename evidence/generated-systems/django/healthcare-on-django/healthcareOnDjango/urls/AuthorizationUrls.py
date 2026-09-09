from django.urls import path
from healthcareOnDjango.views import AuthorizationView

urlpatterns = [
    path('', AuthorizationView.index, name='index'),
	path('create', AuthorizationView.get, name='create'),
	path('get/<int:authorizationId>/', AuthorizationView.get, name='get'),
	path('save', AuthorizationView.save, name='save'),
	path('getAll', AuthorizationView.getAll, name='getAll'),
	path('delete/<int:authorizationId>/', AuthorizationView.delete, name='delete'),
	path('assignCoverage/<int:authorizationId>/<int:CoverageId>/', AuthorizationView.assignCoverage, name='assignCoverage'),
	path('unassignCoverage/<int:authorizationId>/', AuthorizationView.unassignCoverage, name='unassignCoverage'),
	path('assignOrder/<int:authorizationId>/<int:OrderId>/', AuthorizationView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:authorizationId>/', AuthorizationView.unassignOrder, name='unassignOrder'),
]
