class AttestationsController < ApplicationController
  def index
    @attestations = Attestation.all
  end
 
  def show
    @attestation = Attestation.find(params[:id])
  end
 
  def new
    @attestation = Attestation.new
  end
 
  def edit
    @attestation = Attestation.find(params[:id])
  end
 
  def create
    @attestation = Attestation.new(attestation_params)
 
    if @attestation.save
      redirect_to attestations_path
    else
      render 'new'
    end
  end
 
  def update
    @attestation = Attestation.find(params[:id])
 
    if @attestation.update(attestation_params)
      redirect_to attestations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @attestation = Attestation.find(params[:id])
    @attestation.destroy
    redirect_to attestations_path
  end

 
  private
    def attestation_params
      params.require(:attestation).permit(:statement, :attestor, :dateSigned, :Result)
    end
end