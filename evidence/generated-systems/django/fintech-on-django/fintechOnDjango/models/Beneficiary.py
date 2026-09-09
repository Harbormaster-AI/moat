from django.db import models

#======================================================================
# 
# Encapsulates data for model Beneficiary
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Beneficiary Declaration
#======================================================================
class Beneficiary (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	accountIdentifier = AccountIdentifier
	iban = IBAN
	bic = BIC
	address = Address
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.accountIdentifier
		str = str + self.iban
		str = str + self.bic
		str = str + self.address
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Beneficiary";
    
	def objectType(self):
		return "Beneficiary";
