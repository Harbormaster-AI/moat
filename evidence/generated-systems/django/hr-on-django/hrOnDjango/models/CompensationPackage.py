from django.db import models

#======================================================================
# 
# Encapsulates data for model CompensationPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompensationPackage Declaration
#======================================================================
class CompensationPackage (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	effectiveFrom = models.DateField(null=True)
	effectiveTo = models.DateField(null=True)
	currency = models.CharField(max_length=200, null=True)
	contract = models.OneToOneField('EmploymentContract', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	salaryComponents = models.ManyToManyField('SalaryComponent',  blank=True, related_name='+')
	bonusPlans = models.ManyToManyField('BonusPlan',  blank=True, related_name='+')
	equityGrants = models.ManyToManyField('EquityGrant',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.effectiveFrom
		str = str + self.effectiveTo
		str = str + self.currency
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CompensationPackage";
    
	def objectType(self):
		return "CompensationPackage";
