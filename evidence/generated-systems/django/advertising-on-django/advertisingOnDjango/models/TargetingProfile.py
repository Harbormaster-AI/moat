from django.db import models

#======================================================================
# 
# Encapsulates data for model TargetingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TargetingProfile Declaration
#======================================================================
class TargetingProfile (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	audienceSegments = models.ManyToManyField('AudienceSegment',  blank=True, related_name='+')
	geoRegions = models.ManyToManyField('GeoRegion',  blank=True, related_name='+')
	contentCategories = models.ManyToManyField('ContentCategory',  blank=True, related_name='+')
	brandSafetyPolicy = models.ForeignKey('BrandSafetyPolicy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	deviceCriteria = models.ManyToManyField('DeviceCriterion',  blank=True, related_name='+')

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
		return "TargetingProfile";
    
	def objectType(self):
		return "TargetingProfile";
