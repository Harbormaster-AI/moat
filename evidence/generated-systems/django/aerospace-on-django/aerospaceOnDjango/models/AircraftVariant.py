from django.db import models

#======================================================================
# 
# Encapsulates data for model AircraftVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftVariant Declaration
#======================================================================
class AircraftVariant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	variantCode = models.CharField(max_length=200, null=True)
	rangeNm = models.IntegerField(null=True)
	maxTakeoffWeightKg = models.CharField(max_length=64, null=True)
	model = models.ForeignKey('AircraftModel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	engineType = models.ForeignKey('EngineType', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	avionicsSuite = models.ForeignKey('AvionicsSuite', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	apu = models.ForeignKey('APU', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	landingGear = models.ForeignKey('LandingGear', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	cabinLayouts = models.ManyToManyField('CabinLayout',  blank=True, related_name='+')
	options = models.ManyToManyField('AircraftOption',  blank=True, related_name='+')
	packages = models.ManyToManyField('AircraftPackage',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.variantCode
		str = str + self.rangeNm
		str = str + self.maxTakeoffWeightKg
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftVariant";
    
	def objectType(self):
		return "AircraftVariant";
