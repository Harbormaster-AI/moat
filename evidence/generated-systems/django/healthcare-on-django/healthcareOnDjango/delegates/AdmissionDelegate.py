from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Admission import Admission
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Admission
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdmissionDelegate Declaration
#======================================================================
class AdmissionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, admissionId ):
		try:	
			admission = Admission.objects.filter(id=admissionId)
			return admission.first();
		except Admission.DoesNotExist:
			raise ProcessingError("Admission with id " + str(admissionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, admission):
		for model in serializers.deserialize("json", admission):
			model.save()
			return model;

	def create(self, admission):
		admission.save()
		return admission;

	def saveFromJson(self, admission):
		for model in serializers.deserialize("json", admission):
			model.save()
			return admission;
	
	def save(self, admission):
		admission.save()
		return admission;
	
	def delete(self, admissionId ):
		errMsg = "Failed to delete Admission from db using id " + str(admissionId)
		
		try:
			admission = Admission.objects.get(id=admissionId)
			admission.delete()
			return True
		except Admission.DoesNotExist:
			raise ProcessingError("Admission with id " + str(admissionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Admission.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Admission from db")
		except Exception:
			return None;
		
	def assignEncounter( self, admissionId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Admission"

		try:
			# get the Admission from db
			admission = self.get( admissionId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			admission.encounter = encounter
			
			#save it
			admission.save()

			# reload and return the appropriate version					
			return self.get( admissionId );
		except Admission.DoesNotExist:
			raise ProcessingError(errMsg + " : Admission with id " + str(admissionId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, admissionId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Admission"

		try:
			# get the Admission from db
			admission = self.get( admissionId ).first()	
			
			# assign to None for unassignment
			admission.encounter = None			

			#save it
			admission.save()

			# reload and return the appropriate version					
			return self.get( admissionId );
		except Admission.DoesNotExist:
			raise ProcessingError(errMsg + " : Admission with id " + str(admissionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFacility( self, admissionId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on Admission"

		try:
			# get the Admission from db
			admission = self.get( admissionId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			admission.facility = facility
			
			#save it
			admission.save()

			# reload and return the appropriate version					
			return self.get( admissionId );
		except Admission.DoesNotExist:
			raise ProcessingError(errMsg + " : Admission with id " + str(admissionId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, admissionId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on Admission"

		try:
			# get the Admission from db
			admission = self.get( admissionId ).first()	
			
			# assign to None for unassignment
			admission.facility = None			

			#save it
			admission.save()

			# reload and return the appropriate version					
			return self.get( admissionId );
		except Admission.DoesNotExist:
			raise ProcessingError(errMsg + " : Admission with id " + str(admissionId) + " does not exist.")
		except Exception:
			return None;
		
