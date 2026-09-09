from django.urls import path
from ecommerceOnDjango.views import WishlistView

urlpatterns = [
    path('', WishlistView.index, name='index'),
	path('create', WishlistView.get, name='create'),
	path('get/<int:wishlistId>/', WishlistView.get, name='get'),
	path('save', WishlistView.save, name='save'),
	path('getAll', WishlistView.getAll, name='getAll'),
	path('delete/<int:wishlistId>/', WishlistView.delete, name='delete'),
	path('assignCustomer/<int:wishlistId>/<int:CustomerId>/', WishlistView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:wishlistId>/', WishlistView.unassignCustomer, name='unassignCustomer'),
	path('addItems/<int:wishlistId>/<ItemsIds>/', WishlistView.addItems, name='addItems'),
	path('removeItems/<int:wishlistId>/<ItemsIds>/', WishlistView.removeItems, name='removeItems'),
]
