from django.db import models

#======================================================================
# 
# Encapsulates data for model RateCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateCard Declaration
#======================================================================
class RateCard (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	currency = models.CharField(max_length=200, null=True)
	publisher = models.ForeignKey('Publisher', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	rates = models.ManyToManyField('Rate',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.effectiveDate
		str = str + self.currency
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RateCard";
    
	def objectType(self):
		return "RateCard";
