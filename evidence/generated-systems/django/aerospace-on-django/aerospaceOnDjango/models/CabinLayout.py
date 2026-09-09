from django.db import models

#======================================================================
# 
# Encapsulates data for model CabinLayout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CabinLayout Declaration
#======================================================================
class CabinLayout (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	layoutCode = models.CharField(max_length=200, null=True)
	totalSeats = models.IntegerField(null=True)
	classLayout = models.CharField(max_length=200, null=True)
	variant = models.ForeignKey('AircraftVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	aircraft = models.ManyToManyField('Aircraft',  blank=True, related_name='+')
	options = models.ManyToManyField('AircraftOption',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.layoutCode
		str = str + self.totalSeats
		str = str + self.classLayout
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CabinLayout";
    
	def objectType(self):
		return "CabinLayout";
