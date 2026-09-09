from django.db import models
from crmOnDjango.models.OrderStatus import OrderStatus

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
	orderDate = models.DateField(null=True)
	totalAmount = Money
	taxAmount = Money
	shippingAmount = Money
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	quote = models.OneToOneField('Quote', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	items = models.ManyToManyField('OrderItem',  blank=True, related_name='+')
	contract = models.ForeignKey('Contract', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	priceBook = models.ForeignKey('PriceBook', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.orderDate
		str = str + self.totalAmount
		str = str + self.taxAmount
		str = str + self.shippingAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Order";
    
	def objectType(self):
		return "Order";
