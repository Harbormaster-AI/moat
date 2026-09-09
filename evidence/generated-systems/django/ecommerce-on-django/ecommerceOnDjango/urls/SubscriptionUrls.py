from django.urls import path
from ecommerceOnDjango.views import SubscriptionView

urlpatterns = [
    path('', SubscriptionView.index, name='index'),
	path('create', SubscriptionView.get, name='create'),
	path('get/<int:subscriptionId>/', SubscriptionView.get, name='get'),
	path('save', SubscriptionView.save, name='save'),
	path('getAll', SubscriptionView.getAll, name='getAll'),
	path('delete/<int:subscriptionId>/', SubscriptionView.delete, name='delete'),
	path('assignCustomer/<int:subscriptionId>/<int:CustomerId>/', SubscriptionView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:subscriptionId>/', SubscriptionView.unassignCustomer, name='unassignCustomer'),
	path('assignVariant/<int:subscriptionId>/<int:VariantId>/', SubscriptionView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:subscriptionId>/', SubscriptionView.unassignVariant, name='unassignVariant'),
	path('assignPaymentProvider/<int:subscriptionId>/<int:PaymentProviderId>/', SubscriptionView.assignPaymentProvider, name='assignPaymentProvider'),
	path('unassignPaymentProvider/<int:subscriptionId>/', SubscriptionView.unassignPaymentProvider, name='unassignPaymentProvider'),
	path('assignChannel/<int:subscriptionId>/<int:ChannelId>/', SubscriptionView.assignChannel, name='assignChannel'),
	path('unassignChannel/<int:subscriptionId>/', SubscriptionView.unassignChannel, name='unassignChannel'),
]
