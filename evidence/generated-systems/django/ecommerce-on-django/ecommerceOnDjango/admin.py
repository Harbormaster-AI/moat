from django.contrib import admin

# Register your models here.
from .models.Merchant import Merchant
from .models.Channel import Channel
from .models.Brand import Brand
from .models.Catalog import Catalog
from .models.Category import Category
from .models.Product import Product
from .models.ProductVariant import ProductVariant
from .models.ProductPricing import ProductPricing
from .models.MediaAsset import MediaAsset
from .models.FulfillmentCenter import FulfillmentCenter
from .models.InventoryItem import InventoryItem
from .models.Supplier import Supplier
from .models.Seller import Seller
from .models.Customer import Customer
from .models.CustomerAddress import CustomerAddress
from .models.Wishlist import Wishlist
from .models.WishlistItem import WishlistItem
from .models.Cart import Cart
from .models.CartItem import CartItem
from .models.Order import Order
from .models.OrderLine import OrderLine
from .models.Payment import Payment
from .models.Refund import Refund
from .models.Shipment import Shipment
from .models.ShipmentItem import ShipmentItem
from .models.ReturnRequest import ReturnRequest
from .models.ReturnItem import ReturnItem
from .models.Promotion import Promotion
from .models.Coupon import Coupon
from .models.CouponRedemption import CouponRedemption
from .models.TaxRule import TaxRule
from .models.ShippingMethod import ShippingMethod
from .models.CarrierService import CarrierService
from .models.Review import Review
from .models.Subscription import Subscription
from .models.PaymentProvider import PaymentProvider
from .models.Invoice import Invoice
from .models.GiftCard import GiftCard
from .models.GiftCardRedemption import GiftCardRedemption
from .models.Payout import Payout

# Need to add this for each model that requires managing

admin.site.register(Merchant)
admin.site.register(Channel)
admin.site.register(Brand)
admin.site.register(Catalog)
admin.site.register(Category)
admin.site.register(Product)
admin.site.register(ProductVariant)
admin.site.register(ProductPricing)
admin.site.register(MediaAsset)
admin.site.register(FulfillmentCenter)
admin.site.register(InventoryItem)
admin.site.register(Supplier)
admin.site.register(Seller)
admin.site.register(Customer)
admin.site.register(CustomerAddress)
admin.site.register(Wishlist)
admin.site.register(WishlistItem)
admin.site.register(Cart)
admin.site.register(CartItem)
admin.site.register(Order)
admin.site.register(OrderLine)
admin.site.register(Payment)
admin.site.register(Refund)
admin.site.register(Shipment)
admin.site.register(ShipmentItem)
admin.site.register(ReturnRequest)
admin.site.register(ReturnItem)
admin.site.register(Promotion)
admin.site.register(Coupon)
admin.site.register(CouponRedemption)
admin.site.register(TaxRule)
admin.site.register(ShippingMethod)
admin.site.register(CarrierService)
admin.site.register(Review)
admin.site.register(Subscription)
admin.site.register(PaymentProvider)
admin.site.register(Invoice)
admin.site.register(GiftCard)
admin.site.register(GiftCardRedemption)
admin.site.register(Payout)
