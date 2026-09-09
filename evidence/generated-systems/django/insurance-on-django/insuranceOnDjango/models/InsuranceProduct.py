from django.db import models
from insuranceOnDjango.models.LineOfBusiness import LineOfBusiness

#======================================================================
# 
# Encapsulates data for model InsuranceProduct
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuranceProduct Declaration
#======================================================================
class InsuranceProduct (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	productCode = models.CharField(max_length=200, null=True)
	insurer = models.ForeignKey('Insurer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	coverageDefinitions = models.ManyToManyField('CoverageDefinition',  blank=True, related_name='+')
	lineOfBusiness = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LineOfBusiness])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.productCode
		str = str + self.lineOfBusiness
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InsuranceProduct";
    
	def objectType(self):
		return "InsuranceProduct";
