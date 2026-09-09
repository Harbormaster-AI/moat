from django.urls import path
from crmOnDjango.views import TeamView

urlpatterns = [
    path('', TeamView.index, name='index'),
	path('create', TeamView.get, name='create'),
	path('get/<int:teamId>/', TeamView.get, name='get'),
	path('save', TeamView.save, name='save'),
	path('getAll', TeamView.getAll, name='getAll'),
	path('delete/<int:teamId>/', TeamView.delete, name='delete'),
	path('assignOrganization/<int:teamId>/<int:OrganizationId>/', TeamView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:teamId>/', TeamView.unassignOrganization, name='unassignOrganization'),
	path('addUsers/<int:teamId>/<UsersIds>/', TeamView.addUsers, name='addUsers'),
	path('removeUsers/<int:teamId>/<UsersIds>/', TeamView.removeUsers, name='removeUsers'),
	path('addAccounts/<int:teamId>/<AccountsIds>/', TeamView.addAccounts, name='addAccounts'),
	path('removeAccounts/<int:teamId>/<AccountsIds>/', TeamView.removeAccounts, name='removeAccounts'),
	path('addOpportunities/<int:teamId>/<OpportunitiesIds>/', TeamView.addOpportunities, name='addOpportunities'),
	path('removeOpportunities/<int:teamId>/<OpportunitiesIds>/', TeamView.removeOpportunities, name='removeOpportunities'),
	path('addCases/<int:teamId>/<CasesIds>/', TeamView.addCases, name='addCases'),
	path('removeCases/<int:teamId>/<CasesIds>/', TeamView.removeCases, name='removeCases'),
	path('addCampaigns/<int:teamId>/<CampaignsIds>/', TeamView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:teamId>/<CampaignsIds>/', TeamView.removeCampaigns, name='removeCampaigns'),
]
