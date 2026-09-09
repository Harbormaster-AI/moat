from django.urls import path
from fintechOnDjango.views import KYCProfileView

urlpatterns = [
    path('', KYCProfileView.index, name='index'),
	path('create', KYCProfileView.get, name='create'),
	path('get/<int:kYCProfileId>/', KYCProfileView.get, name='get'),
	path('save', KYCProfileView.save, name='save'),
	path('getAll', KYCProfileView.getAll, name='getAll'),
	path('delete/<int:kYCProfileId>/', KYCProfileView.delete, name='delete'),
	path('assignCustomer/<int:kYCProfileId>/<int:CustomerId>/', KYCProfileView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:kYCProfileId>/', KYCProfileView.unassignCustomer, name='unassignCustomer'),
	path('addDocuments/<int:kYCProfileId>/<DocumentsIds>/', KYCProfileView.addDocuments, name='addDocuments'),
	path('removeDocuments/<int:kYCProfileId>/<DocumentsIds>/', KYCProfileView.removeDocuments, name='removeDocuments'),
	path('addScreenings/<int:kYCProfileId>/<ScreeningsIds>/', KYCProfileView.addScreenings, name='addScreenings'),
	path('removeScreenings/<int:kYCProfileId>/<ScreeningsIds>/', KYCProfileView.removeScreenings, name='removeScreenings'),
	path('addAddresses/<int:kYCProfileId>/<AddressesIds>/', KYCProfileView.addAddresses, name='addAddresses'),
	path('removeAddresses/<int:kYCProfileId>/<AddressesIds>/', KYCProfileView.removeAddresses, name='removeAddresses'),
]
