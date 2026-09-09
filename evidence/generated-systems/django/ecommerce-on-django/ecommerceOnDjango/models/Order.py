from django.db import models
from ecommerceOnDjango.models.OrderStatus import OrderStatus

#======================================================================
# 
# Encapsulates data for model Order
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Order Declaration
#======================================================================
class Order (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	placedDate = models.DateField(null=True)
	subtotal = Money
	discountTotal = Money
	shippingTotal = Money
	taxTotal = Money
	grandTotal = Money
	shippingAddress = Address
	billingAddress = Address
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channel = models.ForeignKey('Channel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orderLines = models.ManyToManyField('OrderLine',  blank=True, related_name='+')
	payments = models.ManyToManyField('Payment',  blank=True, related_name='+')
	shipments = models.ManyToManyField('Shipment',  blank=True, related_name='+')
	refunds = models.ManyToManyField('Refund',  blank=True, related_name='+')
	appliedPromotions = models.ManyToManyField('Promotion',  blank=True, related_name='+')
	seller = models.ForeignKey('Seller', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	giftCardRedemptions = models.ManyToManyField('GiftCardRedemption',  blank=True, related_name='+')
	couponRedemptions = models.ManyToManyField('CouponRedemption',  blank=True, related_name='+')
	returnRequests = models.ManyToManyField('ReturnRequest',  blank=True, related_name='+')
	invoice = models.OneToOneField('Invoice', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.placedDate
		str = str + self.subtotal
		str = str + self.discountTotal
		str = str + self.shippingTotal
		str = str + self.taxTotal
		str = str + self.grandTotal
		str = str + self.shippingAddress
		str = str + self.billingAddress
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Order";
    
	def objectType(self):
		return "Order";
