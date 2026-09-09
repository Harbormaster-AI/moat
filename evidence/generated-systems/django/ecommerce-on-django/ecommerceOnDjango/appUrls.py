"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('Merchant/', include('ecommerceOnDjango.urls.MerchantUrls')),
    path('Channel/', include('ecommerceOnDjango.urls.ChannelUrls')),
    path('Brand/', include('ecommerceOnDjango.urls.BrandUrls')),
    path('Catalog/', include('ecommerceOnDjango.urls.CatalogUrls')),
    path('Category/', include('ecommerceOnDjango.urls.CategoryUrls')),
    path('Product/', include('ecommerceOnDjango.urls.ProductUrls')),
    path('ProductVariant/', include('ecommerceOnDjango.urls.ProductVariantUrls')),
    path('ProductPricing/', include('ecommerceOnDjango.urls.ProductPricingUrls')),
    path('MediaAsset/', include('ecommerceOnDjango.urls.MediaAssetUrls')),
    path('FulfillmentCenter/', include('ecommerceOnDjango.urls.FulfillmentCenterUrls')),
    path('InventoryItem/', include('ecommerceOnDjango.urls.InventoryItemUrls')),
    path('Supplier/', include('ecommerceOnDjango.urls.SupplierUrls')),
    path('Seller/', include('ecommerceOnDjango.urls.SellerUrls')),
    path('Customer/', include('ecommerceOnDjango.urls.CustomerUrls')),
    path('CustomerAddress/', include('ecommerceOnDjango.urls.CustomerAddressUrls')),
    path('Wishlist/', include('ecommerceOnDjango.urls.WishlistUrls')),
    path('WishlistItem/', include('ecommerceOnDjango.urls.WishlistItemUrls')),
    path('Cart/', include('ecommerceOnDjango.urls.CartUrls')),
    path('CartItem/', include('ecommerceOnDjango.urls.CartItemUrls')),
    path('Order/', include('ecommerceOnDjango.urls.OrderUrls')),
    path('OrderLine/', include('ecommerceOnDjango.urls.OrderLineUrls')),
    path('Payment/', include('ecommerceOnDjango.urls.PaymentUrls')),
    path('Refund/', include('ecommerceOnDjango.urls.RefundUrls')),
    path('Shipment/', include('ecommerceOnDjango.urls.ShipmentUrls')),
    path('ShipmentItem/', include('ecommerceOnDjango.urls.ShipmentItemUrls')),
    path('ReturnRequest/', include('ecommerceOnDjango.urls.ReturnRequestUrls')),
    path('ReturnItem/', include('ecommerceOnDjango.urls.ReturnItemUrls')),
    path('Promotion/', include('ecommerceOnDjango.urls.PromotionUrls')),
    path('Coupon/', include('ecommerceOnDjango.urls.CouponUrls')),
    path('CouponRedemption/', include('ecommerceOnDjango.urls.CouponRedemptionUrls')),
    path('TaxRule/', include('ecommerceOnDjango.urls.TaxRuleUrls')),
    path('ShippingMethod/', include('ecommerceOnDjango.urls.ShippingMethodUrls')),
    path('CarrierService/', include('ecommerceOnDjango.urls.CarrierServiceUrls')),
    path('Review/', include('ecommerceOnDjango.urls.ReviewUrls')),
    path('Subscription/', include('ecommerceOnDjango.urls.SubscriptionUrls')),
    path('PaymentProvider/', include('ecommerceOnDjango.urls.PaymentProviderUrls')),
    path('Invoice/', include('ecommerceOnDjango.urls.InvoiceUrls')),
    path('GiftCard/', include('ecommerceOnDjango.urls.GiftCardUrls')),
    path('GiftCardRedemption/', include('ecommerceOnDjango.urls.GiftCardRedemptionUrls')),
    path('Payout/', include('ecommerceOnDjango.urls.PayoutUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]