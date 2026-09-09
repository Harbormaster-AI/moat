from django.db import models
from advertisingOnDjango.models.GeoRegionType import GeoRegionType

#======================================================================
# 
# Encapsulates data for model GeoRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GeoRegion Declaration
#======================================================================
class GeoRegion (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	parent = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	children = models.ManyToManyField('GeoRegion',  blank=True, related_name='+')
	regionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in GeoRegionType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.name
		str = str + self.regionType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "GeoRegion";
    
	def objectType(self):
		return "GeoRegion";
