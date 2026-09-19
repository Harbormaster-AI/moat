class AdmissionsController < ApplicationController
  def index
    @admissions = Admission.all
  end
 
  def show
    @admission = Admission.find(params[:id])
  end
 
  def new
    @admission = Admission.new
  end
 
  def edit
    @admission = Admission.find(params[:id])
  end
 
  def create
    @admission = Admission.new(admission_params)
 
    if @admission.save
      redirect_to admissions_path
    else
      render 'new'
    end
  end
 
  def update
    @admission = Admission.find(params[:id])
 
    if @admission.update(admission_params)
      redirect_to admissions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @admission = Admission.find(params[:id])
    @admission.destroy
    redirect_to admissions_path
  end

 
  private
    def admission_params
      params.require(:admission).permit(:admitDateTime, :bed, :AdmissionType)
    end
end