from django.urls import path
from advertisingOnDjango.views import AgencyView

urlpatterns = [
    path('', AgencyView.index, name='index'),
	path('create', AgencyView.get, name='create'),
	path('get/<int:agencyId>/', AgencyView.get, name='get'),
	path('save', AgencyView.save, name='save'),
	path('getAll', AgencyView.getAll, name='getAll'),
	path('delete/<int:agencyId>/', AgencyView.delete, name='delete'),
	path('addAdvertisers/<int:agencyId>/<AdvertisersIds>/', AgencyView.addAdvertisers, name='addAdvertisers'),
	path('removeAdvertisers/<int:agencyId>/<AdvertisersIds>/', AgencyView.removeAdvertisers, name='removeAdvertisers'),
	path('addTeams/<int:agencyId>/<TeamsIds>/', AgencyView.addTeams, name='addTeams'),
	path('removeTeams/<int:agencyId>/<TeamsIds>/', AgencyView.removeTeams, name='removeTeams'),
	path('addUsers/<int:agencyId>/<UsersIds>/', AgencyView.addUsers, name='addUsers'),
	path('removeUsers/<int:agencyId>/<UsersIds>/', AgencyView.removeUsers, name='removeUsers'),
	path('addInsertionOrders/<int:agencyId>/<InsertionOrdersIds>/', AgencyView.addInsertionOrders, name='addInsertionOrders'),
	path('removeInsertionOrders/<int:agencyId>/<InsertionOrdersIds>/', AgencyView.removeInsertionOrders, name='removeInsertionOrders'),
]
