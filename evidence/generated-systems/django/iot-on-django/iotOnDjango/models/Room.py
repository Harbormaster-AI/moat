from django.db import models

#======================================================================
# 
# Encapsulates data for model Room
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Room Declaration
#======================================================================
class Room (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	floor = models.ForeignKey('Floor', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	devices = models.ManyToManyField('IoTDevice',  blank=True, related_name='+')
	gateways = models.ManyToManyField('Gateway',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Room";
    
	def objectType(self):
		return "Room";
