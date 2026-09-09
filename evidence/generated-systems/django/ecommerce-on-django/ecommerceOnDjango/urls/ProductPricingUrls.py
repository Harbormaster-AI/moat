from django.urls import path
from ecommerceOnDjango.views import ProductPricingView

urlpatterns = [
    path('', ProductPricingView.index, name='index'),
	path('create', ProductPricingView.get, name='create'),
	path('get/<int:productPricingId>/', ProductPricingView.get, name='get'),
	path('save', ProductPricingView.save, name='save'),
	path('getAll', ProductPricingView.getAll, name='getAll'),
	path('delete/<int:productPricingId>/', ProductPricingView.delete, name='delete'),
	path('assignVariant/<int:productPricingId>/<int:VariantId>/', ProductPricingView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:productPricingId>/', ProductPricingView.unassignVariant, name='unassignVariant'),
	path('assignChannel/<int:productPricingId>/<int:ChannelId>/', ProductPricingView.assignChannel, name='assignChannel'),
	path('unassignChannel/<int:productPricingId>/', ProductPricingView.unassignChannel, name='unassignChannel'),
]
