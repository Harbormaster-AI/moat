from django.db import models
from advertisingOnDjango.models.BrandSafetyLevel import BrandSafetyLevel
from advertisingOnDjango.models.ContentRating import ContentRating

#======================================================================
# 
# Encapsulates data for model BrandSafetyPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandSafetyPolicy Declaration
#======================================================================
class BrandSafetyPolicy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	targetingProfiles = models.ManyToManyField('TargetingProfile',  blank=True, related_name='+')
	level = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BrandSafetyLevel])
	contentRatingThreshold = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ContentRating])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.level
		str = str + self.contentRatingThreshold
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BrandSafetyPolicy";
    
	def objectType(self):
		return "BrandSafetyPolicy";
