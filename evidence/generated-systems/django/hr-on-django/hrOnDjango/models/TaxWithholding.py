from django.db import models
from hrOnDjango.models.FilingStatus import FilingStatus

#======================================================================
# 
# Encapsulates data for model TaxWithholding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxWithholding Declaration
#======================================================================
class TaxWithholding (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	taxId = TaxId
	allowances = models.IntegerField(null=True)
	additionalAmount = Money
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	filingStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FilingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.taxId
		str = str + self.allowances
		str = str + self.additionalAmount
		str = str + self.filingStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TaxWithholding";
    
	def objectType(self):
		return "TaxWithholding";
