class EvidencesController < ApplicationController
  def index
    @evidences = Evidence.all
  end
 
  def show
    @evidence = Evidence.find(params[:id])
  end
 
  def new
    @evidence = Evidence.new
  end
 
  def edit
    @evidence = Evidence.find(params[:id])
  end
 
  def create
    @evidence = Evidence.new(evidence_params)
 
    if @evidence.save
      redirect_to evidences_path
    else
      render 'new'
    end
  end
 
  def update
    @evidence = Evidence.find(params[:id])
 
    if @evidence.update(evidence_params)
      redirect_to evidences_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @evidence = Evidence.find(params[:id])
    @evidence.destroy
    redirect_to evidences_path
  end

 
  private
    def evidence_params
      params.require(:evidence).permit(:title, :locationUrl, :receivedDate, :EvidenceType)
    end
end