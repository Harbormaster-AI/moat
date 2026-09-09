from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.GeoRegion import GeoRegion
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model GeoRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GeoRegionDelegate Declaration
#======================================================================
class GeoRegionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, geoRegionId ):
		try:	
			geoRegion = GeoRegion.objects.filter(id=geoRegionId)
			return geoRegion.first();
		except GeoRegion.DoesNotExist:
			raise ProcessingError("GeoRegion with id " + str(geoRegionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, geoRegion):
		for model in serializers.deserialize("json", geoRegion):
			model.save()
			return model;

	def create(self, geoRegion):
		geoRegion.save()
		return geoRegion;

	def saveFromJson(self, geoRegion):
		for model in serializers.deserialize("json", geoRegion):
			model.save()
			return geoRegion;
	
	def save(self, geoRegion):
		geoRegion.save()
		return geoRegion;
	
	def delete(self, geoRegionId ):
		errMsg = "Failed to delete GeoRegion from db using id " + str(geoRegionId)
		
		try:
			geoRegion = GeoRegion.objects.get(id=geoRegionId)
			geoRegion.delete()
			return True
		except GeoRegion.DoesNotExist:
			raise ProcessingError("GeoRegion with id " + str(geoRegionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = GeoRegion.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all GeoRegion from db")
		except Exception:
			return None;
		
	def assignParent( self, geoRegionId, parentId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

		errMsg = "Failed to assign element " + str(parentId) + " for Parent on GeoRegion"

		try:
			# get the GeoRegion from db
			geoRegion = self.get( geoRegionId ).first()	
			
			# get the GeoRegion from db
			geoRegion = GeoRegionDelegate().get(parentId).first();
			
			# assign the Parent		
			geoRegion.parent = geoRegion
			
			#save it
			geoRegion.save()

			# reload and return the appropriate version					
			return self.get( geoRegionId );
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion with id " + str(geoRegionId) + " does not exist.")
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion with id " + str(parentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParent( self, geoRegionId ):
		errMsg = "Failed to unassign element " + str(parentId) + " for Parent on GeoRegion"

		try:
			# get the GeoRegion from db
			geoRegion = self.get( geoRegionId ).first()	
			
			# assign to None for unassignment
			geoRegion.geoRegion = None			

			#save it
			geoRegion.save()

			# reload and return the appropriate version					
			return self.get( geoRegionId );
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion with id " + str(geoRegionId) + " does not exist.")
		except Exception:
			return None;
		
	def addChildren( self, geoRegionId, childrenIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

		errMsg = "Failed to add elements " + str(childrenIds) + " for Children on GeoRegion"

		try:
			# get the GeoRegion
			geoRegion = self.get( geoRegionId ).first()
				
			# split on a comma with no spaces
			idList = childrenIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GeoRegion		
				geoRegion = GeoRegionDelegate().get(id).first();	
				# add the GeoRegion
				geoRegion.children.add(geoRegion)
				
			# save it		
			geoRegion.save()
			
			# reload and return the appropriate version
			return self.get( geoRegionId );
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion with id " + str(geoRegionId) + " does not exist.")
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChildren( self, geoRegionId, childrenIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

		errMsg = "Failed to remove elements " + str(childrenIds) + " for Children on GeoRegion"

		try:
			# get the GeoRegion
			geoRegion = self.get( geoRegionId ).first()
				
			# split on a comma with no spaces
			idList = childrenIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GeoRegion		
				geoRegion = GeoRegionDelegate().get(id).first();	
				# add the GeoRegion
				geoRegion.children.remove(geoRegion)
				
			# save it		
			geoRegion.save()
			
			# reload and return the appropriate version
			return self.get( geoRegionId );
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion with id " + str(geoRegionId) + " does not exist.")
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
