class CliniciansController < ApplicationController
  def index
    @clinicians = Clinician.all
  end
 
  def show
    @clinician = Clinician.find(params[:id])
  end
 
  def new
    @clinician = Clinician.new
  end
 
  def edit
    @clinician = Clinician.find(params[:id])
  end
 
  def create
    @clinician = Clinician.new(clinician_params)
 
    if @clinician.save
      redirect_to clinicians_path
    else
      render 'new'
    end
  end
 
  def update
    @clinician = Clinician.find(params[:id])
 
    if @clinician.update(clinician_params)
      redirect_to clinicians_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @clinician = Clinician.find(params[:id])
    @clinician.destroy
    redirect_to clinicians_path
  end

 
  private
    def clinician_params
      params.require(:clinician).permit(:firstName, :lastName, :licenseNumber, :ClinicianType, :Specialty)
    end
end