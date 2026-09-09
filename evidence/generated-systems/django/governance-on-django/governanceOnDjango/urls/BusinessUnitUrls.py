from django.urls import path
from governanceOnDjango.views import BusinessUnitView

urlpatterns = [
    path('', BusinessUnitView.index, name='index'),
	path('create', BusinessUnitView.get, name='create'),
	path('get/<int:businessUnitId>/', BusinessUnitView.get, name='get'),
	path('save', BusinessUnitView.save, name='save'),
	path('getAll', BusinessUnitView.getAll, name='getAll'),
	path('delete/<int:businessUnitId>/', BusinessUnitView.delete, name='delete'),
	path('assignOrganization/<int:businessUnitId>/<int:OrganizationId>/', BusinessUnitView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:businessUnitId>/', BusinessUnitView.unassignOrganization, name='unassignOrganization'),
	path('addAudits/<int:businessUnitId>/<AuditsIds>/', BusinessUnitView.addAudits, name='addAudits'),
	path('removeAudits/<int:businessUnitId>/<AuditsIds>/', BusinessUnitView.removeAudits, name='removeAudits'),
]
