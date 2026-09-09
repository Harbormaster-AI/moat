from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Appointment import Appointment
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Appointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppointmentDelegate Declaration
#======================================================================
class AppointmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, appointmentId ):
		try:	
			appointment = Appointment.objects.filter(id=appointmentId)
			return appointment.first();
		except Appointment.DoesNotExist:
			raise ProcessingError("Appointment with id " + str(appointmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, appointment):
		for model in serializers.deserialize("json", appointment):
			model.save()
			return model;

	def create(self, appointment):
		appointment.save()
		return appointment;

	def saveFromJson(self, appointment):
		for model in serializers.deserialize("json", appointment):
			model.save()
			return appointment;
	
	def save(self, appointment):
		appointment.save()
		return appointment;
	
	def delete(self, appointmentId ):
		errMsg = "Failed to delete Appointment from db using id " + str(appointmentId)
		
		try:
			appointment = Appointment.objects.get(id=appointmentId)
			appointment.delete()
			return True
		except Appointment.DoesNotExist:
			raise ProcessingError("Appointment with id " + str(appointmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Appointment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Appointment from db")
		except Exception:
			return None;
		
	def assignPatient( self, appointmentId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			appointment.patient = patient
			
			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, appointmentId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# assign to None for unassignment
			appointment.patient = None			

			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignClinician( self, appointmentId, clinicianId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to assign element " + str(clinicianId) + " for Clinician on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# get the Clinician from db
			clinician = ClinicianDelegate().get(clinicianId).first();
			
			# assign the Clinician		
			appointment.clinician = clinician
			
			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClinician( self, appointmentId ):
		errMsg = "Failed to unassign element " + str(clinicianId) + " for Clinician on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# assign to None for unassignment
			appointment.clinician = None			

			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFacility( self, appointmentId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			appointment.facility = facility
			
			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, appointmentId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# assign to None for unassignment
			appointment.facility = None			

			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEncounter( self, appointmentId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			appointment.encounter = encounter
			
			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, appointmentId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Appointment"

		try:
			# get the Appointment from db
			appointment = self.get( appointmentId ).first()	
			
			# assign to None for unassignment
			appointment.encounter = None			

			#save it
			appointment.save()

			# reload and return the appropriate version					
			return self.get( appointmentId );
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Exception:
			return None;
		
