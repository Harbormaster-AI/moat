from django.db import models
from healthcareOnDjango.models.AdministrativeSex import AdministrativeSex
from healthcareOnDjango.models.BloodType import BloodType

#======================================================================
# 
# Encapsulates data for model Patient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Patient Declaration
#======================================================================
class Patient (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	mrn = MRN
	dateOfBirth = models.DateField(null=True)
	address = Address
	primaryLanguage = models.CharField(max_length=200, null=True)
	appointments = models.ManyToManyField('Appointment',  blank=True, related_name='+')
	encounters = models.ManyToManyField('Encounter',  blank=True, related_name='+')
	carePlans = models.ManyToManyField('CarePlan',  blank=True, related_name='+')
	allergies = models.ManyToManyField('Allergy',  blank=True, related_name='+')
	conditions = models.ManyToManyField('Condition',  blank=True, related_name='+')
	medicationOrders = models.ManyToManyField('MedicationOrder',  blank=True, related_name='+')
	labOrders = models.ManyToManyField('LaboratoryOrder',  blank=True, related_name='+')
	imagingOrders = models.ManyToManyField('ImagingOrder',  blank=True, related_name='+')
	coverages = models.ManyToManyField('Coverage',  blank=True, related_name='+')
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	devices = models.ManyToManyField('MedicalDevice',  blank=True, related_name='+')
	observations = models.ManyToManyField('Observation',  blank=True, related_name='+')
	sexAtBirth = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdministrativeSex])
	bloodType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BloodType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.mrn
		str = str + self.dateOfBirth
		str = str + self.address
		str = str + self.primaryLanguage
		str = str + self.sexAtBirth
		str = str + self.bloodType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Patient";
    
	def objectType(self):
		return "Patient";
