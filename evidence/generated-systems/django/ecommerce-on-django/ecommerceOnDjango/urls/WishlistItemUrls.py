from django.urls import path
from ecommerceOnDjango.views import WishlistItemView

urlpatterns = [
    path('', WishlistItemView.index, name='index'),
	path('create', WishlistItemView.get, name='create'),
	path('get/<int:wishlistItemId>/', WishlistItemView.get, name='get'),
	path('save', WishlistItemView.save, name='save'),
	path('getAll', WishlistItemView.getAll, name='getAll'),
	path('delete/<int:wishlistItemId>/', WishlistItemView.delete, name='delete'),
	path('assignWishlist/<int:wishlistItemId>/<int:WishlistId>/', WishlistItemView.assignWishlist, name='assignWishlist'),
	path('unassignWishlist/<int:wishlistItemId>/', WishlistItemView.unassignWishlist, name='unassignWishlist'),
	path('assignVariant/<int:wishlistItemId>/<int:VariantId>/', WishlistItemView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:wishlistItemId>/', WishlistItemView.unassignVariant, name='unassignVariant'),
]
