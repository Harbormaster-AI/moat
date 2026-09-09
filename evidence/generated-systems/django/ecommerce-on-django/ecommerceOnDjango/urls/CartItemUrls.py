from django.urls import path
from ecommerceOnDjango.views import CartItemView

urlpatterns = [
    path('', CartItemView.index, name='index'),
	path('create', CartItemView.get, name='create'),
	path('get/<int:cartItemId>/', CartItemView.get, name='get'),
	path('save', CartItemView.save, name='save'),
	path('getAll', CartItemView.getAll, name='getAll'),
	path('delete/<int:cartItemId>/', CartItemView.delete, name='delete'),
	path('assignCart/<int:cartItemId>/<int:CartId>/', CartItemView.assignCart, name='assignCart'),
	path('unassignCart/<int:cartItemId>/', CartItemView.unassignCart, name='unassignCart'),
	path('assignVariant/<int:cartItemId>/<int:VariantId>/', CartItemView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:cartItemId>/', CartItemView.unassignVariant, name='unassignVariant'),
	path('addAppliedPromotions/<int:cartItemId>/<AppliedPromotionsIds>/', CartItemView.addAppliedPromotions, name='addAppliedPromotions'),
	path('removeAppliedPromotions/<int:cartItemId>/<AppliedPromotionsIds>/', CartItemView.removeAppliedPromotions, name='removeAppliedPromotions'),
]
