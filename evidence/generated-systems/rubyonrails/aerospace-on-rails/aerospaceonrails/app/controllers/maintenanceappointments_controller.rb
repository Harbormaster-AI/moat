class MaintenanceAppointmentsController < ApplicationController
  def index
    @maintenanceAppointments = MaintenanceAppointment.all
  end
 
  def show
    @maintenanceAppointment = MaintenanceAppointment.find(params[:id])
  end
 
  def new
    @maintenanceAppointment = MaintenanceAppointment.new
  end
 
  def edit
    @maintenanceAppointment = MaintenanceAppointment.find(params[:id])
  end
 
  def create
    @maintenanceAppointment = MaintenanceAppointment.new(maintenanceAppointment_params)
 
    if @maintenanceAppointment.save
      redirect_to maintenanceAppointments_path
    else
      render 'new'
    end
  end
 
  def update
    @maintenanceAppointment = MaintenanceAppointment.find(params[:id])
 
    if @maintenanceAppointment.update(maintenanceAppointment_params)
      redirect_to maintenanceAppointments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @maintenanceAppointment = MaintenanceAppointment.find(params[:id])
    @maintenanceAppointment.destroy
    redirect_to maintenanceAppointments_path
  end

 
  private
    def maintenanceAppointment_params
      params.require(:maintenanceAppointment).permit(:appointmentDate, :Status)
    end
end