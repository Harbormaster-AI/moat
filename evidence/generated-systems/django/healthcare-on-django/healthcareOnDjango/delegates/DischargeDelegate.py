from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Discharge import Discharge
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Discharge
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DischargeDelegate Declaration
#======================================================================
class DischargeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dischargeId ):
		try:	
			discharge = Discharge.objects.filter(id=dischargeId)
			return discharge.first();
		except Discharge.DoesNotExist:
			raise ProcessingError("Discharge with id " + str(dischargeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, discharge):
		for model in serializers.deserialize("json", discharge):
			model.save()
			return model;

	def create(self, discharge):
		discharge.save()
		return discharge;

	def saveFromJson(self, discharge):
		for model in serializers.deserialize("json", discharge):
			model.save()
			return discharge;
	
	def save(self, discharge):
		discharge.save()
		return discharge;
	
	def delete(self, dischargeId ):
		errMsg = "Failed to delete Discharge from db using id " + str(dischargeId)
		
		try:
			discharge = Discharge.objects.get(id=dischargeId)
			discharge.delete()
			return True
		except Discharge.DoesNotExist:
			raise ProcessingError("Discharge with id " + str(dischargeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Discharge.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Discharge from db")
		except Exception:
			return None;
		
	def assignEncounter( self, dischargeId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Discharge"

		try:
			# get the Discharge from db
			discharge = self.get( dischargeId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			discharge.encounter = encounter
			
			#save it
			discharge.save()

			# reload and return the appropriate version					
			return self.get( dischargeId );
		except Discharge.DoesNotExist:
			raise ProcessingError(errMsg + " : Discharge with id " + str(dischargeId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, dischargeId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Discharge"

		try:
			# get the Discharge from db
			discharge = self.get( dischargeId ).first()	
			
			# assign to None for unassignment
			discharge.encounter = None			

			#save it
			discharge.save()

			# reload and return the appropriate version					
			return self.get( dischargeId );
		except Discharge.DoesNotExist:
			raise ProcessingError(errMsg + " : Discharge with id " + str(dischargeId) + " does not exist.")
		except Exception:
			return None;
		
