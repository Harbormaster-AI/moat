from django.db import models
from aerospaceOnDjango.models.AircraftOrderStatus import AircraftOrderStatus

#======================================================================
# 
# Encapsulates data for model AircraftOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOrder Declaration
#======================================================================
class AircraftOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	totalAmount = Money
	operator = models.ForeignKey('Operator', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variant = models.ForeignKey('AircraftVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	quote = models.OneToOneField('Quote', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	purchaseAgreement = models.OneToOneField('PurchaseAgreement', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AircraftOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.totalAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftOrder";
    
	def objectType(self):
		return "AircraftOrder";
