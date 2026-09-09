from django.db import models
from insuranceOnDjango.models.ReserveType import ReserveType
from insuranceOnDjango.models.ReserveStatus import ReserveStatus

#======================================================================
# 
# Encapsulates data for model ClaimReserve
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimReserve Declaration
#======================================================================
class ClaimReserve (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	amount = Money
	setDate = models.DateField(null=True)
	claim = models.ForeignKey('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	exposure = models.ForeignKey('Exposure', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reserveType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReserveType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReserveStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.amount
		str = str + self.setDate
		str = str + self.reserveType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ClaimReserve";
    
	def objectType(self):
		return "ClaimReserve";
