from django.db import models

#======================================================================
# 
# Encapsulates data for model PurchaseAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseAgreement Declaration
#======================================================================
class PurchaseAgreement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	agreementNumber = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	aircraftOrder = models.OneToOneField('AircraftOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.agreementNumber
		str = str + self.effectiveDate
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PurchaseAgreement";
    
	def objectType(self):
		return "PurchaseAgreement";
