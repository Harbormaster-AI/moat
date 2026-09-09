from django.db import models
from fintechOnDjango.models.AgreementType import AgreementType
from fintechOnDjango.models.AgreementStatus import AgreementStatus

#======================================================================
# 
# Encapsulates data for model Agreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Agreement Declaration
#======================================================================
class Agreement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	agreementNumber = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	productOffering = models.ForeignKey('ProductOffering', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	agreementType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AgreementType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AgreementStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.agreementNumber
		str = str + self.effectiveDate
		str = str + self.agreementType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Agreement";
    
	def objectType(self):
		return "Agreement";
