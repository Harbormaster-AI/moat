from django.db import models
from manufacturingOnDjango.models.ForecastMethod import ForecastMethod

#======================================================================
# 
# Encapsulates data for model Forecast
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Forecast Declaration
#======================================================================
class Forecast (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	forecastNumber = models.CharField(max_length=200, null=True)
	forecastHorizonStart = models.DateField(null=True)
	forecastHorizonEnd = models.DateField(null=True)
	lines = models.ManyToManyField('ForecastLine',  blank=True, related_name='+')
	method = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ForecastMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.forecastNumber
		str = str + self.forecastHorizonStart
		str = str + self.forecastHorizonEnd
		str = str + self.method
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Forecast";
    
	def objectType(self):
		return "Forecast";
