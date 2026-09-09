from django.urls import path
from advertisingOnDjango.views import BillingProfileView

urlpatterns = [
    path('', BillingProfileView.index, name='index'),
	path('create', BillingProfileView.get, name='create'),
	path('get/<int:billingProfileId>/', BillingProfileView.get, name='get'),
	path('save', BillingProfileView.save, name='save'),
	path('getAll', BillingProfileView.getAll, name='getAll'),
	path('delete/<int:billingProfileId>/', BillingProfileView.delete, name='delete'),
	path('assignAdvertiser/<int:billingProfileId>/<int:AdvertiserId>/', BillingProfileView.assignAdvertiser, name='assignAdvertiser'),
	path('unassignAdvertiser/<int:billingProfileId>/', BillingProfileView.unassignAdvertiser, name='unassignAdvertiser'),
	path('addPaymentMethods/<int:billingProfileId>/<PaymentMethodsIds>/', BillingProfileView.addPaymentMethods, name='addPaymentMethods'),
	path('removePaymentMethods/<int:billingProfileId>/<PaymentMethodsIds>/', BillingProfileView.removePaymentMethods, name='removePaymentMethods'),
	path('addAdAccounts/<int:billingProfileId>/<AdAccountsIds>/', BillingProfileView.addAdAccounts, name='addAdAccounts'),
	path('removeAdAccounts/<int:billingProfileId>/<AdAccountsIds>/', BillingProfileView.removeAdAccounts, name='removeAdAccounts'),
]
