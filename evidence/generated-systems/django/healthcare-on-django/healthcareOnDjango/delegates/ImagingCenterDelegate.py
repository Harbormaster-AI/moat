from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.ImagingCenter import ImagingCenter
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.ImagingOrder import ImagingOrder
from healthcareOnDjango.models.ImagingReport import ImagingReport
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ImagingCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingCenterDelegate Declaration
#======================================================================
class ImagingCenterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, imagingCenterId ):
		try:	
			imagingCenter = ImagingCenter.objects.filter(id=imagingCenterId)
			return imagingCenter.first();
		except ImagingCenter.DoesNotExist:
			raise ProcessingError("ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, imagingCenter):
		for model in serializers.deserialize("json", imagingCenter):
			model.save()
			return model;

	def create(self, imagingCenter):
		imagingCenter.save()
		return imagingCenter;

	def saveFromJson(self, imagingCenter):
		for model in serializers.deserialize("json", imagingCenter):
			model.save()
			return imagingCenter;
	
	def save(self, imagingCenter):
		imagingCenter.save()
		return imagingCenter;
	
	def delete(self, imagingCenterId ):
		errMsg = "Failed to delete ImagingCenter from db using id " + str(imagingCenterId)
		
		try:
			imagingCenter = ImagingCenter.objects.get(id=imagingCenterId)
			imagingCenter.delete()
			return True
		except ImagingCenter.DoesNotExist:
			raise ProcessingError("ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ImagingCenter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ImagingCenter from db")
		except Exception:
			return None;
		
	def assignFacility( self, imagingCenterId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on ImagingCenter"

		try:
			# get the ImagingCenter from db
			imagingCenter = self.get( imagingCenterId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			imagingCenter.facility = facility
			
			#save it
			imagingCenter.save()

			# reload and return the appropriate version					
			return self.get( imagingCenterId );
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, imagingCenterId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on ImagingCenter"

		try:
			# get the ImagingCenter from db
			imagingCenter = self.get( imagingCenterId ).first()	
			
			# assign to None for unassignment
			imagingCenter.facility = None			

			#save it
			imagingCenter.save()

			# reload and return the appropriate version					
			return self.get( imagingCenterId );
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except Exception:
			return None;
		
	def addImagingOrders( self, imagingCenterId, imagingOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to add elements " + str(imagingOrdersIds) + " for ImagingOrders on ImagingCenter"

		try:
			# get the ImagingCenter
			imagingCenter = self.get( imagingCenterId ).first()
				
			# split on a comma with no spaces
			idList = imagingOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingOrder		
				imagingOrder = ImagingOrderDelegate().get(id).first();	
				# add the ImagingOrder
				imagingCenter.imagingOrders.add(imagingOrder)
				
			# save it		
			imagingCenter.save()
			
			# reload and return the appropriate version
			return self.get( imagingCenterId );
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeImagingOrders( self, imagingCenterId, imagingOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to remove elements " + str(imagingOrdersIds) + " for ImagingOrders on ImagingCenter"

		try:
			# get the ImagingCenter
			imagingCenter = self.get( imagingCenterId ).first()
				
			# split on a comma with no spaces
			idList = imagingOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingOrder		
				imagingOrder = ImagingOrderDelegate().get(id).first();	
				# add the ImagingOrder
				imagingCenter.imagingOrders.remove(imagingOrder)
				
			# save it		
			imagingCenter.save()
			
			# reload and return the appropriate version
			return self.get( imagingCenterId );
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addImagingReports( self, imagingCenterId, imagingReportsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

		errMsg = "Failed to add elements " + str(imagingReportsIds) + " for ImagingReports on ImagingCenter"

		try:
			# get the ImagingCenter
			imagingCenter = self.get( imagingCenterId ).first()
				
			# split on a comma with no spaces
			idList = imagingReportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingReport		
				imagingReport = ImagingReportDelegate().get(id).first();	
				# add the ImagingReport
				imagingCenter.imagingReports.add(imagingReport)
				
			# save it		
			imagingCenter.save()
			
			# reload and return the appropriate version
			return self.get( imagingCenterId );
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeImagingReports( self, imagingCenterId, imagingReportsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

		errMsg = "Failed to remove elements " + str(imagingReportsIds) + " for ImagingReports on ImagingCenter"

		try:
			# get the ImagingCenter
			imagingCenter = self.get( imagingCenterId ).first()
				
			# split on a comma with no spaces
			idList = imagingReportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingReport		
				imagingReport = ImagingReportDelegate().get(id).first();	
				# add the ImagingReport
				imagingCenter.imagingReports.remove(imagingReport)
				
			# save it		
			imagingCenter.save()
			
			# reload and return the appropriate version
			return self.get( imagingCenterId );
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
