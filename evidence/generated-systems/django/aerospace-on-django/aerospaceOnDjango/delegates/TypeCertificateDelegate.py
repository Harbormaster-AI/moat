from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.TypeCertificate import TypeCertificate
from aerospaceOnDjango.models.AircraftProgram import AircraftProgram
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TypeCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TypeCertificateDelegate Declaration
#======================================================================
class TypeCertificateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, typeCertificateId ):
		try:	
			typeCertificate = TypeCertificate.objects.filter(id=typeCertificateId)
			return typeCertificate.first();
		except TypeCertificate.DoesNotExist:
			raise ProcessingError("TypeCertificate with id " + str(typeCertificateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, typeCertificate):
		for model in serializers.deserialize("json", typeCertificate):
			model.save()
			return model;

	def create(self, typeCertificate):
		typeCertificate.save()
		return typeCertificate;

	def saveFromJson(self, typeCertificate):
		for model in serializers.deserialize("json", typeCertificate):
			model.save()
			return typeCertificate;
	
	def save(self, typeCertificate):
		typeCertificate.save()
		return typeCertificate;
	
	def delete(self, typeCertificateId ):
		errMsg = "Failed to delete TypeCertificate from db using id " + str(typeCertificateId)
		
		try:
			typeCertificate = TypeCertificate.objects.get(id=typeCertificateId)
			typeCertificate.delete()
			return True
		except TypeCertificate.DoesNotExist:
			raise ProcessingError("TypeCertificate with id " + str(typeCertificateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TypeCertificate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TypeCertificate from db")
		except Exception:
			return None;
		
	def assignProgram( self, typeCertificateId, programId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftProgramDelegate import AircraftProgramDelegate

		errMsg = "Failed to assign element " + str(programId) + " for Program on TypeCertificate"

		try:
			# get the TypeCertificate from db
			typeCertificate = self.get( typeCertificateId ).first()	
			
			# get the AircraftProgram from db
			aircraftProgram = AircraftProgramDelegate().get(programId).first();
			
			# assign the Program		
			typeCertificate.program = aircraftProgram
			
			#save it
			typeCertificate.save()

			# reload and return the appropriate version					
			return self.get( typeCertificateId );
		except TypeCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : TypeCertificate with id " + str(typeCertificateId) + " does not exist.")
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(programId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProgram( self, typeCertificateId ):
		errMsg = "Failed to unassign element " + str(programId) + " for Program on TypeCertificate"

		try:
			# get the TypeCertificate from db
			typeCertificate = self.get( typeCertificateId ).first()	
			
			# assign to None for unassignment
			typeCertificate.aircraftProgram = None			

			#save it
			typeCertificate.save()

			# reload and return the appropriate version					
			return self.get( typeCertificateId );
		except TypeCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : TypeCertificate with id " + str(typeCertificateId) + " does not exist.")
		except Exception:
			return None;
		
