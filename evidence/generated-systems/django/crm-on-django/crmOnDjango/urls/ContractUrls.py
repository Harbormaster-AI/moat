from django.urls import path
from crmOnDjango.views import ContractView

urlpatterns = [
    path('', ContractView.index, name='index'),
	path('create', ContractView.get, name='create'),
	path('get/<int:contractId>/', ContractView.get, name='get'),
	path('save', ContractView.save, name='save'),
	path('getAll', ContractView.getAll, name='getAll'),
	path('delete/<int:contractId>/', ContractView.delete, name='delete'),
	path('assignOrganization/<int:contractId>/<int:OrganizationId>/', ContractView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:contractId>/', ContractView.unassignOrganization, name='unassignOrganization'),
	path('assignAccount/<int:contractId>/<int:AccountId>/', ContractView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:contractId>/', ContractView.unassignAccount, name='unassignAccount'),
	path('assignOwner/<int:contractId>/<int:OwnerId>/', ContractView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:contractId>/', ContractView.unassignOwner, name='unassignOwner'),
	path('addOrders/<int:contractId>/<OrdersIds>/', ContractView.addOrders, name='addOrders'),
	path('removeOrders/<int:contractId>/<OrdersIds>/', ContractView.removeOrders, name='removeOrders'),
	path('addCases/<int:contractId>/<CasesIds>/', ContractView.addCases, name='addCases'),
	path('removeCases/<int:contractId>/<CasesIds>/', ContractView.removeCases, name='removeCases'),
]
