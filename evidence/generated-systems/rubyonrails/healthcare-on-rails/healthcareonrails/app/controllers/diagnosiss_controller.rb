class DiagnosissController < ApplicationController
  def index
    @diagnosiss = Diagnosis.all
  end
 
  def show
    @diagnosis = Diagnosis.find(params[:id])
  end
 
  def new
    @diagnosis = Diagnosis.new
  end
 
  def edit
    @diagnosis = Diagnosis.find(params[:id])
  end
 
  def create
    @diagnosis = Diagnosis.new(diagnosis_params)
 
    if @diagnosis.save
      redirect_to diagnosiss_path
    else
      render 'new'
    end
  end
 
  def update
    @diagnosis = Diagnosis.find(params[:id])
 
    if @diagnosis.update(diagnosis_params)
      redirect_to diagnosiss_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @diagnosis = Diagnosis.find(params[:id])
    @diagnosis.destroy
    redirect_to diagnosiss_path
  end

 
  private
    def diagnosis_params
      params.require(:diagnosis).permit(:code, :description, :onsetDate, :Certainty)
    end
end