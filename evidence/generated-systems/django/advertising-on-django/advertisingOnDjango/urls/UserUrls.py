from django.urls import path
from advertisingOnDjango.views import UserView

urlpatterns = [
    path('', UserView.index, name='index'),
	path('create', UserView.get, name='create'),
	path('get/<int:userId>/', UserView.get, name='get'),
	path('save', UserView.save, name='save'),
	path('getAll', UserView.getAll, name='getAll'),
	path('delete/<int:userId>/', UserView.delete, name='delete'),
	path('assignAgency/<int:userId>/<int:AgencyId>/', UserView.assignAgency, name='assignAgency'),
	path('unassignAgency/<int:userId>/', UserView.unassignAgency, name='unassignAgency'),
	path('addTeams/<int:userId>/<TeamsIds>/', UserView.addTeams, name='addTeams'),
	path('removeTeams/<int:userId>/<TeamsIds>/', UserView.removeTeams, name='removeTeams'),
	path('addAdAccounts/<int:userId>/<AdAccountsIds>/', UserView.addAdAccounts, name='addAdAccounts'),
	path('removeAdAccounts/<int:userId>/<AdAccountsIds>/', UserView.removeAdAccounts, name='removeAdAccounts'),
]
