from django.db import models
from ecommerceOnDjango.models.SubscriptionStatus import SubscriptionStatus
from ecommerceOnDjango.models.SubscriptionInterval import SubscriptionInterval

#======================================================================
# 
# Encapsulates data for model Subscription
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Subscription Declaration
#======================================================================
class Subscription (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	subscriptionNumber = models.CharField(max_length=200, null=True)
	nextBillingDate = models.DateField(null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	paymentProvider = models.ForeignKey('PaymentProvider', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channel = models.ForeignKey('Channel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SubscriptionStatus])
	interval = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SubscriptionInterval])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.subscriptionNumber
		str = str + self.nextBillingDate
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.status
		str = str + self.interval
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Subscription";
    
	def objectType(self):
		return "Subscription";
