from django.db import models

#======================================================================
# 
# Encapsulates data for model Branch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Branch Declaration
#======================================================================
class Branch (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	branchCode = models.CharField(max_length=200, null=True)
	address = Address
	institution = models.ForeignKey('FinancialInstitution', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.branchCode
		str = str + self.address
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Branch";
    
	def objectType(self):
		return "Branch";
