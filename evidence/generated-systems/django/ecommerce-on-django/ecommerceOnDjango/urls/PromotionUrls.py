from django.urls import path
from ecommerceOnDjango.views import PromotionView

urlpatterns = [
    path('', PromotionView.index, name='index'),
	path('create', PromotionView.get, name='create'),
	path('get/<int:promotionId>/', PromotionView.get, name='get'),
	path('save', PromotionView.save, name='save'),
	path('getAll', PromotionView.getAll, name='getAll'),
	path('delete/<int:promotionId>/', PromotionView.delete, name='delete'),
	path('assignMerchant/<int:promotionId>/<int:MerchantId>/', PromotionView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:promotionId>/', PromotionView.unassignMerchant, name='unassignMerchant'),
	path('addChannels/<int:promotionId>/<ChannelsIds>/', PromotionView.addChannels, name='addChannels'),
	path('removeChannels/<int:promotionId>/<ChannelsIds>/', PromotionView.removeChannels, name='removeChannels'),
	path('addApplicableProducts/<int:promotionId>/<ApplicableProductsIds>/', PromotionView.addApplicableProducts, name='addApplicableProducts'),
	path('removeApplicableProducts/<int:promotionId>/<ApplicableProductsIds>/', PromotionView.removeApplicableProducts, name='removeApplicableProducts'),
	path('addApplicableCategories/<int:promotionId>/<ApplicableCategoriesIds>/', PromotionView.addApplicableCategories, name='addApplicableCategories'),
	path('removeApplicableCategories/<int:promotionId>/<ApplicableCategoriesIds>/', PromotionView.removeApplicableCategories, name='removeApplicableCategories'),
	path('addCoupons/<int:promotionId>/<CouponsIds>/', PromotionView.addCoupons, name='addCoupons'),
	path('removeCoupons/<int:promotionId>/<CouponsIds>/', PromotionView.removeCoupons, name='removeCoupons'),
]
