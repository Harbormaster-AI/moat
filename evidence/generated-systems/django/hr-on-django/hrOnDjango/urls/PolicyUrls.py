from django.urls import path
from hrOnDjango.views import PolicyView

urlpatterns = [
    path('', PolicyView.index, name='index'),
	path('create', PolicyView.get, name='create'),
	path('get/<int:policyId>/', PolicyView.get, name='get'),
	path('save', PolicyView.save, name='save'),
	path('getAll', PolicyView.getAll, name='getAll'),
	path('delete/<int:policyId>/', PolicyView.delete, name='delete'),
	path('assignOrganization/<int:policyId>/<int:OrganizationId>/', PolicyView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:policyId>/', PolicyView.unassignOrganization, name='unassignOrganization'),
	path('addAcknowledgements/<int:policyId>/<AcknowledgementsIds>/', PolicyView.addAcknowledgements, name='addAcknowledgements'),
	path('removeAcknowledgements/<int:policyId>/<AcknowledgementsIds>/', PolicyView.removeAcknowledgements, name='removeAcknowledgements'),
]
