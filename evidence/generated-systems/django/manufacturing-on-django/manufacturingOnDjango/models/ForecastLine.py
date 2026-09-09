from django.db import models

#======================================================================
# 
# Encapsulates data for model ForecastLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastLine Declaration
#======================================================================
class ForecastLine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	period = models.DateField(null=True)
	quantity = Quantity
	confidence = Percentage
	forecast = models.ForeignKey('Forecast', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.period
		str = str + self.quantity
		str = str + self.confidence
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ForecastLine";
    
	def objectType(self):
		return "ForecastLine";
