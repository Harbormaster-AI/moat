import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.MaintenanceAppointment import MaintenanceAppointment
from aerospaceOnDjango.delegates.MaintenanceAppointmentDelegate import MaintenanceAppointmentDelegate

 #======================================================================
# 
# Encapsulates data for model MaintenanceAppointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceAppointmentTest Declaration
#======================================================================
class MaintenanceAppointmentTest (TestCase) :
	def test_crud(self) :
		maintenanceAppointment = MaintenanceAppointment()
		maintenanceAppointment.appointmentDate = datetime.datetime.now()
		maintenanceAppointment.status = "default status field value"
		
		delegate = MaintenanceAppointmentDelegate()
		responseObj = delegate.create(maintenanceAppointment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


