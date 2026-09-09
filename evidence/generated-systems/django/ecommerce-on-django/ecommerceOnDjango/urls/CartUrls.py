from django.urls import path
from ecommerceOnDjango.views import CartView

urlpatterns = [
    path('', CartView.index, name='index'),
	path('create', CartView.get, name='create'),
	path('get/<int:cartId>/', CartView.get, name='get'),
	path('save', CartView.save, name='save'),
	path('getAll', CartView.getAll, name='getAll'),
	path('delete/<int:cartId>/', CartView.delete, name='delete'),
	path('assignCustomer/<int:cartId>/<int:CustomerId>/', CartView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:cartId>/', CartView.unassignCustomer, name='unassignCustomer'),
	path('assignChannel/<int:cartId>/<int:ChannelId>/', CartView.assignChannel, name='assignChannel'),
	path('unassignChannel/<int:cartId>/', CartView.unassignChannel, name='unassignChannel'),
	path('addItems/<int:cartId>/<ItemsIds>/', CartView.addItems, name='addItems'),
	path('removeItems/<int:cartId>/<ItemsIds>/', CartView.removeItems, name='removeItems'),
	path('addAppliedPromotions/<int:cartId>/<AppliedPromotionsIds>/', CartView.addAppliedPromotions, name='addAppliedPromotions'),
	path('removeAppliedPromotions/<int:cartId>/<AppliedPromotionsIds>/', CartView.removeAppliedPromotions, name='removeAppliedPromotions'),
]
