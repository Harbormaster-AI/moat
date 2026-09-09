from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Incident import Incident
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.InsuredObject import InsuredObject
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Incident
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IncidentDelegate Declaration
#======================================================================
class IncidentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, incidentId ):
		try:	
			incident = Incident.objects.filter(id=incidentId)
			return incident.first();
		except Incident.DoesNotExist:
			raise ProcessingError("Incident with id " + str(incidentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, incident):
		for model in serializers.deserialize("json", incident):
			model.save()
			return model;

	def create(self, incident):
		incident.save()
		return incident;

	def saveFromJson(self, incident):
		for model in serializers.deserialize("json", incident):
			model.save()
			return incident;
	
	def save(self, incident):
		incident.save()
		return incident;
	
	def delete(self, incidentId ):
		errMsg = "Failed to delete Incident from db using id " + str(incidentId)
		
		try:
			incident = Incident.objects.get(id=incidentId)
			incident.delete()
			return True
		except Incident.DoesNotExist:
			raise ProcessingError("Incident with id " + str(incidentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Incident.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Incident from db")
		except Exception:
			return None;
		
	def assignClaim( self, incidentId, claimId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on Incident"

		try:
			# get the Incident from db
			incident = self.get( incidentId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			incident.claim = claim
			
			#save it
			incident.save()

			# reload and return the appropriate version					
			return self.get( incidentId );
		except Incident.DoesNotExist:
			raise ProcessingError(errMsg + " : Incident with id " + str(incidentId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, incidentId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on Incident"

		try:
			# get the Incident from db
			incident = self.get( incidentId ).first()	
			
			# assign to None for unassignment
			incident.claim = None			

			#save it
			incident.save()

			# reload and return the appropriate version					
			return self.get( incidentId );
		except Incident.DoesNotExist:
			raise ProcessingError(errMsg + " : Incident with id " + str(incidentId) + " does not exist.")
		except Exception:
			return None;
		
	def addInsuredObjects( self, incidentId, insuredObjectsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to add elements " + str(insuredObjectsIds) + " for InsuredObjects on Incident"

		try:
			# get the Incident
			incident = self.get( incidentId ).first()
				
			# split on a comma with no spaces
			idList = insuredObjectsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsuredObject		
				insuredObject = InsuredObjectDelegate().get(id).first();	
				# add the InsuredObject
				incident.insuredObjects.add(insuredObject)
				
			# save it		
			incident.save()
			
			# reload and return the appropriate version
			return self.get( incidentId );
		except Incident.DoesNotExist:
			raise ProcessingError(errMsg + " : Incident with id " + str(incidentId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInsuredObjects( self, incidentId, insuredObjectsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to remove elements " + str(insuredObjectsIds) + " for InsuredObjects on Incident"

		try:
			# get the Incident
			incident = self.get( incidentId ).first()
				
			# split on a comma with no spaces
			idList = insuredObjectsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsuredObject		
				insuredObject = InsuredObjectDelegate().get(id).first();	
				# add the InsuredObject
				incident.insuredObjects.remove(insuredObject)
				
			# save it		
			incident.save()
			
			# reload and return the appropriate version
			return self.get( incidentId );
		except Incident.DoesNotExist:
			raise ProcessingError(errMsg + " : Incident with id " + str(incidentId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
