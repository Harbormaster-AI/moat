from django.urls import path
from crmOnDjango.views import OrderItemView

urlpatterns = [
    path('', OrderItemView.index, name='index'),
	path('create', OrderItemView.get, name='create'),
	path('get/<int:orderItemId>/', OrderItemView.get, name='get'),
	path('save', OrderItemView.save, name='save'),
	path('getAll', OrderItemView.getAll, name='getAll'),
	path('delete/<int:orderItemId>/', OrderItemView.delete, name='delete'),
	path('assignOrder/<int:orderItemId>/<int:OrderId>/', OrderItemView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:orderItemId>/', OrderItemView.unassignOrder, name='unassignOrder'),
	path('assignProduct/<int:orderItemId>/<int:ProductId>/', OrderItemView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:orderItemId>/', OrderItemView.unassignProduct, name='unassignProduct'),
	path('assignPriceBookEntry/<int:orderItemId>/<int:PriceBookEntryId>/', OrderItemView.assignPriceBookEntry, name='assignPriceBookEntry'),
	path('unassignPriceBookEntry/<int:orderItemId>/', OrderItemView.unassignPriceBookEntry, name='unassignPriceBookEntry'),
]
