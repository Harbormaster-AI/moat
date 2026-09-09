from django.urls import path
from crmOnDjango.views import OrganizationView

urlpatterns = [
    path('', OrganizationView.index, name='index'),
	path('create', OrganizationView.get, name='create'),
	path('get/<int:organizationId>/', OrganizationView.get, name='get'),
	path('save', OrganizationView.save, name='save'),
	path('getAll', OrganizationView.getAll, name='getAll'),
	path('delete/<int:organizationId>/', OrganizationView.delete, name='delete'),
	path('addUsers/<int:organizationId>/<UsersIds>/', OrganizationView.addUsers, name='addUsers'),
	path('removeUsers/<int:organizationId>/<UsersIds>/', OrganizationView.removeUsers, name='removeUsers'),
	path('addAccounts/<int:organizationId>/<AccountsIds>/', OrganizationView.addAccounts, name='addAccounts'),
	path('removeAccounts/<int:organizationId>/<AccountsIds>/', OrganizationView.removeAccounts, name='removeAccounts'),
	path('addTeams/<int:organizationId>/<TeamsIds>/', OrganizationView.addTeams, name='addTeams'),
	path('removeTeams/<int:organizationId>/<TeamsIds>/', OrganizationView.removeTeams, name='removeTeams'),
	path('addTerritories/<int:organizationId>/<TerritoriesIds>/', OrganizationView.addTerritories, name='addTerritories'),
	path('removeTerritories/<int:organizationId>/<TerritoriesIds>/', OrganizationView.removeTerritories, name='removeTerritories'),
	path('addProducts/<int:organizationId>/<ProductsIds>/', OrganizationView.addProducts, name='addProducts'),
	path('removeProducts/<int:organizationId>/<ProductsIds>/', OrganizationView.removeProducts, name='removeProducts'),
	path('addPriceBooks/<int:organizationId>/<PriceBooksIds>/', OrganizationView.addPriceBooks, name='addPriceBooks'),
	path('removePriceBooks/<int:organizationId>/<PriceBooksIds>/', OrganizationView.removePriceBooks, name='removePriceBooks'),
	path('addCampaigns/<int:organizationId>/<CampaignsIds>/', OrganizationView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:organizationId>/<CampaignsIds>/', OrganizationView.removeCampaigns, name='removeCampaigns'),
]
