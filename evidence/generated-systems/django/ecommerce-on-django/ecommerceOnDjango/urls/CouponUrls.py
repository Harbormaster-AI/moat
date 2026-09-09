from django.urls import path
from ecommerceOnDjango.views import CouponView

urlpatterns = [
    path('', CouponView.index, name='index'),
	path('create', CouponView.get, name='create'),
	path('get/<int:couponId>/', CouponView.get, name='get'),
	path('save', CouponView.save, name='save'),
	path('getAll', CouponView.getAll, name='getAll'),
	path('delete/<int:couponId>/', CouponView.delete, name='delete'),
	path('assignPromotion/<int:couponId>/<int:PromotionId>/', CouponView.assignPromotion, name='assignPromotion'),
	path('unassignPromotion/<int:couponId>/', CouponView.unassignPromotion, name='unassignPromotion'),
	path('addRedemptions/<int:couponId>/<RedemptionsIds>/', CouponView.addRedemptions, name='addRedemptions'),
	path('removeRedemptions/<int:couponId>/<RedemptionsIds>/', CouponView.removeRedemptions, name='removeRedemptions'),
]
