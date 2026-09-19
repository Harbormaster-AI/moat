class MedicalSuppliersController < ApplicationController
  def index
    @medicalSuppliers = MedicalSupplier.all
  end
 
  def show
    @medicalSupplier = MedicalSupplier.find(params[:id])
  end
 
  def new
    @medicalSupplier = MedicalSupplier.new
  end
 
  def edit
    @medicalSupplier = MedicalSupplier.find(params[:id])
  end
 
  def create
    @medicalSupplier = MedicalSupplier.new(medicalSupplier_params)
 
    if @medicalSupplier.save
      redirect_to medicalSuppliers_path
    else
      render 'new'
    end
  end
 
  def update
    @medicalSupplier = MedicalSupplier.find(params[:id])
 
    if @medicalSupplier.update(medicalSupplier_params)
      redirect_to medicalSuppliers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @medicalSupplier = MedicalSupplier.find(params[:id])
    @medicalSupplier.destroy
    redirect_to medicalSuppliers_path
  end

 
  private
    def medicalSupplier_params
      params.require(:medicalSupplier).permit(:name, :website, :SupplierTier)
    end
end