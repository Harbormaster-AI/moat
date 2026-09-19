class MedicationDispensesController < ApplicationController
  def index
    @medicationDispenses = MedicationDispense.all
  end
 
  def show
    @medicationDispense = MedicationDispense.find(params[:id])
  end
 
  def new
    @medicationDispense = MedicationDispense.new
  end
 
  def edit
    @medicationDispense = MedicationDispense.find(params[:id])
  end
 
  def create
    @medicationDispense = MedicationDispense.new(medicationDispense_params)
 
    if @medicationDispense.save
      redirect_to medicationDispenses_path
    else
      render 'new'
    end
  end
 
  def update
    @medicationDispense = MedicationDispense.find(params[:id])
 
    if @medicationDispense.update(medicationDispense_params)
      redirect_to medicationDispenses_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @medicationDispense = MedicationDispense.find(params[:id])
    @medicationDispense.destroy
    redirect_to medicationDispenses_path
  end

 
  private
    def medicationDispense_params
      params.require(:medicationDispense).permit(:dispenseNumber, :quantity, :whenPrepared, :Status)
    end
end