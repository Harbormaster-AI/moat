from django.urls import path
from advertisingOnDjango.views import TeamView

urlpatterns = [
    path('', TeamView.index, name='index'),
	path('create', TeamView.get, name='create'),
	path('get/<int:teamId>/', TeamView.get, name='get'),
	path('save', TeamView.save, name='save'),
	path('getAll', TeamView.getAll, name='getAll'),
	path('delete/<int:teamId>/', TeamView.delete, name='delete'),
	path('assignAgency/<int:teamId>/<int:AgencyId>/', TeamView.assignAgency, name='assignAgency'),
	path('unassignAgency/<int:teamId>/', TeamView.unassignAgency, name='unassignAgency'),
	path('addUsers/<int:teamId>/<UsersIds>/', TeamView.addUsers, name='addUsers'),
	path('removeUsers/<int:teamId>/<UsersIds>/', TeamView.removeUsers, name='removeUsers'),
	path('addAdAccounts/<int:teamId>/<AdAccountsIds>/', TeamView.addAdAccounts, name='addAdAccounts'),
	path('removeAdAccounts/<int:teamId>/<AdAccountsIds>/', TeamView.removeAdAccounts, name='removeAdAccounts'),
]
