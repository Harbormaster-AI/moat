from django.db import models
from fintechOnDjango.models.DisputeReason import DisputeReason
from fintechOnDjango.models.DisputeStatus import DisputeStatus

#======================================================================
# 
# Encapsulates data for model Dispute
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Dispute Declaration
#======================================================================
class Dispute (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	disputeReference = models.CharField(max_length=200, null=True)
	openedAt = DateTime
	closedAt = DateTime
	transaction = models.ForeignKey('Transaction', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	card = models.ForeignKey('PaymentCard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	chargebacks = models.ManyToManyField('Chargeback',  blank=True, related_name='+')
	reason = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DisputeReason])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DisputeStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.disputeReference
		str = str + self.openedAt
		str = str + self.closedAt
		str = str + self.reason
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Dispute";
    
	def objectType(self):
		return "Dispute";
