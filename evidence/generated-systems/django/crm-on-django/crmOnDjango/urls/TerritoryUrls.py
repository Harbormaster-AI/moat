from django.urls import path
from crmOnDjango.views import TerritoryView

urlpatterns = [
    path('', TerritoryView.index, name='index'),
	path('create', TerritoryView.get, name='create'),
	path('get/<int:territoryId>/', TerritoryView.get, name='get'),
	path('save', TerritoryView.save, name='save'),
	path('getAll', TerritoryView.getAll, name='getAll'),
	path('delete/<int:territoryId>/', TerritoryView.delete, name='delete'),
	path('assignOrganization/<int:territoryId>/<int:OrganizationId>/', TerritoryView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:territoryId>/', TerritoryView.unassignOrganization, name='unassignOrganization'),
	path('addAccounts/<int:territoryId>/<AccountsIds>/', TerritoryView.addAccounts, name='addAccounts'),
	path('removeAccounts/<int:territoryId>/<AccountsIds>/', TerritoryView.removeAccounts, name='removeAccounts'),
	path('addUsers/<int:territoryId>/<UsersIds>/', TerritoryView.addUsers, name='addUsers'),
	path('removeUsers/<int:territoryId>/<UsersIds>/', TerritoryView.removeUsers, name='removeUsers'),
]
