from django.db import models

#======================================================================
# 
# Encapsulates data for model Placement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Placement Declaration
#======================================================================
class Placement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	flight = DateRange
	goalImpressions = models.IntegerField(null=True)
	lineItem = models.ForeignKey('LineItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adSlot = models.ForeignKey('AdSlot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	deal = models.ForeignKey('Deal', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.flight
		str = str + self.goalImpressions
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Placement";
    
	def objectType(self):
		return "Placement";
