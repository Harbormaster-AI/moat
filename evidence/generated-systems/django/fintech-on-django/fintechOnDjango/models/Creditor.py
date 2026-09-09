from django.db import models

#======================================================================
# 
# Encapsulates data for model Creditor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Creditor Declaration
#======================================================================
class Creditor (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	bic = BIC
	address = Address
	mandates = models.ManyToManyField('DirectDebitMandate',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.bic
		str = str + self.address
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Creditor";
    
	def objectType(self):
		return "Creditor";
