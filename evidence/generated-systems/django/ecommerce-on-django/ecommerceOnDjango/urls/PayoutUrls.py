from django.urls import path
from ecommerceOnDjango.views import PayoutView

urlpatterns = [
    path('', PayoutView.index, name='index'),
	path('create', PayoutView.get, name='create'),
	path('get/<int:payoutId>/', PayoutView.get, name='get'),
	path('save', PayoutView.save, name='save'),
	path('getAll', PayoutView.getAll, name='getAll'),
	path('delete/<int:payoutId>/', PayoutView.delete, name='delete'),
	path('assignSeller/<int:payoutId>/<int:SellerId>/', PayoutView.assignSeller, name='assignSeller'),
	path('unassignSeller/<int:payoutId>/', PayoutView.unassignSeller, name='unassignSeller'),
	path('addOrders/<int:payoutId>/<OrdersIds>/', PayoutView.addOrders, name='addOrders'),
	path('removeOrders/<int:payoutId>/<OrdersIds>/', PayoutView.removeOrders, name='removeOrders'),
]
