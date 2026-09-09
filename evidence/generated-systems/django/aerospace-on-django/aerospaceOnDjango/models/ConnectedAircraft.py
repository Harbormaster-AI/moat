from django.db import models
from aerospaceOnDjango.models.ConnectivityStatus import ConnectivityStatus

#======================================================================
# 
# Encapsulates data for model ConnectedAircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConnectedAircraft Declaration
#======================================================================
class ConnectedAircraft (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	communicationsProvider = models.CharField(max_length=200, null=True)
	aircraft = models.OneToOneField('Aircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	flightHealthEvents = models.ManyToManyField('FlightHealthEvent',  blank=True, related_name='+')
	softwareLoads = models.ManyToManyField('SoftwareLoad',  blank=True, related_name='+')
	connectivityStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConnectivityStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.communicationsProvider
		str = str + self.connectivityStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ConnectedAircraft";
    
	def objectType(self):
		return "ConnectedAircraft";
