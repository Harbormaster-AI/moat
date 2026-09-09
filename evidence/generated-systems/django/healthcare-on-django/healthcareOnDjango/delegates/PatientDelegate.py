from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Appointment import Appointment
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.CarePlan import CarePlan
from healthcareOnDjango.models.Allergy import Allergy
from healthcareOnDjango.models.Condition import Condition
from healthcareOnDjango.models.MedicationOrder import MedicationOrder
from healthcareOnDjango.models.LaboratoryOrder import LaboratoryOrder
from healthcareOnDjango.models.ImagingOrder import ImagingOrder
from healthcareOnDjango.models.Coverage import Coverage
from healthcareOnDjango.models.Claim import Claim
from healthcareOnDjango.models.MedicalDevice import MedicalDevice
from healthcareOnDjango.models.Observation import Observation
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Patient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PatientDelegate Declaration
#======================================================================
class PatientDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, patientId ):
		try:	
			patient = Patient.objects.filter(id=patientId)
			return patient.first();
		except Patient.DoesNotExist:
			raise ProcessingError("Patient with id " + str(patientId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, patient):
		for model in serializers.deserialize("json", patient):
			model.save()
			return model;

	def create(self, patient):
		patient.save()
		return patient;

	def saveFromJson(self, patient):
		for model in serializers.deserialize("json", patient):
			model.save()
			return patient;
	
	def save(self, patient):
		patient.save()
		return patient;
	
	def delete(self, patientId ):
		errMsg = "Failed to delete Patient from db using id " + str(patientId)
		
		try:
			patient = Patient.objects.get(id=patientId)
			patient.delete()
			return True
		except Patient.DoesNotExist:
			raise ProcessingError("Patient with id " + str(patientId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Patient.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Patient from db")
		except Exception:
			return None;
		
	def addAppointments( self, patientId, appointmentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

		errMsg = "Failed to add elements " + str(appointmentsIds) + " for Appointments on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = appointmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Appointment		
				appointment = AppointmentDelegate().get(id).first();	
				# add the Appointment
				patient.appointments.add(appointment)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppointments( self, patientId, appointmentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

		errMsg = "Failed to remove elements " + str(appointmentsIds) + " for Appointments on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = appointmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Appointment		
				appointment = AppointmentDelegate().get(id).first();	
				# add the Appointment
				patient.appointments.remove(appointment)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEncounters( self, patientId, encountersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to add elements " + str(encountersIds) + " for Encounters on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = encountersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Encounter		
				encounter = EncounterDelegate().get(id).first();	
				# add the Encounter
				patient.encounters.add(encounter)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEncounters( self, patientId, encountersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to remove elements " + str(encountersIds) + " for Encounters on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = encountersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Encounter		
				encounter = EncounterDelegate().get(id).first();	
				# add the Encounter
				patient.encounters.remove(encounter)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCarePlans( self, patientId, carePlansIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CarePlanDelegate import CarePlanDelegate

		errMsg = "Failed to add elements " + str(carePlansIds) + " for CarePlans on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = carePlansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CarePlan		
				carePlan = CarePlanDelegate().get(id).first();	
				# add the CarePlan
				patient.carePlans.add(carePlan)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCarePlans( self, patientId, carePlansIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CarePlanDelegate import CarePlanDelegate

		errMsg = "Failed to remove elements " + str(carePlansIds) + " for CarePlans on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = carePlansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CarePlan		
				carePlan = CarePlanDelegate().get(id).first();	
				# add the CarePlan
				patient.carePlans.remove(carePlan)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAllergies( self, patientId, allergiesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AllergyDelegate import AllergyDelegate

		errMsg = "Failed to add elements " + str(allergiesIds) + " for Allergies on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = allergiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Allergy		
				allergy = AllergyDelegate().get(id).first();	
				# add the Allergy
				patient.allergies.add(allergy)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Allergy.DoesNotExist:
			raise ProcessingError(errMsg + " : Allergy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAllergies( self, patientId, allergiesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AllergyDelegate import AllergyDelegate

		errMsg = "Failed to remove elements " + str(allergiesIds) + " for Allergies on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = allergiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Allergy		
				allergy = AllergyDelegate().get(id).first();	
				# add the Allergy
				patient.allergies.remove(allergy)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Allergy.DoesNotExist:
			raise ProcessingError(errMsg + " : Allergy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addConditions( self, patientId, conditionsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ConditionDelegate import ConditionDelegate

		errMsg = "Failed to add elements " + str(conditionsIds) + " for Conditions on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = conditionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Condition		
				condition = ConditionDelegate().get(id).first();	
				# add the Condition
				patient.conditions.add(condition)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Condition.DoesNotExist:
			raise ProcessingError(errMsg + " : Condition does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeConditions( self, patientId, conditionsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ConditionDelegate import ConditionDelegate

		errMsg = "Failed to remove elements " + str(conditionsIds) + " for Conditions on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = conditionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Condition		
				condition = ConditionDelegate().get(id).first();	
				# add the Condition
				patient.conditions.remove(condition)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Condition.DoesNotExist:
			raise ProcessingError(errMsg + " : Condition does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMedicationOrders( self, patientId, medicationOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to add elements " + str(medicationOrdersIds) + " for MedicationOrders on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = medicationOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicationOrder		
				medicationOrder = MedicationOrderDelegate().get(id).first();	
				# add the MedicationOrder
				patient.medicationOrders.add(medicationOrder)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMedicationOrders( self, patientId, medicationOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to remove elements " + str(medicationOrdersIds) + " for MedicationOrders on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = medicationOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicationOrder		
				medicationOrder = MedicationOrderDelegate().get(id).first();	
				# add the MedicationOrder
				patient.medicationOrders.remove(medicationOrder)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLabOrders( self, patientId, labOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to add elements " + str(labOrdersIds) + " for LabOrders on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = labOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LaboratoryOrder		
				laboratoryOrder = LaboratoryOrderDelegate().get(id).first();	
				# add the LaboratoryOrder
				patient.labOrders.add(laboratoryOrder)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLabOrders( self, patientId, labOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to remove elements " + str(labOrdersIds) + " for LabOrders on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = labOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LaboratoryOrder		
				laboratoryOrder = LaboratoryOrderDelegate().get(id).first();	
				# add the LaboratoryOrder
				patient.labOrders.remove(laboratoryOrder)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addImagingOrders( self, patientId, imagingOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to add elements " + str(imagingOrdersIds) + " for ImagingOrders on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = imagingOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingOrder		
				imagingOrder = ImagingOrderDelegate().get(id).first();	
				# add the ImagingOrder
				patient.imagingOrders.add(imagingOrder)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeImagingOrders( self, patientId, imagingOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to remove elements " + str(imagingOrdersIds) + " for ImagingOrders on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = imagingOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingOrder		
				imagingOrder = ImagingOrderDelegate().get(id).first();	
				# add the ImagingOrder
				patient.imagingOrders.remove(imagingOrder)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCoverages( self, patientId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

		errMsg = "Failed to add elements " + str(coveragesIds) + " for Coverages on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Coverage		
				coverage = CoverageDelegate().get(id).first();	
				# add the Coverage
				patient.coverages.add(coverage)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCoverages( self, patientId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

		errMsg = "Failed to remove elements " + str(coveragesIds) + " for Coverages on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Coverage		
				coverage = CoverageDelegate().get(id).first();	
				# add the Coverage
				patient.coverages.remove(coverage)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addClaims( self, patientId, claimsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				patient.claims.add(claim)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, patientId, claimsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				patient.claims.remove(claim)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDevices( self, patientId, devicesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalDeviceDelegate import MedicalDeviceDelegate

		errMsg = "Failed to add elements " + str(devicesIds) + " for Devices on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = devicesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicalDevice		
				medicalDevice = MedicalDeviceDelegate().get(id).first();	
				# add the MedicalDevice
				patient.devices.add(medicalDevice)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDevices( self, patientId, devicesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalDeviceDelegate import MedicalDeviceDelegate

		errMsg = "Failed to remove elements " + str(devicesIds) + " for Devices on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = devicesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicalDevice		
				medicalDevice = MedicalDeviceDelegate().get(id).first();	
				# add the MedicalDevice
				patient.devices.remove(medicalDevice)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addObservations( self, patientId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to add elements " + str(observationsIds) + " for Observations on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				patient.observations.add(observation)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObservations( self, patientId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to remove elements " + str(observationsIds) + " for Observations on Patient"

		try:
			# get the Patient
			patient = self.get( patientId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				patient.observations.remove(observation)
				
			# save it		
			patient.save()
			
			# reload and return the appropriate version
			return self.get( patientId );
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
