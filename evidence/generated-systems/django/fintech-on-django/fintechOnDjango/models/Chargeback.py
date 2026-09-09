from django.db import models
from fintechOnDjango.models.ChargebackStage import ChargebackStage
from fintechOnDjango.models.ChargebackStatus import ChargebackStatus

#======================================================================
# 
# Encapsulates data for model Chargeback
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Chargeback Declaration
#======================================================================
class Chargeback (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	chargebackReference = models.CharField(max_length=200, null=True)
	amount = Money
	postedAt = DateTime
	dispute = models.ForeignKey('Dispute', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transaction = models.ForeignKey('Transaction', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	stage = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ChargebackStage])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ChargebackStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.chargebackReference
		str = str + self.amount
		str = str + self.postedAt
		str = str + self.stage
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Chargeback";
    
	def objectType(self):
		return "Chargeback";
