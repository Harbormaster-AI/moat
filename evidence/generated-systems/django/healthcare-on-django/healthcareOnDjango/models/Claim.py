from django.db import models
from healthcareOnDjango.models.ClaimStatus import ClaimStatus

#======================================================================
# 
# Encapsulates data for model Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Claim Declaration
#======================================================================
class Claim (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	claimNumber = models.CharField(max_length=200, null=True)
	totalAmount = Money
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	coverage = models.ForeignKey('Coverage', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	invoices = models.ManyToManyField('Invoice',  blank=True, related_name='+')
	payer = models.ForeignKey('InsurancePayer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ClaimStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.claimNumber
		str = str + self.totalAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Claim";
    
	def objectType(self):
		return "Claim";
