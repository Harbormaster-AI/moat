from django.urls import path
from ecommerceOnDjango.views import PaymentProviderView

urlpatterns = [
    path('', PaymentProviderView.index, name='index'),
	path('create', PaymentProviderView.get, name='create'),
	path('get/<int:paymentProviderId>/', PaymentProviderView.get, name='get'),
	path('save', PaymentProviderView.save, name='save'),
	path('getAll', PaymentProviderView.getAll, name='getAll'),
	path('delete/<int:paymentProviderId>/', PaymentProviderView.delete, name='delete'),
	path('assignMerchant/<int:paymentProviderId>/<int:MerchantId>/', PaymentProviderView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:paymentProviderId>/', PaymentProviderView.unassignMerchant, name='unassignMerchant'),
	path('addChannels/<int:paymentProviderId>/<ChannelsIds>/', PaymentProviderView.addChannels, name='addChannels'),
	path('removeChannels/<int:paymentProviderId>/<ChannelsIds>/', PaymentProviderView.removeChannels, name='removeChannels'),
	path('addPayments/<int:paymentProviderId>/<PaymentsIds>/', PaymentProviderView.addPayments, name='addPayments'),
	path('removePayments/<int:paymentProviderId>/<PaymentsIds>/', PaymentProviderView.removePayments, name='removePayments'),
	path('addSubscriptions/<int:paymentProviderId>/<SubscriptionsIds>/', PaymentProviderView.addSubscriptions, name='addSubscriptions'),
	path('removeSubscriptions/<int:paymentProviderId>/<SubscriptionsIds>/', PaymentProviderView.removeSubscriptions, name='removeSubscriptions'),
]
