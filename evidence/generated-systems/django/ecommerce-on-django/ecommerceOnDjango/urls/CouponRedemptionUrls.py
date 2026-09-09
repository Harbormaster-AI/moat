from django.urls import path
from ecommerceOnDjango.views import CouponRedemptionView

urlpatterns = [
    path('', CouponRedemptionView.index, name='index'),
	path('create', CouponRedemptionView.get, name='create'),
	path('get/<int:couponRedemptionId>/', CouponRedemptionView.get, name='get'),
	path('save', CouponRedemptionView.save, name='save'),
	path('getAll', CouponRedemptionView.getAll, name='getAll'),
	path('delete/<int:couponRedemptionId>/', CouponRedemptionView.delete, name='delete'),
	path('assignCoupon/<int:couponRedemptionId>/<int:CouponId>/', CouponRedemptionView.assignCoupon, name='assignCoupon'),
	path('unassignCoupon/<int:couponRedemptionId>/', CouponRedemptionView.unassignCoupon, name='unassignCoupon'),
	path('assignOrder/<int:couponRedemptionId>/<int:OrderId>/', CouponRedemptionView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:couponRedemptionId>/', CouponRedemptionView.unassignOrder, name='unassignOrder'),
	path('assignCustomer/<int:couponRedemptionId>/<int:CustomerId>/', CouponRedemptionView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:couponRedemptionId>/', CouponRedemptionView.unassignCustomer, name='unassignCustomer'),
]
