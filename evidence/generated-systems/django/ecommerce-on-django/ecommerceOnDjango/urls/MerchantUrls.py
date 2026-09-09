from django.urls import path
from ecommerceOnDjango.views import MerchantView

urlpatterns = [
    path('', MerchantView.index, name='index'),
	path('create', MerchantView.get, name='create'),
	path('get/<int:merchantId>/', MerchantView.get, name='get'),
	path('save', MerchantView.save, name='save'),
	path('getAll', MerchantView.getAll, name='getAll'),
	path('delete/<int:merchantId>/', MerchantView.delete, name='delete'),
	path('addChannels/<int:merchantId>/<ChannelsIds>/', MerchantView.addChannels, name='addChannels'),
	path('removeChannels/<int:merchantId>/<ChannelsIds>/', MerchantView.removeChannels, name='removeChannels'),
	path('addBrands/<int:merchantId>/<BrandsIds>/', MerchantView.addBrands, name='addBrands'),
	path('removeBrands/<int:merchantId>/<BrandsIds>/', MerchantView.removeBrands, name='removeBrands'),
	path('addFulfillmentCenters/<int:merchantId>/<FulfillmentCentersIds>/', MerchantView.addFulfillmentCenters, name='addFulfillmentCenters'),
	path('removeFulfillmentCenters/<int:merchantId>/<FulfillmentCentersIds>/', MerchantView.removeFulfillmentCenters, name='removeFulfillmentCenters'),
	path('addTaxRules/<int:merchantId>/<TaxRulesIds>/', MerchantView.addTaxRules, name='addTaxRules'),
	path('removeTaxRules/<int:merchantId>/<TaxRulesIds>/', MerchantView.removeTaxRules, name='removeTaxRules'),
	path('addPaymentProviders/<int:merchantId>/<PaymentProvidersIds>/', MerchantView.addPaymentProviders, name='addPaymentProviders'),
	path('removePaymentProviders/<int:merchantId>/<PaymentProvidersIds>/', MerchantView.removePaymentProviders, name='removePaymentProviders'),
	path('addSellers/<int:merchantId>/<SellersIds>/', MerchantView.addSellers, name='addSellers'),
	path('removeSellers/<int:merchantId>/<SellersIds>/', MerchantView.removeSellers, name='removeSellers'),
	path('addPromotions/<int:merchantId>/<PromotionsIds>/', MerchantView.addPromotions, name='addPromotions'),
	path('removePromotions/<int:merchantId>/<PromotionsIds>/', MerchantView.removePromotions, name='removePromotions'),
]
