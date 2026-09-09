from django.db import models
from fintechOnDjango.models.SettlementStatus import SettlementStatus

#======================================================================
# 
# Encapsulates data for model SettlementBatch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SettlementBatch Declaration
#======================================================================
class SettlementBatch (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	batchId = models.CharField(max_length=200, null=True)
	periodStart = DateTime
	periodEnd = DateTime
	totalVolume = Money
	totalCount = models.IntegerField(null=True)
	processor = models.ForeignKey('PaymentProcessor', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payouts = models.ManyToManyField('Payout',  blank=True, related_name='+')
	transactions = models.ManyToManyField('Transaction',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SettlementStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.batchId
		str = str + self.periodStart
		str = str + self.periodEnd
		str = str + self.totalVolume
		str = str + self.totalCount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SettlementBatch";
    
	def objectType(self):
		return "SettlementBatch";
