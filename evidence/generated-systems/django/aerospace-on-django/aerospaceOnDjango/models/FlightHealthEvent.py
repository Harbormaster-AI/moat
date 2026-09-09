from django.db import models
from aerospaceOnDjango.models.EventSeverity import EventSeverity

#======================================================================
# 
# Encapsulates data for model FlightHealthEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FlightHealthEvent Declaration
#======================================================================
class FlightHealthEvent (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	eventCode = models.CharField(max_length=200, null=True)
	connectedAircraft = models.ForeignKey('ConnectedAircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EventSeverity])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.eventCode
		str = str + self.severity
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FlightHealthEvent";
    
	def objectType(self):
		return "FlightHealthEvent";
