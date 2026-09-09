from django.db import models

#======================================================================
# 
# Encapsulates data for model AirworthinessDirective
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AirworthinessDirective Declaration
#======================================================================
class AirworthinessDirective (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	directiveNumber = models.CharField(max_length=200, null=True)
	title = models.CharField(max_length=200, null=True)
	workOrders = models.ManyToManyField('MaintenanceWorkOrder',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.directiveNumber
		str = str + self.title
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AirworthinessDirective";
    
	def objectType(self):
		return "AirworthinessDirective";
