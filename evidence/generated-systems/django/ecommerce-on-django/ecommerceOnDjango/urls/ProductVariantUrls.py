from django.urls import path
from ecommerceOnDjango.views import ProductVariantView

urlpatterns = [
    path('', ProductVariantView.index, name='index'),
	path('create', ProductVariantView.get, name='create'),
	path('get/<int:productVariantId>/', ProductVariantView.get, name='get'),
	path('save', ProductVariantView.save, name='save'),
	path('getAll', ProductVariantView.getAll, name='getAll'),
	path('delete/<int:productVariantId>/', ProductVariantView.delete, name='delete'),
	path('assignProduct/<int:productVariantId>/<int:ProductId>/', ProductVariantView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:productVariantId>/', ProductVariantView.unassignProduct, name='unassignProduct'),
	path('addPricing/<int:productVariantId>/<PricingIds>/', ProductVariantView.addPricing, name='addPricing'),
	path('removePricing/<int:productVariantId>/<PricingIds>/', ProductVariantView.removePricing, name='removePricing'),
	path('addInventoryItems/<int:productVariantId>/<InventoryItemsIds>/', ProductVariantView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:productVariantId>/<InventoryItemsIds>/', ProductVariantView.removeInventoryItems, name='removeInventoryItems'),
	path('addMediaAssets/<int:productVariantId>/<MediaAssetsIds>/', ProductVariantView.addMediaAssets, name='addMediaAssets'),
	path('removeMediaAssets/<int:productVariantId>/<MediaAssetsIds>/', ProductVariantView.removeMediaAssets, name='removeMediaAssets'),
	path('addSubscriptions/<int:productVariantId>/<SubscriptionsIds>/', ProductVariantView.addSubscriptions, name='addSubscriptions'),
	path('removeSubscriptions/<int:productVariantId>/<SubscriptionsIds>/', ProductVariantView.removeSubscriptions, name='removeSubscriptions'),
	path('addCartItems/<int:productVariantId>/<CartItemsIds>/', ProductVariantView.addCartItems, name='addCartItems'),
	path('removeCartItems/<int:productVariantId>/<CartItemsIds>/', ProductVariantView.removeCartItems, name='removeCartItems'),
	path('addOrderLines/<int:productVariantId>/<OrderLinesIds>/', ProductVariantView.addOrderLines, name='addOrderLines'),
	path('removeOrderLines/<int:productVariantId>/<OrderLinesIds>/', ProductVariantView.removeOrderLines, name='removeOrderLines'),
	path('addWishlistItems/<int:productVariantId>/<WishlistItemsIds>/', ProductVariantView.addWishlistItems, name='addWishlistItems'),
	path('removeWishlistItems/<int:productVariantId>/<WishlistItemsIds>/', ProductVariantView.removeWishlistItems, name='removeWishlistItems'),
]
