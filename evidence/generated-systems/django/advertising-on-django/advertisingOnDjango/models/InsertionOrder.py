from django.db import models
from advertisingOnDjango.models.IOStatus import IOStatus

#======================================================================
# 
# Encapsulates data for model InsertionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsertionOrder Declaration
#======================================================================
class InsertionOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	ioNumber = models.CharField(max_length=200, null=True)
	agreedBudget = Money
	flight = DateRange
	advertiser = models.ForeignKey('Advertiser', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	agency = models.ForeignKey('Agency', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	publisher = models.ForeignKey('Publisher', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in IOStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.ioNumber
		str = str + self.agreedBudget
		str = str + self.flight
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InsertionOrder";
    
	def objectType(self):
		return "InsertionOrder";
