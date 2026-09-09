from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.HealthSystem import HealthSystem
from healthcareOnDjango.models.Department import Department
from healthcareOnDjango.models.CareTeam import CareTeam
from healthcareOnDjango.models.Laboratory import Laboratory
from healthcareOnDjango.models.ImagingCenter import ImagingCenter
from healthcareOnDjango.models.Pharmacy import Pharmacy
from healthcareOnDjango.models.InventoryItem import InventoryItem
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Facility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FacilityDelegate Declaration
#======================================================================
class FacilityDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, facilityId ):
		try:	
			facility = Facility.objects.filter(id=facilityId)
			return facility.first();
		except Facility.DoesNotExist:
			raise ProcessingError("Facility with id " + str(facilityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, facility):
		for model in serializers.deserialize("json", facility):
			model.save()
			return model;

	def create(self, facility):
		facility.save()
		return facility;

	def saveFromJson(self, facility):
		for model in serializers.deserialize("json", facility):
			model.save()
			return facility;
	
	def save(self, facility):
		facility.save()
		return facility;
	
	def delete(self, facilityId ):
		errMsg = "Failed to delete Facility from db using id " + str(facilityId)
		
		try:
			facility = Facility.objects.get(id=facilityId)
			facility.delete()
			return True
		except Facility.DoesNotExist:
			raise ProcessingError("Facility with id " + str(facilityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Facility.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Facility from db")
		except Exception:
			return None;
		
	def assignHealthSystem( self, facilityId, healthSystemId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.HealthSystemDelegate import HealthSystemDelegate

		errMsg = "Failed to assign element " + str(healthSystemId) + " for HealthSystem on Facility"

		try:
			# get the Facility from db
			facility = self.get( facilityId ).first()	
			
			# get the HealthSystem from db
			healthSystem = HealthSystemDelegate().get(healthSystemId).first();
			
			# assign the HealthSystem		
			facility.healthSystem = healthSystem
			
			#save it
			facility.save()

			# reload and return the appropriate version					
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except HealthSystem.DoesNotExist:
			raise ProcessingError(errMsg + " : HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignHealthSystem( self, facilityId ):
		errMsg = "Failed to unassign element " + str(healthSystemId) + " for HealthSystem on Facility"

		try:
			# get the Facility from db
			facility = self.get( facilityId ).first()	
			
			# assign to None for unassignment
			facility.healthSystem = None			

			#save it
			facility.save()

			# reload and return the appropriate version					
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
		
	def addDepartments( self, facilityId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to add elements " + str(departmentsIds) + " for Departments on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				facility.departments.add(department)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDepartments( self, facilityId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to remove elements " + str(departmentsIds) + " for Departments on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				facility.departments.remove(department)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCareTeams( self, facilityId, careTeamsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to add elements " + str(careTeamsIds) + " for CareTeams on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = careTeamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CareTeam		
				careTeam = CareTeamDelegate().get(id).first();	
				# add the CareTeam
				facility.careTeams.add(careTeam)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCareTeams( self, facilityId, careTeamsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to remove elements " + str(careTeamsIds) + " for CareTeams on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = careTeamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CareTeam		
				careTeam = CareTeamDelegate().get(id).first();	
				# add the CareTeam
				facility.careTeams.remove(careTeam)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLaboratories( self, facilityId, laboratoriesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryDelegate import LaboratoryDelegate

		errMsg = "Failed to add elements " + str(laboratoriesIds) + " for Laboratories on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = laboratoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Laboratory		
				laboratory = LaboratoryDelegate().get(id).first();	
				# add the Laboratory
				facility.laboratories.add(laboratory)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLaboratories( self, facilityId, laboratoriesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryDelegate import LaboratoryDelegate

		errMsg = "Failed to remove elements " + str(laboratoriesIds) + " for Laboratories on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = laboratoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Laboratory		
				laboratory = LaboratoryDelegate().get(id).first();	
				# add the Laboratory
				facility.laboratories.remove(laboratory)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addImagingCenters( self, facilityId, imagingCentersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingCenterDelegate import ImagingCenterDelegate

		errMsg = "Failed to add elements " + str(imagingCentersIds) + " for ImagingCenters on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = imagingCentersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingCenter		
				imagingCenter = ImagingCenterDelegate().get(id).first();	
				# add the ImagingCenter
				facility.imagingCenters.add(imagingCenter)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeImagingCenters( self, facilityId, imagingCentersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingCenterDelegate import ImagingCenterDelegate

		errMsg = "Failed to remove elements " + str(imagingCentersIds) + " for ImagingCenters on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = imagingCentersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingCenter		
				imagingCenter = ImagingCenterDelegate().get(id).first();	
				# add the ImagingCenter
				facility.imagingCenters.remove(imagingCenter)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPharmacies( self, facilityId, pharmaciesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PharmacyDelegate import PharmacyDelegate

		errMsg = "Failed to add elements " + str(pharmaciesIds) + " for Pharmacies on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = pharmaciesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Pharmacy		
				pharmacy = PharmacyDelegate().get(id).first();	
				# add the Pharmacy
				facility.pharmacies.add(pharmacy)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePharmacies( self, facilityId, pharmaciesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PharmacyDelegate import PharmacyDelegate

		errMsg = "Failed to remove elements " + str(pharmaciesIds) + " for Pharmacies on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = pharmaciesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Pharmacy		
				pharmacy = PharmacyDelegate().get(id).first();	
				# add the Pharmacy
				facility.pharmacies.remove(pharmacy)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, facilityId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				facility.inventoryItems.add(inventoryItem)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, facilityId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on Facility"

		try:
			# get the Facility
			facility = self.get( facilityId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				facility.inventoryItems.remove(inventoryItem)
				
			# save it		
			facility.save()
			
			# reload and return the appropriate version
			return self.get( facilityId );
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
